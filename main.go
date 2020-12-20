package main

import (
	"github.com/whereabouts/chassis/logger"
	"github.com/whereabouts/web-template/config"
	"github.com/whereabouts/web-template/engine/server"
	"github.com/whereabouts/web-template/routes"
)

func main() {
	if err := server.NewServer(config.GetConfig()).Router(routes.Routes).Run(); err != nil {
		logger.Fatalf("server run with http_error: %v", err)
	}
}
