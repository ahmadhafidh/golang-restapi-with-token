package models

// ResClient struct for returning value to client
type ResClient struct {
	Status  string      `json:"response_code"`
	Message string      `json:"response_message"`
	Data    interface{} `json:"data,omitempty"`
}
