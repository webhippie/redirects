package config

import (
	"time"
)

type yaml struct {
	Enabled bool
	File    string
}

type json struct {
	Enabled bool
	File    string
}

type toml struct {
	Enabled bool
	File    string
}

type zookeeper struct {
	Enabled   bool
	Endpoints []string
	Timeout   time.Duration
	Prefix    string
}

type etcd struct {
	Enabled   bool
	Endpoints []string
	Timeout   time.Duration
	Prefix    string
}

type consul struct {
	Enabled   bool
	Endpoints []string
	Timeout   time.Duration
	Prefix    string
}

type server struct {
	Addr        string
	Cert        string
	Key         string
	LetsEncrypt bool
	Pprof       bool
}

var (
	// Debug represents the flag to enable or disable debug logging.
	Debug bool

	// YAML represents the YAML storage configuration details.
	YAML = &yaml{}

	// JSON represents the JSON storage configuration details.
	JSON = &json{}

	// TOML represents the TOML storage configuration details.
	TOML = &toml{}

	// Zookeeper represents the Zookeeper storage configuration details.
	Zookeeper = &zookeeper{}

	// Etcd represents the Etcd storage configuration details.
	Etcd = &etcd{}

	// Consul represents the Consul storage configuration details.
	Consul = &consul{}

	// Server represents the informations about the server bindings.
	Server = &server{}
)
