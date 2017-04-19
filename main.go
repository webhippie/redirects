package main

import (
	"os"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/webhippie/redirects/cmd"
	"github.com/webhippie/redirects/config"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("REDIRECTS_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "redirects",
		Version:  config.Version,
		Usage:    "Simple pattern-based redirect server",
		Compiled: time.Now(),

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Activate debug information",
				EnvVars:     []string{"REDIRECTS_DEBUG"},
				Destination: &config.Debug,
				Hidden:      true,
			},
			&cli.BoolFlag{
				Name:        "yaml",
				Usage:       "Enable YAML storage",
				EnvVars:     []string{"REDIRECTS_YAML"},
				Destination: &config.YAML.Enabled,
			},
			&cli.StringFlag{
				Name:        "yaml-file",
				Value:       "file://storage/redirects.yaml",
				Usage:       "Define YAML storage file",
				EnvVars:     []string{"REDIRECTS_YAML_FILE"},
				Destination: &config.YAML.File,
			},
			&cli.BoolFlag{
				Name:        "json",
				Usage:       "Enable JSON storage",
				EnvVars:     []string{"REDIRECTS_JSON"},
				Destination: &config.JSON.Enabled,
			},
			&cli.StringFlag{
				Name:        "json-file",
				Value:       "file://storage/redirects.json",
				Usage:       "Define JSON storage file",
				EnvVars:     []string{"REDIRECTS_JSON_FILE"},
				Destination: &config.JSON.File,
			},
			&cli.BoolFlag{
				Name:        "toml",
				Usage:       "Enable TOML storage",
				EnvVars:     []string{"REDIRECTS_TOML"},
				Destination: &config.TOML.Enabled,
			},
			&cli.StringFlag{
				Name:        "toml-file",
				Value:       "file://storage/redirects.toml",
				Usage:       "Define TOML storage file",
				EnvVars:     []string{"REDIRECTS_TOML_FILE"},
				Destination: &config.TOML.File,
			},
			&cli.BoolFlag{
				Name:        "zk",
				Usage:       "Enable Zookeeper storage",
				EnvVars:     []string{"REDIRECTS_ZK"},
				Destination: &config.Zookeeper.Enabled,
			},
			&cli.StringSliceFlag{
				Name:    "zk-endpoint",
				Value:   cli.NewStringSlice("127.0.0.1:2181"),
				Usage:   "Used Zookeeper endpoints",
				EnvVars: []string{"REDIRECTS_ZK_ENDPOINTS"},
			},
			&cli.DurationFlag{
				Name:        "zk-timeout",
				Value:       10 * time.Second,
				Usage:       "Connection timeout for Zookeeper storage",
				EnvVars:     []string{"REDIRECTS_ZK_TIMEOUT"},
				Destination: &config.Zookeeper.Timeout,
			},
			&cli.StringFlag{
				Name:        "zk-prefix",
				Value:       "/redirects",
				Usage:       "Define Zookeeper storage prefix",
				EnvVars:     []string{"REDIRECTS_ZK_PREFIX"},
				Destination: &config.Zookeeper.Prefix,
			},
			&cli.BoolFlag{
				Name:        "etcd",
				Usage:       "Enable Etcd storage",
				EnvVars:     []string{"REDIRECTS_ETCD"},
				Destination: &config.Etcd.Enabled,
			},
			&cli.StringSliceFlag{
				Name:    "etcd-endpoint",
				Value:   cli.NewStringSlice("127.0.0.1:2379"),
				Usage:   "Used Etcd endpoints",
				EnvVars: []string{"REDIRECTS_ETCD_ENDPOINTS"},
			},
			&cli.DurationFlag{
				Name:        "etcd-timeout",
				Value:       10 * time.Second,
				Usage:       "Connection timeout for Etcd storage",
				EnvVars:     []string{"REDIRECTS_ETCD_TIMEOUT"},
				Destination: &config.Etcd.Timeout,
			},
			&cli.StringFlag{
				Name:        "etcd-prefix",
				Value:       "/redirects",
				Usage:       "Define Etcd storage prefix",
				EnvVars:     []string{"REDIRECTS_ETCD_PREFIX"},
				Destination: &config.Etcd.Prefix,
			},
			&cli.StringFlag{
				Name:        "etcd-username",
				Value:       "",
				Usage:       "Username to access Etcd",
				EnvVars:     []string{"REDIRECTS_ETCD_USERNAME"},
				Destination: &config.Etcd.Username,
			},
			&cli.StringFlag{
				Name:        "etcd-password",
				Value:       "",
				Usage:       "Password to access Etcd",
				EnvVars:     []string{"REDIRECTS_ETCD_PASSWORD"},
				Destination: &config.Etcd.Password,
			},
			&cli.StringFlag{
				Name:        "etcd-ca",
				Value:       "",
				Usage:       "CA certificate to access Etcd",
				EnvVars:     []string{"REDIRECTS_ETCD_CA"},
				Destination: &config.Etcd.CA,
			},
			&cli.StringFlag{
				Name:        "etcd-cert",
				Value:       "",
				Usage:       "SSL certificate to access Etcd",
				EnvVars:     []string{"REDIRECTS_ETCD_CERT"},
				Destination: &config.Etcd.Cert,
			},
			&cli.StringFlag{
				Name:        "etcd-key",
				Value:       "",
				Usage:       "SSL key to access Etcd",
				EnvVars:     []string{"REDIRECTS_ETCD_KEY"},
				Destination: &config.Etcd.Key,
			},
			&cli.BoolFlag{
				Name:        "etcd-skip-verify",
				Usage:       "Skip SSL verification for Consul",
				EnvVars:     []string{"REDIRECTS_ETCD_SKIP_VERIFY"},
				Destination: &config.Etcd.SkipVerify,
			},
			&cli.BoolFlag{
				Name:        "consul",
				Usage:       "Enable Consul storage",
				EnvVars:     []string{"REDIRECTS_CONSUL"},
				Destination: &config.Consul.Enabled,
			},
			&cli.StringSliceFlag{
				Name:    "consul-endpoint",
				Value:   cli.NewStringSlice("127.0.0.1:8500"),
				Usage:   "Used Consul endpoints",
				EnvVars: []string{"REDIRECTS_CONSUL_ENDPOINTS"},
			},
			&cli.DurationFlag{
				Name:        "consul-timeout",
				Value:       10 * time.Second,
				Usage:       "Connection timeout for Consul storage",
				EnvVars:     []string{"REDIRECTS_CONSUL_TIMEOUT"},
				Destination: &config.Consul.Timeout,
			},
			&cli.StringFlag{
				Name:        "consul-prefix",
				Value:       "/redirects",
				Usage:       "Define Consul storage prefix",
				EnvVars:     []string{"REDIRECTS_CONSUL_PREFIX"},
				Destination: &config.Consul.Prefix,
			},
			&cli.StringFlag{
				Name:        "consul-ca",
				Value:       "",
				Usage:       "CA certificate to access Consul",
				EnvVars:     []string{"REDIRECTS_CONSUL_CA"},
				Destination: &config.Consul.CA,
			},
			&cli.StringFlag{
				Name:        "consul-cert",
				Value:       "",
				Usage:       "SSL certificate to access Consul",
				EnvVars:     []string{"REDIRECTS_CONSUL_CERT"},
				Destination: &config.Consul.Cert,
			},
			&cli.StringFlag{
				Name:        "consul-key",
				Value:       "",
				Usage:       "SSL key to access Consul",
				EnvVars:     []string{"REDIRECTS_CONSUL_KEY"},
				Destination: &config.Consul.Key,
			},
			&cli.BoolFlag{
				Name:        "consul-skip-verify",
				Usage:       "Skip SSL verification for Consul",
				EnvVars:     []string{"REDIRECTS_CONSUL_SKIP_VERIFY"},
				Destination: &config.Consul.SkipVerify,
			},
		},

		Before: func(c *cli.Context) error {
			logrus.SetOutput(os.Stdout)

			if config.Debug {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}

			if len(c.StringSlice("zk-endpoint")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Zookeeper.Endpoints = c.StringSlice("zk-endpoint")
			}

			if len(c.StringSlice("etcd-endpoint")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Etcd.Endpoints = c.StringSlice("etcd-endpoint")
			}

			if len(c.StringSlice("consul-endpoint")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Consul.Endpoints = c.StringSlice("consul-endpoint")
			}

			return nil
		},

		Commands: []*cli.Command{
			cmd.Server(),
			cmd.List(),
			cmd.Show(),
			cmd.Create(),
			cmd.Update(),
			cmd.Remove(),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
