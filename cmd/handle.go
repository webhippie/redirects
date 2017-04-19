package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/webhippie/redirects/config"
	"github.com/webhippie/redirects/store"
	"github.com/webhippie/redirects/store/consul"
	"github.com/webhippie/redirects/store/etcd"
	"github.com/webhippie/redirects/store/json"
	"github.com/webhippie/redirects/store/toml"
	"github.com/webhippie/redirects/store/yaml"
	"github.com/webhippie/redirects/store/zookeeper"
	"gopkg.in/urfave/cli.v2"
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
	s, err := initStore()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}

	if err := fn(c, s); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(3)
	}

	return nil
}

// initStore initializes the store for CLI usage.
func initStore() (store.Store, error) {
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

	return nil, fmt.Errorf("No storage method specified")
}
