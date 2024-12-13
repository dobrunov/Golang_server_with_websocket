package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket connection configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Use origin check in production
	},
}

// WebSocket Handler function
func WSHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("New connection established successfully")

	// Send "Connected" message to client
	connectedMessage := Message{
		Type: "Connected",
		Data: json.RawMessage(`{"message": "Welcome!"}`),
	}
	if err := conn.WriteJSON(connectedMessage); err != nil {
		log.Println("Error sending connected message:", err)
		return
	}
	log.Println("Connected message sent:", prettyPrintJSON(connectedMessage))

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		log.Printf("Message received: %s\n", prettyPrintJSON(msg))

		// Parse the Data field
		var data Data
		if err := json.Unmarshal(msg.Data, &data); err != nil {
			var simpleData string
			if err := json.Unmarshal(msg.Data, &simpleData); err == nil && simpleData == "1" {
				data = Data{Value: 1}
			} else {
				log.Println("Error parsing Data:", err)
				continue
			}
		}

		// Check message type and send a response with type "UpdateCounter"
		if msg.Type == IncrementCounter && data.Value == 1 {
			response := Message{
				Type: UpdateCounter,
				Data: json.RawMessage(`{"value": 1}`),
			}

			err = conn.WriteJSON(response)
			if err != nil {
				log.Println("Error sending response:", err)
				return
			}
			log.Println("Response sent:", prettyPrintJSON(response))
		}
	}
}
