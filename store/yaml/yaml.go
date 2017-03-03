package yaml

import (
	"github.com/Sirupsen/logrus"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"path/filepath"
	"strings"
	"sync"
)

// data is a basic struct that iplements the Store interface.
type data struct {
	dsn   string
	mutex sync.Mutex
}

// New initializes a new YAML store.
func New(config string) store.Store {
	return &data{
		dsn:   config,
		mutex: sync.Mutex{},
	}
}

// Load initializes the YAML storage.
func Load() store.Store {
	driver := config.Storage.Driver
	connect := filepath.Clean(strings.TrimPrefix(config.Storage.DSN, "file://"))

	logrus.Infof("Using storage driver %s", driver)
	logrus.Infof("Using storage DSN %s", connect)

	return New(
		connect,
	)
}
