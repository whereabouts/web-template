package main

import (
	"github.com/whereabouts/chassis/logger"
	"github.com/whereabouts/web-template/config"
	"github.com/whereabouts/web-template/engine/server"
	"github.com/whereabouts/web-template/init"
	"github.com/whereabouts/web-template/routes"
)

func main() {
	// Quickly create an initial server and start it:
	// server.DefaultServer().Router(routes.Routes).Run()
	// if you need to use the custom configuration, call server.NewServer(conf)
	if err := server.NewServer(config.GetConfig()).Router(routes.Routes).Init(init.Init).Run(); err != nil {
		logger.Fatalf("server run with http_error: %v", err)
	}
}
