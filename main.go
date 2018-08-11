package main // import "github.com/qnib/go-webtest"

import (
	"github.com/codegangsta/cli"
	"github.com/qnib/go-webtest/lib"
	"os"
)



func main() {
	app := cli.NewApp()
	app.Name = "Webserver for testing purposes"
	app.Usage = "webtest [options]"
	app.Version = "0.1.2"
	app.Flags = webtest.DefaultFlags()
	app.Action = webtest.Run
	app.Run(os.Args)
}
