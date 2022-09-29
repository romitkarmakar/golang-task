package sockets

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Member struct {
	ID uuid.UUID
	Name string
	Conn *websocket.Conn
	Pool *Pool
}

type setNameReqBody struct {
	Name string `json:"name"`
}

type messageReqBody struct {
	Message string `json:"message"`
}

func closeConn(member *Member) {
	member.Pool.Unregister <- member
	member.Conn.Close()
}

func(member *Member) Read(){
	defer closeConn(member) // close connection at end of function
	for {
		messageType, p, err := member.Conn.ReadMessage()

		if err != nil {
			fmt.Printf("\nerror while reading message of %s",member.ID)
			continue
		}

		setName := &setNameReqBody{}
		messageReqBody := &messageReqBody{}

		memberError := MemberError{
			Member: member,
			Error: "",
		}

		err = json.Unmarshal(p,&messageReqBody)

		if err != nil {
			memberError.Error = "Invalid Format"
			member.Pool.Error <- memberError
			continue
		
		}

		if(messageReqBody.Message == "") {
			// look for name if not message
			err = json.Unmarshal(p,&setName)
			if err != nil || setName.Name == "" {
				memberError.Error = "Invalid Format"
				member.Pool.Error <- memberError
				continue
			
			}
			member.Name = setName.Name
			member.Pool.AssignName <- member
			memberError.Error = "Name Changed to " + member.Name
			member.Pool.Error <- memberError
			continue
		}

		if member.Name == "" {
			memberError.Error = "Set a name first"
			member.Pool.Error <- memberError
			continue
		}

		message := Message{
			Type: messageType,
			Body: member.Name + string(messageReqBody.Message),
		}

		member.Pool.Broadcast <- message
	}
}
