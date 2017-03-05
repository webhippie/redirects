package toml

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

// New initializes a new TOML store.
func New(f string) store.Store {
	return &data{
		file:  f,
		mutex: sync.Mutex{},
	}
}

// Load initializes the TOML storage.
func Load() store.Store {
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
				// TODO: Handle this error properly
				panic(fmt.Sprintf("TODO: Failed to create storage folder"))
			}
		}

		if err := ioutil.WriteFile(f, []byte(""), 0640); err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to create storage file"))
		}
	}

	return New(
		f,
	)
}
