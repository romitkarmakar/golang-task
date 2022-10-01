package routes

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/sswastik02/PublicRoom/sockets"
)

func ServeWs(receiver bool)(func(pool *sockets.Pool, w http.ResponseWriter, r *http.Request)) {

return func (pool *sockets.Pool,w http.ResponseWriter,r *http.Request) {

	ws,err := sockets.Upgrader.Upgrade(w,r,nil)
	// Upgrade TCP to Socket

	if err != nil {
		fmt.Println(err)
		return
	}

	member := &sockets.Member{ID: uuid.New(),Conn: ws,Pool: pool,Receiver: receiver}

	pool.Register <- member
	member.Read()
}
}