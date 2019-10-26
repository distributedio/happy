# TiKV Spanner transection interface SDK

## RO/RW Transection lifetime

### Begin()
* do nothing, gen txn struct with RO/RW

### Get()
* get TS if not exist
* start heartbeat with transection nodes
* call tikv get RO/RW with TS

### Set()/Delete()
* store key value in local storage(map)

### Commit()
* map & gather all keys with regions
* choose coordinator leader
* call commit on every perticipators and leader
