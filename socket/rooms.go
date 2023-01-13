package socket

import (
	"fmt"
	"letschat/models"
	"math/rand"
	"time"
)

type Room struct {
	Id         int32
	Client     map[*Client]bool
	Broadcast  chan *models.Message
	Register   chan *Client
	Unregister chan *Client
}

func NewRoom() *Room {
	rand.Seed(time.Now().UnixNano())
	Room := &Room{
		Id:         rand.Int31(),
		Broadcast:  make(chan *models.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Client:     make(map[*Client]bool),
	}
	println("creating new room")
	go Room.run()
	return Room
}

func (r *Room) run() {
	for {

		select {
		case client := <-r.Register:
			fmt.Println("new connection is registered in room")
			r.Client[client] = true
			fmt.Println("Size of connection in this room", len(r.Client))
			for client := range r.Client {

				if err := client.Conn.WriteJSON(models.Message{Type: "added", Message: "new user added"}); err != nil {
				}
			}
		case client := <-r.Unregister:
			if _, ok := r.Client[client]; ok {
				delete(r.Client, client)
				close(client.Send)
				for client := range r.Client {
					fmt.Println("broadcasting message to all users from room ")
					if err := client.Conn.WriteJSON(models.Message{Type: "removed", Message: "user removed"}); err != nil {
					}
				}
			}
		case message := <-r.Broadcast:
			for client := range r.Client {
				println(message.Message)
				println("message is send to all clients")
				if err := client.Conn.WriteJSON(message); err != nil {
				}
			}

		}
	}
}
