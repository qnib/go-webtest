package webtest

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net"
	"log"
	"net/http"
	"github.com/codegangsta/cli"
	)

func Run(ctx *cli.Context) {
	r := httprouter.New()

	//TODO: Add Prometheus
	// r.GET("/metrics",
	r.GET("/", getIP)
	r.GET("/cntname", getName)
	r.GET("/task", getTask)
	r.GET("/gpus", getGPUs)

	addr := fmt.Sprintf("%s:%s", ctx.String("http-host"), ctx.String("http-port"))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start Webserver on %s (v%s)", addr, ctx.App.Version)
	log.Fatal(http.Serve(l, r))
}