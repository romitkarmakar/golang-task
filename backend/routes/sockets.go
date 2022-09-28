package routes

import (
	"fmt"
	"net/http"

	"github.com/sswastik02/PublicRoom/sockets"
)

func ServeWs(w http.ResponseWriter,r *http.Request) {
	fmt.Println(r.Host)
	// Print Host

	ws,err := sockets.Upgrader.Upgrade(w,r,nil)
	// Upgrade TCP to Socket

	if err != nil {
		fmt.Println(err)
	}

	sockets.Reader(ws)
}
