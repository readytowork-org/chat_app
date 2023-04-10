package socket

import (
	"encoding/json"
	"fmt"
	"letschat/api/helper"
	"letschat/api/services"
	"letschat/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
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
	ID             string `json:"id"`
	wsServer       *WsServer
	Conn           *websocket.Conn
	Send           chan []byte
	rooms          map[*Room]bool
	userService    services.UserService
	roomService    services.RoomService
	messageService services.MessageService
}

func ServeWs(wsServer *WsServer, c *gin.Context) {
	conn, err := helper.Upgrade(c.Writer, c.Request)
	if err != nil {
		println("the errror is", err)
	}
	id := c.Query("id")

	if len(id) < 1 {
		log.Println("Url Param 'Id' is missing")
		return
	}
	client := *newClient(conn, wsServer, id)
	wsServer.Register <- &client
	client.connect()
	go client.writeMessage()
	go client.readMessage()
}

func newClient(conn *websocket.Conn, wsServer *WsServer, name string) *Client {
	//validate the client id
	return &Client{
		ID:             name,
		Conn:           conn,
		rooms:          make(map[*Room]bool),
		wsServer:       wsServer,
		Send:           make(chan []byte),
		userService:    wsServer.userService,
		roomService:    wsServer.roomService,
		messageService: wsServer.messageService,
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

func (client *Client) connect() {
	fmt.Println(client.ID)
	err := client.userService.UpdateUserStatus(client.ID, true)
	if err != nil {
		fmt.Println("there is error while updating user status")
	}
	rooms, err := client.userService.GetAllRooms(client.ID)
	if err != nil {
		fmt.Println("error while fetching user rooms")
	}
	onlineRooms := client.wsServer.FindMultipleRoomByID(rooms)
	for _, onlineRoom := range onlineRooms {
		onlineRoom.Register <- client
		client.rooms[onlineRoom] = true
	}
}

func (client *Client) disconnect() {
	err := client.userService.UpdateUserStatus(client.ID, false)
	if err != nil {
		fmt.Println(err)
		fmt.Println("there is error while updating user status")
	}
	//todo
	for room := range client.rooms {
		room.Unregister <- client
		if len(room.Clients) == 0 {
			client.wsServer.DeleteRoom(room)
		}
	}
	client.wsServer.Unregister <- client
}

func (client *Client) handleNewMessages(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}
	message.Sender = client.ID

	switch message.Action {
	case SendMessageAction:
		//save the message in database over here
		client.handleSendMessage(message)

	case CreateRoomAction:
		client.handleCreateRoom(message)

	case DeleteRoomAction:
		client.handleDeleteRoom(message)

	case DeleteMessageAction:
		client.handleDeleteMessage(message)

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
	dbMessage := models.MessageM{
		Message:    message.Message,
		Created_At: time.Now(),
		Created_By: message.Sender,
		Type:       message.Type,
		Room_Id:    message.RoomId,
	}
	dbMessage, err := client.messageService.Create(dbMessage)
	if err != nil {
		fmt.Println("Error while deleting the message")
		return
	}
	err = client.roomService.UpdateLastMessage(dbMessage.Room_Id, dbMessage.MessageId.Hex())
	if err != nil {
		fmt.Println("Error while deleting the message")
		return
	}
	room.Broadcast <- &message
}

func (client *Client) handleDeleteMessage(message Message) {
	room := client.findRoomByID(message.RoomId)
	if room == nil {
		println("The room you are trying to delete message doesnot present")
		return
	}
	err := client.messageService.Delete(message.MessageId)
	if err != nil {
		fmt.Println("Error while deleting the message")
	}

	// todo update last message id in database if this message was last
	room.Broadcast <- &message
}

func (client *Client) handleCreateRoom(message Message) {
	var room models.Room
	room = models.Room{
		Name:       "Any room",
		Createdby:  client.ID,
		ModifeidAt: time.Now(),
		StartedAt:  time.Now(),
	}
	room.Members = []string{client.ID}
	if len(message.Members) == 0 {
		fmt.Println("members are empty")
	} else {
		fmt.Println("There are members")
		for _, member := range message.Members {
			_, err := client.userService.FindOne(member)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Println("No user found")
				} else {
					fmt.Println("Some thing went wrong")
				}
				return
			}
			room.Members = append(room.Members, member)
		}
	}
	room, err := client.roomService.Create(room)
	if err != nil {
		fmt.Println("cannot create room")
	}
	for _, member := range room.Members {
		client.userService.AddRoom(member, room.RoomId.Hex())
	}
	wsRoom := client.wsServer.createRoom(message.RoomId, false)
	Clients := client.wsServer.FindMultipleClientsByID(room.Members)
	for _, clie := range Clients {
		fmt.Println("this is the online client")
		fmt.Println(clie)
		wsRoom.Register <- clie
		clie.rooms[wsRoom] = true
	}
	wsRoom.Broadcast <- &message
}

func (client *Client) handleDeleteRoom(message Message) {
	members, err := client.roomService.GetAllMembers(message.RoomId)
	if err != nil {
		fmt.Println("there is no room")
	}
	for _, member := range members {
		err := client.userService.DeleteRoom(member, message.RoomId)
		if err != nil {
			fmt.Println("error while delete")
		}
	}
	err = client.roomService.DeleteAllMember(message.RoomId)
	if err != nil {
		fmt.Println("Error while deleting Members from room")
	}
}

func (client *Client) handleJoinRoomMessage(message Message) {
	err := client.roomService.AddMember(message.RoomId, message.Members[0])
	if err != nil {
		println("there is error while adding member to the databse")
		return
	}
	err = client.userService.AddRoom(message.Members[0], message.RoomId)
	if err != nil {
		println("there is error while adding room to the user databse")
		return
	}
	dbMessage := models.MessageM{
		Message:    message.Message,
		Created_At: time.Now(),
		Created_By: message.Sender,
		Type:       message.Type,
		Room_Id:    message.RoomId,
	}
	dbMessage, err = client.messageService.Create(dbMessage)
	if err != nil {
		fmt.Println("Error while creating the message")
		return
	}
	err = client.roomService.UpdateLastMessage(dbMessage.Room_Id, dbMessage.MessageId.Hex())
	if err != nil {
		fmt.Println("Error while updating messsage the message")
		return
	}
	room := client.wsServer.findRoomByID(message.RoomId)
	if room == nil {
		members, err := client.roomService.GetAllMembers(message.RoomId)
		if err != nil {
			return
		}
		room = client.wsServer.createRoom(message.RoomId, false)
		Clients := client.wsServer.FindMultipleClientsByID(members)
		for _, clie := range Clients {
			room.Register <- clie
			clie.rooms[room] = true
		}
	}
	room.Broadcast <- &message
}

func (client *Client) handleLeaveRoomMessage(message Message) {
	err := client.roomService.DeleteMember(message.RoomId, message.Members[0])
	if err != nil {
		println("there is error while adding member to the databse")
		return
	}
	err = client.userService.DeleteRoom(message.Members[0], message.RoomId)
	if err != nil {
		println("there is error while adding room to the user databse")
		return
	}
	dbMessage := models.MessageM{
		Message:    message.Message,
		Created_At: time.Now(),
		Created_By: message.Sender,
		Type:       message.Type,
		Room_Id:    message.RoomId,
	}
	dbMessage, err = client.messageService.Create(dbMessage)
	if err != nil {
		fmt.Println("Error while creating the message")
		return
	}
	err = client.roomService.UpdateLastMessage(dbMessage.Room_Id, dbMessage.MessageId.Hex())
	if err != nil {
		fmt.Println("Error while updating messsage the message")
		return
	}
	room := client.wsServer.findRoomByID(message.RoomId)
	if room == nil {
		members, err := client.roomService.GetAllMembers(message.RoomId)
		if err != nil {
			return
		}
		room = client.wsServer.createRoom(message.RoomId, false)
		Clients := client.wsServer.FindMultipleClientsByID(members)
		for _, clie := range Clients {
			room.Register <- clie
			clie.rooms[room] = true
		}
	}
	room.Broadcast <- &message
}

// message read update -> to the databse and also to ther users  of that romm who are online
func (client *Client) findRoomByID(roomID string) *Room {
	var Room *Room
	//find room in client itself
	for room := range client.rooms {
		if room.GetId() == roomID {
			Room = room
			break
		}
	}
	if Room == nil {
		// check if client is in the room
		present, err := client.userService.IsRoomPresent(client.ID, roomID)
		if err != nil {
			fmt.Println("There is problem while finding room")
			return nil
		}
		if present {
			// check if room is in websocketserver
			Room = client.wsServer.findRoomByID(roomID)
			if Room == nil {
				roo, err := client.roomService.FindOne(roomID)
				if err != nil {
					println("There is no room associated with the id")
					return nil
				}
				Room = client.wsServer.createRoom(roo.RoomId.Hex(), false)
				Clients := client.wsServer.FindMultipleClientsByID(roo.Members)
				for _, clie := range Clients {
					Room.Register <- clie
					clie.rooms[Room] = true
				}
			}
		}
	}
	return Room
}
