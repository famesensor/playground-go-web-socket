package main

import (
	"github.com/famesensor/go-template/config"
	"github.com/famesensor/go-template/handler"
	"github.com/famesensor/go-template/infra"
)

func main() {
	cfg := config.NewConfig()
	upgrader := infra.NewWebsocketUpgrader()

	wsHdl := handler.NewWebsocketHandler(upgrader)

	httpServer := infra.NewHTTPServer(cfg)

	httpServer.RegisterRoute(wsHdl)

	httpServer.Run()
}
