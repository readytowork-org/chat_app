package socket

import (
	"encoding/json"
	"fmt"
)

const SendMessageAction = "send-message"
const DeleteMessageAction = "delete-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const CreateRoomAction = "create-room"
const DeleteRoomAction = "delete-room"

// todo message validation
type Message struct {
	MessageId string   `json:"messageId"`
	Action    string   `json:"action"`
	Message   string   `json:"message"`
	RoomId    string   `json:"roomId"`
	Sender    string   `json:"sender"`
	Members   []string `json:"members"`
	Type      string   `json:"type"`
}

func (message *Message) encode() []byte {
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		fmt.Printf("Message struct values: %#v\n", message)
		return nil
	}
	return jsonData
}
