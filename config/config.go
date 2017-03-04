package config

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

	// Server represents the informations about the server bindings.
	Server = &server{}
)
