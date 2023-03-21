package socket

const welcomeMessage = "%s joined the room"

type Room struct {
	ID         string `json:"id"`
	Private    bool   `json:"private"`
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

// NewRoom creates a new Room
func NewRoom(name string, private bool) *Room {
	return &Room{
		ID:         name,
		Private:    private,
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

// RunRoom runs our room, accepting various requests
func (room *Room) RunRoom() {
	for {
		select {

		case client := <-room.Register:
			room.registerClientInRoom(client)

		case client := <-room.Unregister:
			room.unregisterClientInRoom(client)

		case message := <-room.Broadcast:
			room.broadcastToClientsInRoom(message.encode(), message.Sender)
		}
	}
}
func (room *Room) GetId() string {
	return room.ID
}

// register and notify others users
func (room *Room) registerClientInRoom(client *Client) {
	// check if the user is in the room alredy or not if not ?
	// register clients in the database and register here
	room.Clients[client] = true
}

// unregister client from the room
func (room *Room) unregisterClientInRoom(client *Client) {
	// delete this client from the room database
	if _, ok := room.Clients[client]; ok {
		delete(room.Clients, client)
	}
}

func (room *Room) broadcastToClientsInRoom(message []byte, clientId string) {
	// save this message to room
	// broadcast to all online users
	// get all client from the database and if some clients are not online than send them messages as notification when they are online

	for client := range room.Clients {
		if client.ID != clientId {
			client.Send <- message
		}
	}
}
