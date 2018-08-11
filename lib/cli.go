package webtest

import (
	"github.com/codegangsta/cli"
)

func DefaultFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "http-host",
			Value:  "0.0.0.0",
			Usage:  "IP to bind webserver",
			EnvVar: "WEBTEST_HTTP_HOST",
		},
		cli.StringFlag{
			Name:   "http-port",
			Value:  "8081",
			Usage:  "IP to bind webserver",
			EnvVar: "WEBTEST_HTTP_PORT",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Be more verbose..",
			EnvVar: "RCUDAD_DEBUG",
		},
	}
}