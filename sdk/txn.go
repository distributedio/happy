// package sdk defined spanner sdk for tikv transection
package sdk

// Transection defiend spanner transection interface, implemented by ro transection and rw transection
type Transection interface{}

func Begin()
