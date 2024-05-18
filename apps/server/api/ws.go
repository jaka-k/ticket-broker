package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func OrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return // ends the loop if client disconnects
		}
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
