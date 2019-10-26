# TiKV Spanner transaction interface SDK

## RO/RW Transaction lifetime

### Begin()
* do nothing, gen txn struct with RO/RW

### Get()
* get TS if not exist
* start heartbeat with transaction nodes
* call tikv get RO/RW with TS

### Set()/Delete()
* store key value in local storage(map)

### Commit()
* map & gather all keys with regions
* choose coordinator leader
* call commit on every perticipators and leader
