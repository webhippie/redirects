package config

type storage struct {
	Driver string
	DSN  string
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

	// Storage represents the current storage configuration details.
	Storage = &storage{}

  // Server represents the informations about the server bindings.
  Server = &server{}
)
