package socket

import (
	"fmt"
	"letschat/models"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Room *Room
	Conn *websocket.Conn
	Send chan *models.Message
}

func (c *Client) Read() {
	defer func() {
		print("closing this connection")
		c.Room.Unregister <- c
		c.Conn.Close()
	}()
	for {
		messageType, p, err := c.Conn.ReadMessage()
		print("new message has been read")
		if err != nil {
			log.Panicln(err)
			return
		}

		message := models.Message{
			Message: string(p),
			Type:    fmt.Sprint(messageType),
		}

		c.Room.Broadcast <- &message
	}

}
