package sdk

import (
	"errors"
	"context"
	"github.com/distributedio/kvproto/pkg/spannerpb"
	"google.golang.org/grpc"
)

var (
	ErrNotFound = errors.New("not found")
	ErrRegionError = errors.New("region error")
)

type Type int
const (
	OpPut Type = 0;
	OpDelete Type = 1;
	OpRead Type = 2;

	ErrorCode_NotFound = 1;
	ErrorCode_RegionError = 2;
)

type Operation struct {
	Type Type
	Key []byte
	Value []byte
}

type TiKVClient interface{
	Get(ctx context.Context, txnID uint64, key []byte, version uint64, readOnly bool)([]byte, error)
	Commit(ctx context.Context, txnID uint64, operations []Operation, coordinatorId []byte, participants [][]byte) (uint64, error)
	HeartBeat(ctx context.Context, txnID uint64, ttl uint64) error
	Close() error
}

type tikvClient struct {
	conn *grpc.ClientConn
	cli spannerpb.SpannerClient
}

func NewTiKV(address string) (TiKVClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	cli := spannerpb.NewSpannerClient(conn)

	return &tikvClient{
		conn: conn,
		cli:cli,
	}, nil
}

func (kv *tikvClient) Get(ctx context.Context, txnID uint64, key []byte, version uint64, readOnly bool)([]byte, error) {
	req := &spannerpb.GetRequest{
		TxnId: txnID,
		Key: key,
		Version: version,
		ReadOnly: readOnly,
	}
	var resp *spannerpb.GetResponse
	resp, err := kv.cli.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.GetErrorCode() != 0 {
		// TODO
		return nil, ErrRegionError
	}

	return resp.Value, nil
}

func (kv *tikvClient)Commit(ctx context.Context, txnID uint64, operations []Operation, coordinatorId []byte, participantIDs [][]byte) (uint64, error) {
	req := &spannerpb.CommitRequest{
		TxnId: txnID,
		CoordinatorId: coordinatorId,
		ParticipantIds: participantIDs,
	}

	for i := range operations {
		op := &spannerpb.Operation{
			Key: operations[i].Key,
			Value: operations[i].Value,
		}
		switch operations[i].Type {
		case OpPut:
			op.Type = spannerpb.Type_Put
		case OpDelete:
			op.Type = spannerpb.Type_Delete
		case OpRead:
			op.Type = spannerpb.Type_Read
		default:
			return 0, errors.New("Unknown operation type")
		}
		req.Operations = append(req.Operations, op)
	}

	var resp *spannerpb.CommitResponse
	resp, err := kv.cli.Commit(ctx, req)
	if err != nil {
		return 0, err
	}

	if resp.ErrorCode != 0 {
		//TODO
		return 0, ErrRegionError
	}

	return resp.Version, nil
}

func (kv *tikvClient)HeartBeat(ctx context.Context, txnID uint64, ttl uint64) error {
	req := &spannerpb.HeartBeatRequest{
		TxnId:txnID,
		Ttl:ttl,
	}
	var resp *spannerpb.HeartBeatResponse
	resp, err := kv.cli.HeartBeat(ctx, req)
	if err != nil {
		return err
	}

	if resp.ErrorCode != 0 {
		//TODO
		return ErrRegionError
	}

	return nil
}

func (kv *tikvClient) Close() error {
	return kv.conn.Close()
}