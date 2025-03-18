package infra

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebsocketUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}
