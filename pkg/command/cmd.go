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

		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
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
	_ = viper.BindPFlag("config.file", rootCmd.PersistentFlags().Lookup("config-file"))

	rootCmd.PersistentFlags().String("log-level", "info", "Set logging level")
	viper.SetDefault("log.level", "info")
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().Bool("log-pretty", true, "Enable pretty logging")
	viper.SetDefault("log.pretty", true)
	_ = viper.BindPFlag("log.pretty", rootCmd.PersistentFlags().Lookup("log-pretty"))

	rootCmd.PersistentFlags().Bool("log-color", true, "Enable colored logging")
	viper.SetDefault("log.color", true)
	_ = viper.BindPFlag("log.color", rootCmd.PersistentFlags().Lookup("log-color"))

	rootCmd.PersistentFlags().Bool("yaml-enabled", defaultYAMLEnabled, "Enable YAML storage")
	viper.SetDefault("yaml.enabled", defaultYAMLEnabled)
	_ = viper.BindPFlag("yaml.enabled", rootCmd.PersistentFlags().Lookup("yaml-enabled"))

	rootCmd.PersistentFlags().String("yaml-file", defaultYAMLFile, "Define YAML storage file")
	viper.SetDefault("yaml.file", defaultYAMLFile)
	_ = viper.BindPFlag("yaml.file", rootCmd.PersistentFlags().Lookup("yaml-file"))

	rootCmd.PersistentFlags().Bool("json-enabled", defaultJSONEnabled, "Enable JSON storage")
	viper.SetDefault("json.enabled", defaultJSONEnabled)
	_ = viper.BindPFlag("json.enabled", rootCmd.PersistentFlags().Lookup("json-enabled"))

	rootCmd.PersistentFlags().String("json-file", defaultJSONFile, "Define JSON storage file")
	viper.SetDefault("json.file", defaultJSONFile)
	_ = viper.BindPFlag("json.file", rootCmd.PersistentFlags().Lookup("json-file"))

	rootCmd.PersistentFlags().Bool("toml-enabled", defaultTOMLEnabled, "Enable TOML storage")
	viper.SetDefault("toml.enabled", defaultTOMLEnabled)
	_ = viper.BindPFlag("toml.enabled", rootCmd.PersistentFlags().Lookup("toml-enabled"))

	rootCmd.PersistentFlags().String("toml-file", defaultTOMLFile, "Define TOML storage file")
	viper.SetDefault("toml.file", defaultTOMLFile)
	_ = viper.BindPFlag("toml.file", rootCmd.PersistentFlags().Lookup("toml-file"))

	rootCmd.PersistentFlags().Bool("zk-enabled", defaultZkEnabled, "Enable Zookeeper storage")
	viper.SetDefault("zk.enabled", defaultZkEnabled)
	_ = viper.BindPFlag("zk.enabled", rootCmd.PersistentFlags().Lookup("zk-enabled"))

	rootCmd.PersistentFlags().StringSlice("zk-endpoint", defaultZkEndpoints, "Used Zookeeper endpoints")
	viper.SetDefault("zk.endpoints", defaultZkEndpoints)
	_ = viper.BindPFlag("zk.endpoints", rootCmd.PersistentFlags().Lookup("zk-endpoint"))

	rootCmd.PersistentFlags().Duration("zk-timeout", defaultZkTimeout, "Connection timeout for Zookeeper storage")
	viper.SetDefault("zk.timeout", defaultZkTimeout)
	_ = viper.BindPFlag("zk.timeout", rootCmd.PersistentFlags().Lookup("zk-timeout"))

	rootCmd.PersistentFlags().String("zk-prefix", defaultZkPrefix, "Define Zookeeper storage prefix")
	viper.SetDefault("zk.prefix", defaultZkPrefix)
	_ = viper.BindPFlag("zk.prefix", rootCmd.PersistentFlags().Lookup("zk-prefix"))

	rootCmd.PersistentFlags().Bool("etcd-enabled", defaultTOMLEnabled, "Enable Etcd storage")
	viper.SetDefault("etcd.enabled", defaultEtcdEnabled)
	_ = viper.BindPFlag("etcd.enabled", rootCmd.PersistentFlags().Lookup("etcd-enabled"))

	rootCmd.PersistentFlags().StringSlice("etcd-endpoint", defaultEtcdEndpoints, "Used Etcd endpoints")
	viper.SetDefault("etcd.endpoints", defaultEtcdEndpoints)
	_ = viper.BindPFlag("etcd.endpoints", rootCmd.PersistentFlags().Lookup("etcd-endpoint"))

	rootCmd.PersistentFlags().Duration("etcd-timeout", defaultEtcdTimeout, "Connection timeout for Etcd storage")
	viper.SetDefault("etcd.timeout", defaultEtcdTimeout)
	_ = viper.BindPFlag("etcd.timeout", rootCmd.PersistentFlags().Lookup("etcd-timeout"))

	rootCmd.PersistentFlags().String("etcd-prefix", defaultEtcdPrefix, "Define Etcd storage prefix")
	viper.SetDefault("etcd.prefix", defaultEtcdPrefix)
	_ = viper.BindPFlag("etcd.prefix", rootCmd.PersistentFlags().Lookup("etcd-prefix"))

	rootCmd.PersistentFlags().String("etcd-username", defaultEtcdUsername, "Username to access Etcd")
	viper.SetDefault("etcd.username", defaultEtcdUsername)
	_ = viper.BindPFlag("etcd.username", rootCmd.PersistentFlags().Lookup("etcd-username"))

	rootCmd.PersistentFlags().String("etcd-password", defaultEtcdPassword, "Password to access Etcd")
	viper.SetDefault("etcd.password", defaultEtcdPassword)
	_ = viper.BindPFlag("etcd.password", rootCmd.PersistentFlags().Lookup("etcd-password"))

	rootCmd.PersistentFlags().String("etcd-ca", defaultEtcdCa, "Path to CA cert to access Etcd")
	viper.SetDefault("etcd.ca", defaultEtcdCa)
	_ = viper.BindPFlag("etcd.ca", rootCmd.PersistentFlags().Lookup("etcd-ca"))

	rootCmd.PersistentFlags().String("etcd-cert", defaultEtcdCert, "Path to SSL certificate to access Etcd")
	viper.SetDefault("etcd.cert", defaultEtcdCert)
	_ = viper.BindPFlag("etcd.cert", rootCmd.PersistentFlags().Lookup("etcd-cert"))

	rootCmd.PersistentFlags().String("etcd-key", defaultEtcdKey, "Path to SSL key to access Etcd")
	viper.SetDefault("etcd.key", defaultEtcdKey)
	_ = viper.BindPFlag("etcd.key", rootCmd.PersistentFlags().Lookup("etcd-key"))

	rootCmd.PersistentFlags().Bool("etcd-skip-verify", defaultEtcdSkipVerify, "Skip SSL verification for Etcd")
	viper.SetDefault("etcd.skip_verify", defaultEtcdSkipVerify)
	_ = viper.BindPFlag("etcd.skip_verify", rootCmd.PersistentFlags().Lookup("etcd-skip-verify"))

	rootCmd.PersistentFlags().Bool("consul-enabled", defaultTOMLEnabled, "Enable Consul storage")
	viper.SetDefault("consul.enabled", defaultConsulEnabled)
	_ = viper.BindPFlag("consul.enabled", rootCmd.PersistentFlags().Lookup("consul-enabled"))

	rootCmd.PersistentFlags().StringSlice("consul-endpoint", defaultConsulEndpoints, "Used Consul endpoints")
	viper.SetDefault("consul.endpoints", defaultConsulEndpoints)
	_ = viper.BindPFlag("consul.endpoints", rootCmd.PersistentFlags().Lookup("consul-endpoint"))

	rootCmd.PersistentFlags().Duration("consul-timeout", defaultConsulTimeout, "Connection timeout for Consul storage")
	viper.SetDefault("consul.timeout", defaultConsulTimeout)
	_ = viper.BindPFlag("consul.timeout", rootCmd.PersistentFlags().Lookup("consul-timeout"))

	rootCmd.PersistentFlags().String("consul-prefix", defaultConsulPrefix, "Define Consul storage prefix")
	viper.SetDefault("consul.prefix", defaultConsulPrefix)
	_ = viper.BindPFlag("consul.prefix", rootCmd.PersistentFlags().Lookup("consul-prefix"))

	rootCmd.PersistentFlags().String("consul-ca", defaultConsulCa, "Path to CA cert to access Consul")
	viper.SetDefault("consul.ca", defaultConsulCa)
	_ = viper.BindPFlag("consul.ca", rootCmd.PersistentFlags().Lookup("consul-ca"))

	rootCmd.PersistentFlags().String("consul-cert", defaultConsulCert, "Path to SSL certificate to access Consul")
	viper.SetDefault("consul.cert", defaultConsulCert)
	_ = viper.BindPFlag("consul.cert", rootCmd.PersistentFlags().Lookup("consul-cert"))

	rootCmd.PersistentFlags().String("consul-key", defaultConsulKey, "Path to SSL key to access Consul")
	viper.SetDefault("consul.key", defaultConsulKey)
	_ = viper.BindPFlag("consul.key", rootCmd.PersistentFlags().Lookup("consul-key"))

	rootCmd.PersistentFlags().Bool("consul-skip-verify", defaultConsulSkipVerify, "Skip SSL verification for Consul")
	viper.SetDefault("consul.skip_verify", defaultConsulSkipVerify)
	_ = viper.BindPFlag("consul.skip_verify", rootCmd.PersistentFlags().Lookup("consul-skip-verify"))
}

// Run parses the command line arguments and executes the program.
func Run() error {
	return rootCmd.Execute()
}
