// package sdk defined spanner sdk for tikv transection
package sdk

import (
	"errors"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	level = "debug"
)

func init() {
	// ConfigureZap customize the zap logger
	//func ConfigureZap(name, path, level, pattern string, compress bool) error {
	lv := zap.NewAtomicLevel()
	lv.SetLevel(zap.DebugLevel)
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02 15:04:05.999999999"))
	}
	zap.ReplaceGlobals(zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			NameKey:        "Name",
			StacktraceKey:  "Stack",
			MessageKey:     "Message",
			LevelKey:       "Level",
			TimeKey:        "TimeStamp",
			CallerKey:      "Caller",
			EncodeTime:     timeEncoder,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}), zapcore.AddSync(os.Stdout), lv), zap.AddCaller()))
}

var (
	Logger      = zap.L()
	ErrInternal = errors.New("error internal")
)
