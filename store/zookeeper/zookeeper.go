package zookeeper

import (
	"fmt"
	"github.com/docker/libkv"
	libkvStore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/zookeeper"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"time"
)

// init simply registers Zookeeper on the libkv library.
func init() {
	zookeeper.Register()
}

// data is a basic struct that iplements the Store interface.
type data struct {
	store  libkvStore.Store
	prefix string
}

// New initializes a new Zookeeper store.
func New(s libkvStore.Store, prefix string) store.Store {
	return &data{
		store:  s,
		prefix: prefix,
	}
}

// Load initializes the Zookeeper storage.
func Load() store.Store {
	prefix := config.Zookeeper.Prefix

	cfg := &libkvStore.Config{
		ConnectionTimeout: config.Zookeeper.Timeout * time.Second,
	}

	s, err := libkv.NewStore(
		libkvStore.ZK,
		config.Zookeeper.Endpoints,
		cfg,
	)

	// TODO: Handle this error properly
	if err != nil {
		panic(fmt.Sprintf("TODO: Failed to init store. %s", err))
	}

	if ok, _ := s.Exists(prefix); !ok {
		err := s.Put(
			prefix,
			nil,
			&libkvStore.WriteOptions{
				IsDir: true,
			},
		)

		// TODO: Handle this error properly
		if err != nil {
			panic(fmt.Sprintf("TODO: Failed to create prefix. %s", err))
		}
	}

	return New(
		s,
		prefix,
	)
}
