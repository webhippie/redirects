package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/webhippie/redirects/pkg/config"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
)

// collection represents the internal storage collection.
type collection struct {
	Redirects []*model.Redirect `json:"redirects"`
}

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

// load parses all available records from the storage.
func (db *data) load() (*collection, error) {
	res := &collection{
		Redirects: make([]*model.Redirect, 0),
	}

	content, err := ioutil.ReadFile(db.file)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(content, res); err != nil {
		return nil, err
	}

	return res, nil
}

// write writes the JSON content back to the storage.
func (db *data) write(content *collection) error {
	bytes, err := json.MarshalIndent(content, "", "  ")

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(db.file, bytes, 0640); err != nil {
		return err
	}

	return nil
}

// New initializes a new JSON store.
func New(f string) store.Store {
	return &data{
		file:  f,
		mutex: sync.Mutex{},
	}
}

// Load initializes the JSON storage.
func Load(cfg *config.JSON) (store.Store, error) {
	f := filepath.Clean(
		strings.TrimPrefix(
			cfg.File,
			"file://",
		),
	)

	if _, err := os.Stat(f); err != nil {
		dir := filepath.Dir(f)

		if _, dirErr := os.Stat(f); dir != "" && dirErr != nil {
			if err := os.MkdirAll(dir, 0750); err != nil {
				return nil, fmt.Errorf("failed to create storage folder: %w", err)
			}
		}

		if err := ioutil.WriteFile(f, []byte("{}"), 0640); err != nil {
			return nil, fmt.Errorf("failed to create storage file: %w", err)
		}
	}

	return New(
		f,
	), nil
}
