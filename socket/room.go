package socket

const welcomeMessage = "%s joined the room"

type Room struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
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
		Name:       name,
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

// get room name only if it is not the one on one room
func (room *Room) GetName() string {
	return room.Name
}

// register and notify others users
func (room *Room) registerClientInRoom(client *Client) {
	room.Clients[client] = true
}

// unregister client from the room
func (room *Room) unregisterClientInRoom(client *Client) {
	println("this is to runegister client in a room")
	if _, ok := room.Clients[client]; ok {
		delete(room.Clients, client)
	}
}

func (room *Room) broadcastToClientsInRoom(message []byte, clientId string) {

	for client := range room.Clients {
		if client.ID != clientId {
			client.Send <- message
		}
	}
}
