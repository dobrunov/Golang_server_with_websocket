package main

import (
	"log"
	"net/http"
	"simple_server/pkg/websocket"
)

func main() {
	// Register WebSocket handler at path "/"
	http.HandleFunc("/", websocket.WSHandler)

	// Start server on 127.0.0.1:8042
	log.Println("Server started on port :8042")
	log.Fatal(http.ListenAndServe("127.0.0.1:8042", nil))
}
