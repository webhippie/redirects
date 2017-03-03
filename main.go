package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/tboerger/redirects/cmd"
	"github.com/tboerger/redirects/config"
	"github.com/urfave/cli"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("REDIRECTS_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "redirects"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@tboerger.de>"
	app.Usage = "Simple pattern-based redirect server"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Activate debug information",
			EnvVar:      "REDIRECTS_DEBUG",
			Destination: &config.Debug,
		},
		cli.StringFlag{
			Name:        "driver",
			Value:       "yaml",
			Usage:       "Define the storage driver",
			EnvVar:      "REDIRECTS_DRIVER",
			Destination: &config.Storage.Driver,
		},
		cli.StringFlag{
			Name:        "dsn",
			Value:       "file://redirects.yml",
			Usage:       "Define the storage DSN",
			EnvVar:      "REDIRECTS_DSN",
			Destination: &config.Storage.DSN,
		},
	}

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Server(),
		cmd.List(),
		cmd.Show(),
		cmd.Create(),
		cmd.Update(),
		cmd.Remove(),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
