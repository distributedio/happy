// package sdk defined spanner sdk for tikv transection
package sdk

// Key give out spanner sdk key
type Key struct {
}

func WarpKey(rk string) *Key {
	return &Key{}
}
