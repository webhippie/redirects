package json

import (
	"github.com/Sirupsen/logrus"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
)

// data is a basic struct that iplements the Store interface.
type data struct {
	dsn string
}

// New initializes a new JSON store.
func New(config string) store.Store {
	return &data{
		dsn: config,
	}
}

// Load initializes the JSON storage.
func Load() store.Store {
	driver := config.Storage.Driver
	connect := config.Storage.DSN

	logrus.Infof("Using storage driver %s", driver)
	logrus.Infof("Using storage DSN %s", connect)

	return New(
		connect,
	)
}
