package socket

import (
	"encoding/json"
	"fmt"
	"letschat/api/helper"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second

	// Max time till next pong from peer
	pongWait = 60 * time.Second

	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// client must send their own id whiler joing the room	q
type Client struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	wsServer *WsServer
	Conn     *websocket.Conn
	Send     chan []byte
	rooms    map[*Room]bool //
}

func ServeWs(wsServer *WsServer, c *gin.Context) {
	conn, err := helper.Upgrade(c.Writer, c.Request)
	if err != nil {
		println("the errror is", err)
	}
	name := c.Query("name")

	if len(name) < 1 {
		log.Println("Url Param 'name' is missing")
		return
	}
	client := *newClient(conn, wsServer, name)
	wsServer.Register <- &client
	go client.writeMessage()
	go client.readMessage()
	fmt.Println(client)
}

func newClient(conn *websocket.Conn, wsServer *WsServer, name string) *Client {

	return &Client{
		ID:       name,
		Name:     name,
		Conn:     conn,
		rooms:    make(map[*Room]bool),
		wsServer: wsServer,
		Send:     make(chan []byte),
	}
}

func (client *Client) readMessage() {
	defer func() {
		client.disconnect()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(func(string) error { client.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, jsonMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		client.handleNewMessages(jsonMessage)
	}
}

func (client *Client) writeMessage() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) GetName() string {
	return client.Name
}

// disconnect client from the websocket server and all the rooms he/she was present in
func (client *Client) disconnect() {
	client.wsServer.Unregister <- client
	for room := range client.rooms {
		room.Unregister <- client
	}

}
func (client *Client) findRoomByID(ID string) *Room {
	var Room *Room
	for room := range client.rooms {
		if room.GetId() == ID {
			Room = room
			break
		}
	}
	return Room
}

func (client *Client) handleNewMessages(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}
	message.Sender = client.Name

	switch message.Action {
	case SendMessageAction:
		client.handleSendMessage(message)

	case JoinRoomAction:
		client.handleJoinRoomMessage(message)

	case LeaveRoomAction:
		client.handleLeaveRoomMessage(message)
	}
}

func (client *Client) handleSendMessage(message Message) {
	room := client.findRoomByID(message.RoomId)
	if room == nil {
		println("The room you are trying to send message doesnot present")
		return
	}
	room.Broadcast <- &message
}

func (client *Client) handleJoinRoomMessage(message Message) {
	roomName := message.RoomId
	client.joinRoom(roomName)
}
func (client *Client) joinRoom(roomName string) {
	room := client.wsServer.findRoomByID(roomName)
	if room == nil {
		room = client.wsServer.createRoom(roomName, false)
	}
	client.rooms[room] = true
	room.Register <- client
}

func (client *Client) handleLeaveRoomMessage(message Message) {
	roomId := message.RoomId
	room := client.wsServer.findRoomByID(roomId)
	if _, ok := client.rooms[room]; ok {
		delete(client.rooms, room)
	}
	room.Unregister <- client

}

func (client *Client) isInRoom(room *Room) bool {
	if _, ok := client.rooms[room]; ok {
		return true
	}
	return false
}
