package socket

import (
	"fmt"
	"letschat/api/services"
)

type WsServer struct {
	Client         map[*Client]bool
	Broadcast      chan []byte
	Register       chan *Client
	Unregister     chan *Client
	Rooms          map[*Room]bool
	userService    services.UserService
	roomService    services.RoomService
	messageService services.MessageService
}

func NewWebsocketServer(
	userService services.UserService,
	roomService services.RoomService,
	messageService services.MessageService,
) *WsServer {
	chatServer := &WsServer{
		Broadcast:      make(chan []byte),
		Register:       make(chan *Client),
		Unregister:     make(chan *Client),
		Client:         make(map[*Client]bool),
		Rooms:          make(map[*Room]bool),
		userService:    userService,
		roomService:    roomService,
		messageService: messageService,
	}
	go chatServer.Run()
	return chatServer
}
func (server *WsServer) Run() {
	for {
		select {

		case client := <-server.Register:
			server.registerClient(client)

		case client := <-server.Unregister:
			server.unregisterClient(client)

		}
	}
}

// register client to the server
func (server *WsServer) registerClient(client *Client) {
	server.Client[client] = true
}

// Delete the clent from the server after it's is lost its connection
func (server *WsServer) unregisterClient(client *Client) {
	// broadcast its connectiion to all clients in room associated with this clients
	// make user status offline
	if _, ok := server.Client[client]; ok {
		delete(server.Client, client)
	}
}

// create a room in database
func (server *WsServer) createRoom(name string, private bool) *Room {
	room := NewRoom(name, private)
	go room.RunRoom()
	server.Rooms[room] = true
	return room
}

func (server *WsServer) DeleteRoom(room *Room) {
	delete(server.Rooms, room)
}

// To find room by id .To add clients there . leave clients and send message to the room clients.
func (server *WsServer) findRoomByID(ID string) *Room {
	var foundRoom *Room
	for room := range server.Rooms {
		if room.GetId() == ID {	
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (server *WsServer) FindMultipleRoomByID(roonIDs []string) []*Room {
	var rooms []*Room
	for _, roomID := range roonIDs {
		for room := range server.Rooms {
			fmt.Println("this is the online rooms")
			print(room)
			if room.ID == roomID {
				rooms = append(rooms, room)
				break
			}
		}
	}
	return rooms
}

func (server *WsServer) findClientByID(ID string) *Client {
	var foundClient *Client
	for client := range server.Client {
		if client.ID == ID {
			foundClient = client
			break
		}
	}
	return foundClient
}

func (server *WsServer) FindMultipleClientsByID(clientIDs []string) []*Client {
	var Clients []*Client
	for _, clientID := range clientIDs {
		for client := range server.Client {
			if client.ID == clientID {
				Clients = append(Clients, client)
				break
			}
		}
	}
	return Clients
}
