package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connectionCollection struct {
	rooms     map[string]map[*connection]bool
	send      chan message
	subscribe chan subscription
}

type subscription struct {
	conn   *connection
	roomId *string
}

type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type message struct {
	message []byte
	room    string
}

var connections = connectionCollection{
	rooms:     make(map[string]map[*connection]bool),
	send:      make(chan message),
	subscribe: make(chan subscription),
}

func handler(w http.ResponseWriter, r *http.Request, roomId string) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil) // create connection
	//create a map that stores all the connection

	// create wiriteMessage that accepts the parameter of connection
	// create a readMessage that accepts the parameter of connection
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		log.Println(err)
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func readMessage(s subscription) {

}

func writeMessage(s subscription) {

}

func registerRoomID(room string) {

}

func Chat(c *gin.Context) {
	roomId := c.Param("roomId")
	handler(c.Writer, c.Request, roomId)
}
