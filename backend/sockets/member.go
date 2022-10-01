package sockets

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Member struct {
	ID uuid.UUID
	Receiver bool
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
			if err.Error() == "websocket: close 1001 (going away)"{
				break
			}
			continue
		}

		if (member.Receiver) {
			// Receivers cannot send message (no message to read for server)
			fmt.Println("Receiver Cannot send message")
			continue
		}

		setName := &setNameReqBody{}
		messageReqBody := &messageReqBody{}

		memberDesc := MemberPrivate{
			Member: member,
			Desc: "",
		}

		err = json.Unmarshal(p,&messageReqBody)

		if err != nil {
			memberDesc.Desc = "Invalid Format"
			member.Pool.Private <- memberDesc
			continue
		
		}

		if(messageReqBody.Message == "") {
			// look for name if not message
			err = json.Unmarshal(p,&setName)
			if err != nil || setName.Name == "" {
				memberDesc.Desc = "Invalid Format"
				member.Pool.Private <- memberDesc
				continue
			
			}
			member.Name = setName.Name
			member.Pool.AssignName <- member
			memberDesc.Desc = "Name Changed to " + member.Name
			member.Pool.Private <- memberDesc
			continue
		}

		if member.Name == "" {
			memberDesc.Desc = "Set a name first"
			member.Pool.Private <- memberDesc
			continue
		}

		message := MemberMessage{
			Member: member,
			Type: messageType,
			Body: fmt.Sprintf("%s : %s",member.Name,messageReqBody.Message),
		}

		member.Pool.Broadcast <- message
	}
}
