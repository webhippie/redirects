package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/webhippie/redirects/pkg/store"
	"github.com/webhippie/redirects/pkg/store/consul"
	"github.com/webhippie/redirects/pkg/store/etcd"
	"github.com/webhippie/redirects/pkg/store/json"
	"github.com/webhippie/redirects/pkg/store/toml"
	"github.com/webhippie/redirects/pkg/store/yaml"
	"github.com/webhippie/redirects/pkg/store/zookeeper"
)

func setupLogger() error {
	switch strings.ToLower(viper.GetString("log.level")) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if viper.GetBool("log.pretty") {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !viper.GetBool("log.color"),
			},
		)
	}

	var (
		err error
	)

	storage, err = setupStore()

	if err != nil {
		return err
	}

	return nil
}

func setupConfig() {
	if viper.GetString("config.file") != "" {
		viper.SetConfigFile(viper.GetString("config.file"))
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("/etc/redirects")
		viper.AddConfigPath("$HOME/.redirects")
		viper.AddConfigPath("./redirects")
	}

	viper.SetEnvPrefix("redirects")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); err != nil && !ok {
			log.Error().
				Err(err).
				Msg("Failed to read config file")
		}
	}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Error().
			Err(err).
			Msg("Failed to parse config file")
	}
}

func setupStore() (store.Store, error) {
	switch {
	case cfg.YAML.Enabled:
		return yaml.Load(&cfg.YAML)
	case cfg.JSON.Enabled:
		return json.Load(&cfg.JSON)
	case cfg.TOML.Enabled:
		return toml.Load(&cfg.TOML)
	case cfg.Zookeeper.Enabled:
		return zookeeper.Load(&cfg.Zookeeper)
	case cfg.Etcd.Enabled:
		return etcd.Load(&cfg.Etcd)
	case cfg.Consul.Enabled:
		return consul.Load(&cfg.Consul)
	}

	return nil, fmt.Errorf("no storage method define")
}
