package json

import (
	"fmt"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// data is a basic struct that iplements the Store interface.
type data struct {
	file  string
	mutex sync.Mutex
}

// Name simply returns the name of the store implementation.
func (db *data) Name() string {
	return "JSON"
}

// Config just returns a simple configuration explanation.
func (db *data) Config() string {
	return fmt.Sprintf("file:%s", db.file)
}

// New initializes a new JSON store.
func New(f string) store.Store {
	return &data{
		file:  f,
		mutex: sync.Mutex{},
	}
}

// Load initializes the JSON storage.
func Load() (store.Store, error) {
	f := filepath.Clean(
		strings.TrimPrefix(
			config.JSON.File,
			"file://",
		),
	)

	if _, err := os.Stat(f); err != nil {
		dir := filepath.Dir(f)

		if _, dirErr := os.Stat(f); dir != "" && dirErr != nil {
			if err := os.MkdirAll(dir, 0750); err != nil {
				return nil, fmt.Errorf("Failed to create storage folder")
			}
		}

		if err := ioutil.WriteFile(f, []byte("{}"), 0640); err != nil {
			return nil, fmt.Errorf("Failed to create storage file")
		}
	}

	return New(
		f,
	), nil
}
