package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var allowedOrigin = []string{"localhost", "krajnc.cc"}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		println("This is the r.URL.Hostname(): %v", r.URL.Hostname())
		return contains(allowedOrigin, r.URL.Hostname())
	},
}

func (s *APIServer) OrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println("Upgrade failed:", err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {

			println("ReadMessage failed:", err)
			break // Exit the loop on error
		}
		println("Received message: %s", p)
		if messageType == websocket.TextMessage {
			// Assume p is JSON and unmarshal -> TBD What the JSON will encapsulate
			// Update order status in your data store and send updates
			println("WS CONNECTION AND THE UPDATE FUNC ::: p === %s", p)
			// updateOrderStatus(p)
		}
		// Broadcast the updated status to the client
		conn.WriteMessage(websocket.TextMessage, []byte("Your updated status JSON here"))
	}

}

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
