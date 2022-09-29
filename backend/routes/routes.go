package routes

import (
	"net/http"

	"github.com/sswastik02/PublicRoom/sockets"
)

func SetupRoutes() {
	pool := sockets.CreatePool()
	go pool.Start()
	http.HandleFunc("/",InfoRoute)
	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request) {
		ServeWs(pool,w,r)
	})
}
