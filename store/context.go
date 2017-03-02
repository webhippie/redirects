package store

import (
	"golang.org/x/net/context"
)

const (
	storeKey = "store"
)

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext gets the store from the context.
func FromContext(c context.Context) Store {
	return c.Value(storeKey).(Store)
}

// ToContext injects the store into the context.
func ToContext(c Setter, store Store) {
	c.Set(storeKey, store)
}
