package models

type Message struct {
	Message  string `json:"message"`
	Type     string `json:"type"`
	ClientID string `json:"clientId"`
}
