package controllers

import (
	"encoding/json"
	"letschat/api/helper"
	"letschat/infrastructure"
	"letschat/socket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ThreadMessage thread message structure
// response from websocket server can be of type
// - "pong" echo response
type ThreadMessage struct {
	Type           string `json:"type"`
	Data           string `json:"data"`
	TextMessage    string `json:"text_message"`
	PictureMessage string `json:"picture_message"`
}

// UnmarshalData unmarshal data into given datatype
func (g ThreadMessage) UnmarshalData(dst interface{}) error {
	return json.Unmarshal([]byte(g.Data), dst)
}

// StreamThread stream thread struct
type ThreadController struct {
	logger infrastructure.Logger
	db     infrastructure.Database
}

// NewStreamThread creates new gift transfer
func NewThreadController(
	logger infrastructure.Logger,
	db infrastructure.Database,

) ThreadController {
	return ThreadController{
		logger: logger,
		db:     db,
	}
}

var ConnectionsData []ConnectionData

type ConnectionData struct {
	Connections *websocket.Conn
	RoomID      string
	UserID      string
}

// Handle handles new gift transfer request
func (tc *ThreadController) ServeWs(c *gin.Context) {
	conn, err := helper.Upgrade(c.Writer, c.Request)
	if err != nil {
		println("the errror is", err)
	}
	room := socket.NewRoom()

	client := &socket.Client{
		Room: room,
		Conn: conn,
	}
	if client == nil {
		println("client is nil")
	}
	println("clent =", &client)

	println("room is registed serveWs0")
	room.Register <- client
	println("room is registed serveWs1")
	println("clent =", &client)
	go client.Read()
}
