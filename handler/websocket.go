package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type websocketHandler struct {
	upgrader *websocket.Upgrader
}

func NewWebsocketHandler(upgrader *websocket.Upgrader) *websocketHandler {
	return &websocketHandler{
		upgrader: upgrader,
	}
}

func (h *websocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()

	// Listen for incoming messages
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s\n", message)
		// Echo the message back to the client
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
