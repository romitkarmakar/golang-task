package sockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true},
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(messageType,p)
		// Print Message


		// Echo Back
		err = conn.WriteMessage(messageType,p)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
