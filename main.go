package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Data structure for the Data field
type Data struct {
	Value int `json:"value"`
}

// Message structure
type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// Message type constants
const (
	UpdateCounter    = "UpdateCounter"
	IncrementCounter = "IncrementCounter"
	TestTwo          = "TestTwo"
)

// WebSocket connection configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Use origin check in production
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("New connection established successfully")

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		log.Printf("Message received: %v\n", msg)

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
			log.Println("Response sent:", response)
		}
	}
}

func main() {
	// Register WebSocket handler at path "/"
	http.HandleFunc("/", wsHandler)

	// Start server on 127.0.0.1:8042
	log.Println("Server started on port :8042")
	log.Fatal(http.ListenAndServe("127.0.0.1:8042", nil))
}
