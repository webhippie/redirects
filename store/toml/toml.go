package toml

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/webhippie/redirects/config"
	"github.com/webhippie/redirects/model"
	"github.com/webhippie/redirects/store"
)

// collection represents the internal storage collection.
type collection struct {
	Redirects []*model.Redirect `toml:"redirects"`
}

// data is a basic struct that iplements the Store interface.
type data struct {
	file  string
	mutex sync.Mutex
}

// Name simply returns the name of the store implementation.
func (db *data) Name() string {
	return "TOML"
}

// Config just returns a simple configuration explanation.
func (db *data) Config() string {
	return fmt.Sprintf("file:%s", db.file)
}

// load parses all available records from the storage.
func (db *data) load() (*collection, error) {
	res := &collection{
		Redirects: make([]*model.Redirect, 0),
	}

	content, err := ioutil.ReadFile(db.file)

	if err != nil {
		return nil, err
	}

	if _, err := toml.Decode(string(content), res); err != nil {
		return nil, err
	}

	return res, nil
}

// write writes the TOML content back to the storage.
func (db *data) write(content *collection) error {
	buf := new(bytes.Buffer)

	if err := toml.NewEncoder(buf).Encode(content); err != nil {
		return err
	}

	if err := ioutil.WriteFile(db.file, buf.Bytes(), 0640); err != nil {
		return err
	}

	return nil
}

// New initializes a new TOML store.
func New(f string) store.Store {
	return &data{
		file:  f,
		mutex: sync.Mutex{},
	}
}

// Load initializes the TOML storage.
func Load() (store.Store, error) {
	f := filepath.Clean(
		strings.TrimPrefix(
			config.TOML.File,
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

		if err := ioutil.WriteFile(f, []byte(""), 0640); err != nil {
			return nil, fmt.Errorf("Failed to create storage file")
		}
	}

	return New(
		f,
	), nil
}
