package main

import (
	"fmt"
	"net/http"

	"github.com/famesensor/go-template/config"
	"github.com/famesensor/go-template/handler"
	"github.com/famesensor/go-template/infra"
)

func main() {
	cfg := config.NewConfig()
	upgrader := infra.NewWebsocketUpgrader()

	wsHdl := handler.NewWebsocketHandler(upgrader)

	http.HandleFunc("/ws", wsHdl.Handle)
	fmt.Printf("WebSocket server started on :%s\n", cfg.App.Port)
	err := http.ListenAndServe(":"+cfg.App.Port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
