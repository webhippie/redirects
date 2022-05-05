package command

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/redirects/pkg/config"
	"github.com/webhippie/redirects/pkg/store"
	"github.com/webhippie/redirects/pkg/version"
)

var (
	rootCmd = &cobra.Command{
		Use:           "redirects",
		Short:         "Simple pattern-based redirect server",
		Version:       version.String,
		SilenceErrors: false,
		SilenceUsage:  true,

		PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {
			return setupLogger()
		},

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cfg     *config.Config
	storage store.Store

	defaultYAMLEnabled      = false
	defaultYAMLFile         = "file://storage/redirects.yaml"
	defaultJSONEnabled      = false
	defaultJSONFile         = "file://storage/redirects.json"
	defaultTOMLEnabled      = false
	defaultTOMLFile         = "file://storage/redirects.toml"
	defaultZkEnabled        = false
	defaultZkEndpoints      = []string{"127.0.0.1:2181"}
	defaultZkTimeout        = 10 * time.Second
	defaultZkPrefix         = "/redirects"
	defaultEtcdEnabled      = false
	defaultEtcdEndpoints    = []string{"127.0.0.1:2379"}
	defaultEtcdTimeout      = 10 * time.Second
	defaultEtcdPrefix       = "/redirects"
	defaultEtcdUsername     = ""
	defaultEtcdPassword     = ""
	defaultEtcdCa           = ""
	defaultEtcdCert         = ""
	defaultEtcdKey          = ""
	defaultEtcdSkipVerify   = false
	defaultConsulEnabled    = false
	defaultConsulEndpoints  = []string{"127.0.0.1:8500"}
	defaultConsulTimeout    = 10 * time.Second
	defaultConsulPrefix     = "/redirects"
	defaultConsulCa         = ""
	defaultConsulCert       = ""
	defaultConsulKey        = ""
	defaultConsulSkipVerify = false
)

func init() {
	cfg = config.Load()
	cobra.OnInitialize(setupConfig)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show the help, so what you see now")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the current version of that tool")

	rootCmd.PersistentFlags().String("config-file", "", "Path to optional config file")
	viper.BindPFlag("config.file", rootCmd.PersistentFlags().Lookup("config-file"))

	rootCmd.PersistentFlags().String("log-level", "info", "Set logging level")
	viper.SetDefault("log.level", "info")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().Bool("log-pretty", true, "Enable pretty logging")
	viper.SetDefault("log.pretty", true)
	viper.BindPFlag("log.pretty", rootCmd.PersistentFlags().Lookup("log-pretty"))

	rootCmd.PersistentFlags().Bool("log-color", true, "Enable colored logging")
	viper.SetDefault("log.color", true)
	viper.BindPFlag("log.color", rootCmd.PersistentFlags().Lookup("log-color"))

	rootCmd.PersistentFlags().Bool("yaml-enabled", defaultYAMLEnabled, "Enable YAML storage")
	viper.SetDefault("yaml.enabled", defaultYAMLEnabled)
	viper.BindPFlag("yaml.enabled", rootCmd.PersistentFlags().Lookup("yaml-enabled"))

	rootCmd.PersistentFlags().String("yaml-file", defaultYAMLFile, "Define YAML storage file")
	viper.SetDefault("yaml.file", defaultYAMLFile)
	viper.BindPFlag("yaml.file", rootCmd.PersistentFlags().Lookup("yaml-file"))

	rootCmd.PersistentFlags().Bool("json-enabled", defaultJSONEnabled, "Enable JSON storage")
	viper.SetDefault("json.enabled", defaultJSONEnabled)
	viper.BindPFlag("json.enabled", rootCmd.PersistentFlags().Lookup("json-enabled"))

	rootCmd.PersistentFlags().String("json-file", defaultJSONFile, "Define JSON storage file")
	viper.SetDefault("json.file", defaultJSONFile)
	viper.BindPFlag("json.file", rootCmd.PersistentFlags().Lookup("json-file"))

	rootCmd.PersistentFlags().Bool("toml-enabled", defaultTOMLEnabled, "Enable TOML storage")
	viper.SetDefault("toml.enabled", defaultTOMLEnabled)
	viper.BindPFlag("toml.enabled", rootCmd.PersistentFlags().Lookup("toml-enabled"))

	rootCmd.PersistentFlags().String("toml-file", defaultTOMLFile, "Define TOML storage file")
	viper.SetDefault("toml.file", defaultTOMLFile)
	viper.BindPFlag("toml.file", rootCmd.PersistentFlags().Lookup("toml-file"))

	rootCmd.PersistentFlags().Bool("zk-enabled", defaultZkEnabled, "Enable Zookeeper storage")
	viper.SetDefault("zk.enabled", defaultZkEnabled)
	viper.BindPFlag("zk.enabled", rootCmd.PersistentFlags().Lookup("zk-enabled"))

	rootCmd.PersistentFlags().StringSlice("zk-endpoint", defaultZkEndpoints, "Used Zookeeper endpoints")
	viper.SetDefault("zk.endpoint", defaultZkEndpoints)
	viper.BindPFlag("zk.endpoint", rootCmd.PersistentFlags().Lookup("zk-endpoint"))

	rootCmd.PersistentFlags().Duration("zk-timeout", defaultZkTimeout, "Connection timeout for Zookeeper storage")
	viper.SetDefault("zk.timeout", defaultZkTimeout)
	viper.BindPFlag("zk.timeout", rootCmd.PersistentFlags().Lookup("zk-timeout"))

	rootCmd.PersistentFlags().String("zk-prefix", defaultZkPrefix, "Define Zookeeper storage prefix")
	viper.SetDefault("zk.prefix", defaultZkPrefix)
	viper.BindPFlag("zk.prefix", rootCmd.PersistentFlags().Lookup("zk-prefix"))

	rootCmd.PersistentFlags().Bool("etcd-enabled", defaultTOMLEnabled, "Enable Etcd storage")
	viper.SetDefault("etcd.enabled", defaultEtcdEnabled)
	viper.BindPFlag("etcd.enabled", rootCmd.PersistentFlags().Lookup("etcd-enabled"))

	rootCmd.PersistentFlags().StringSlice("etcd-endpoint", defaultEtcdEndpoints, "Used Etcd endpoints")
	viper.SetDefault("etcd.endpoint", defaultEtcdEndpoints)
	viper.BindPFlag("etcd.endpoint", rootCmd.PersistentFlags().Lookup("etcd-endpoint"))

	rootCmd.PersistentFlags().Duration("etcd-timeout", defaultEtcdTimeout, "Connection timeout for Etcd storage")
	viper.SetDefault("etcd.timeout", defaultEtcdTimeout)
	viper.BindPFlag("etcd.timeout", rootCmd.PersistentFlags().Lookup("etcd-timeout"))

	rootCmd.PersistentFlags().String("etcd-prefix", defaultEtcdPrefix, "Define Etcd storage prefix")
	viper.SetDefault("etcd.prefix", defaultEtcdPrefix)
	viper.BindPFlag("etcd.prefix", rootCmd.PersistentFlags().Lookup("etcd-prefix"))

	rootCmd.PersistentFlags().String("etcd-username", defaultEtcdUsername, "Username to access Etcd")
	viper.SetDefault("etcd.username", defaultEtcdUsername)
	viper.BindPFlag("etcd.username", rootCmd.PersistentFlags().Lookup("etcd-username"))

	rootCmd.PersistentFlags().String("etcd-password", defaultEtcdPassword, "Password to access Etcd")
	viper.SetDefault("etcd.password", defaultEtcdPassword)
	viper.BindPFlag("etcd.password", rootCmd.PersistentFlags().Lookup("etcd-password"))

	rootCmd.PersistentFlags().String("etcd-ca", defaultEtcdCa, "Path to CA cert to access Etcd")
	viper.SetDefault("etcd.ca", defaultEtcdCa)
	viper.BindPFlag("etcd.ca", rootCmd.PersistentFlags().Lookup("etcd-ca"))

	rootCmd.PersistentFlags().String("etcd-cert", defaultEtcdCert, "Path to SSL certificate to access Etcd")
	viper.SetDefault("etcd.cert", defaultEtcdCert)
	viper.BindPFlag("etcd.cert", rootCmd.PersistentFlags().Lookup("etcd-cert"))

	rootCmd.PersistentFlags().String("etcd-key", defaultEtcdKey, "Path to SSL key to access Etcd")
	viper.SetDefault("etcd.key", defaultEtcdKey)
	viper.BindPFlag("etcd.key", rootCmd.PersistentFlags().Lookup("etcd-key"))

	rootCmd.PersistentFlags().Bool("etcd-skip-verify", defaultEtcdSkipVerify, "Skip SSL verification for Etcd")
	viper.SetDefault("etcd.skip_verify", defaultEtcdSkipVerify)
	viper.BindPFlag("etcd.skip_verify", rootCmd.PersistentFlags().Lookup("etcd-skip-verify"))

	rootCmd.PersistentFlags().Bool("consul-enabled", defaultTOMLEnabled, "Enable Consul storage")
	viper.SetDefault("consul.enabled", defaultConsulEnabled)
	viper.BindPFlag("consul.enabled", rootCmd.PersistentFlags().Lookup("consul-enabled"))

	rootCmd.PersistentFlags().StringSlice("consul-endpoint", defaultConsulEndpoints, "Used Consul endpoints")
	viper.SetDefault("consul.endpoint", defaultConsulEndpoints)
	viper.BindPFlag("consul.endpoint", rootCmd.PersistentFlags().Lookup("consul-endpoint"))

	rootCmd.PersistentFlags().Duration("consul-timeout", defaultConsulTimeout, "Connection timeout for Consul storage")
	viper.SetDefault("consul.timeout", defaultConsulTimeout)
	viper.BindPFlag("consul.timeout", rootCmd.PersistentFlags().Lookup("consul-timeout"))

	rootCmd.PersistentFlags().String("consul-prefix", defaultConsulPrefix, "Define Consul storage prefix")
	viper.SetDefault("consul.prefix", defaultConsulPrefix)
	viper.BindPFlag("consul.prefix", rootCmd.PersistentFlags().Lookup("consul-prefix"))

	rootCmd.PersistentFlags().String("consul-ca", defaultConsulCa, "Path to CA cert to access Consul")
	viper.SetDefault("consul.ca", defaultConsulCa)
	viper.BindPFlag("consul.ca", rootCmd.PersistentFlags().Lookup("consul-ca"))

	rootCmd.PersistentFlags().String("consul-cert", defaultConsulCert, "Path to SSL certificate to access Consul")
	viper.SetDefault("consul.cert", defaultConsulCert)
	viper.BindPFlag("consul.cert", rootCmd.PersistentFlags().Lookup("consul-cert"))

	rootCmd.PersistentFlags().String("consul-key", defaultConsulKey, "Path to SSL key to access Consul")
	viper.SetDefault("consul.key", defaultConsulKey)
	viper.BindPFlag("consul.key", rootCmd.PersistentFlags().Lookup("consul-key"))

	rootCmd.PersistentFlags().Bool("consul-skip-verify", defaultConsulSkipVerify, "Skip SSL verification for Consul")
	viper.SetDefault("consul.skip_verify", defaultConsulSkipVerify)
	viper.BindPFlag("consul.skip_verify", rootCmd.PersistentFlags().Lookup("consul-skip-verify"))
}

// Run parses the command line arguments and executes the program.
func Run() error {
	return rootCmd.Execute()
}
