package config

import (
	"time"
)

// Server defines the server configuration.
type Server struct {
	Addr          string `mapstructure:"addr"`
	Host          string `mapstructure:"host"`
	Pprof         bool   `mapstructure:"pprof"`
	Root          string `mapstructure:"root"`
	Cert          string `mapstructure:"cert"`
	Key           string `mapstructure:"key"`
	StrictCurves  bool   `mapstructure:"strict_curves"`
	StrictCiphers bool   `mapstructure:"strict_ciphers"`
	Templates     string `mapstructure:"templates"`
}

// Metrics defines the metrics server configuration.
type Metrics struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// YAML defines the yaml configuration.
type YAML struct {
	Enabled bool   `mapstructure:"enabled"`
	File    string `mapstructure:"file"`
}

// JSON defines the json configuration.
type JSON struct {
	Enabled bool   `mapstructure:"enabled"`
	File    string `mapstructure:"file"`
}

// TOML defines the toml configuration.
type TOML struct {
	Enabled bool   `mapstructure:"enabled"`
	File    string `mapstructure:"file"`
}

// Zookeeper defines the zookeeper configuration.
type Zookeeper struct {
	Enabled   bool          `mapstructure:"enabled"`
	Endpoints []string      `mapstructure:"endpoints"`
	Timeout   time.Duration `mapstructure:"timeout"`
	Prefix    string        `mapstructure:"prefix"`
}

// Etcd defines the etcd configuration.
type Etcd struct {
	Enabled    bool          `mapstructure:"enabled"`
	Endpoints  []string      `mapstructure:"endpoints"`
	Timeout    time.Duration `mapstructure:"timeout"`
	Prefix     string        `mapstructure:"prefix"`
	Username   string        `mapstructure:"username"`
	Password   string        `mapstructure:"password"`
	CA         string        `mapstructure:"ca"`
	Cert       string        `mapstructure:"cert"`
	Key        string        `mapstructure:"key"`
	SkipVerify bool          `mapstructure:"skip_verify"`
}

// Consul defines the consul configuration.
type Consul struct {
	Enabled    bool          `mapstructure:"enabled"`
	Endpoints  []string      `mapstructure:"endpoints"`
	Timeout    time.Duration `mapstructure:"timeout"`
	Prefix     string        `mapstructure:"prefix"`
	CA         string        `mapstructure:"ca"`
	Cert       string        `mapstructure:"cert"`
	Key        string        `mapstructure:"key"`
	SkipVerify bool          `mapstructure:"skip_verify"`
}

// Config defines the general configuration.
type Config struct {
	Server    Server    `mapstructure:"server"`
	Metrics   Metrics   `mapstructure:"metrics"`
	Logs      Logs      `mapstructure:"log"`
	YAML      YAML      `mapstructure:"yaml"`
	JSON      JSON      `mapstructure:"json"`
	TOML      TOML      `mapstructure:"toml"`
	Zookeeper Zookeeper `mapstructure:"zk"`
	Etcd      Etcd      `mapstructure:"etcd"`
	Consul    Consul    `mapstructure:"consul"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
