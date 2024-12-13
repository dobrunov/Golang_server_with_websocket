package websocket

import "encoding/json"

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
)
