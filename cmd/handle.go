package cmd

import (
	"fmt"
	"os"

	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"github.com/tboerger/redirects/store/json"
	"github.com/tboerger/redirects/store/yaml"
	"github.com/urfave/cli"
)

// HandleFunc is the real handle implementation.
type HandleFunc func(c *cli.Context, s store.Store) error

// Handle wraps the command function handler.
func Handle(c *cli.Context, fn HandleFunc) error {
	s := initStore()

	if err := fn(c, s); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}

	return nil
}

// initStore initializes the store for CLI usage.
func initStore() store.Store {
	switch config.Storage.Driver {
	case "yaml":
		return yaml.Load()
	case "json":
		return json.Load()
	}

	return nil
}
