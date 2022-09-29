package sockets

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/slices"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true},
}

// func Reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		fmt.Println(messageType,p)
// 		// Print Message


// 		// Echo Back
// 		err = conn.WriteMessage(messageType,p)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 	}
// }

func broadcast(pool *Pool,message *Message) {
	for _,member := range pool.Members {
		member.Conn.WriteJSON(message.Body)
	}
}

func (pool *Pool) Start() {
	for {
		select {

		case member := <-pool.Register:
			fmt.Printf("\nNew member joined: %s",member.ID)
			// let other members know that new member has joined
			broadcast(pool,&Message{Type: 1, Body: fmt.Sprintf("New Member Joined : %s",member.ID)})
			// then add to pool
			pool.Members = append(pool.Members, member)
			continue

		case member := <-pool.Unregister:
			fmt.Printf("\n Member left : %s",member.ID)
			idx := slices.IndexFunc(pool.Members,func(m *Member) bool {return m.ID == member.ID})
			slices.Delete(pool.Members,idx,idx)
			broadcast(pool,&Message{Type: 1, Body: fmt.Sprintf("Member Left : %s",member.ID)})
			continue

		case message := <-pool.Broadcast:
			broadcast(pool,&message)
			continue

		case memberError := <-pool.Error:
			memberError.Member.Conn.WriteJSON(
				struct {
					Error string `json:"error"`
				} {
					Error: memberError.Error,
				},
			)
			continue

		case member := <-pool.AssignName:
			idx := slices.IndexFunc(pool.Members,func(m *Member) bool {return m.ID == member.ID})
			fmt.Printf("Assigning Name %s to %s",member.Name,member.ID)
			pool.Members[idx] = member
			continue
		}
	}	
}
