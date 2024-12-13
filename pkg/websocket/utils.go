package websocket

import "encoding/json"

// Utility function to print JSON as a single line
func prettyPrintJSON(v interface{}) string {
	data, err := json.Marshal(v) // Serialize to JSON without indentation
	if err != nil {
		return "Error encoding JSON"
	}
	return string(data)
}
