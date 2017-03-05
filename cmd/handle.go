package cmd

import (
	"fmt"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"github.com/tboerger/redirects/store/consul"
	"github.com/tboerger/redirects/store/etcd"
	"github.com/tboerger/redirects/store/json"
	"github.com/tboerger/redirects/store/toml"
	"github.com/tboerger/redirects/store/yaml"
	"github.com/tboerger/redirects/store/zookeeper"
	"gopkg.in/urfave/cli.v2"
	"os"
	"strings"
	"text/template"
)

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"split":    strings.Split,
	"join":     strings.Join,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
}

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
	switch {
	case config.YAML.Enabled:
		return yaml.Load()
	case config.JSON.Enabled:
		return json.Load()
	case config.TOML.Enabled:
		return toml.Load()
	case config.Zookeeper.Enabled:
		return zookeeper.Load()
	case config.Etcd.Enabled:
		return etcd.Load()
	case config.Consul.Enabled:
		return consul.Load()
	}

	panic("No storage method specified")
}
