package handler

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type websocketHandler struct {
	upgrader *websocket.Upgrader
}

func NewWebsocketHandler(upgrader *websocket.Upgrader) *websocketHandler {
	return &websocketHandler{
		upgrader: upgrader,
	}
}

func (h *websocketHandler) RegisterRoutes(e *echo.Echo) {
	e.GET("/ws", h.Handle)
}

// func (h *websocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
func (h *websocketHandler) Handle(c echo.Context) error {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := h.upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Error upgrading:", err)
		return err
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

	return nil
}
