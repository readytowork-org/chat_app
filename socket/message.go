package socket

import (
	"encoding/json"
	"fmt"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const JoinRoomPrivateAction = "join-room-private"
const RoomJoinedAction = "room-joined"

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	RoomId  string `json:"roomId"`
	Sender  string `json:"sender"`
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
