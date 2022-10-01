package routes

import (
	"net/http"

	"github.com/sswastik02/PublicRoom/sockets"
)

func SetupRoutes() {
	pool := sockets.CreatePool()
	go pool.Start()
	http.HandleFunc("/",InfoRoute)
	http.HandleFunc("/ws/receiver",func(w http.ResponseWriter, r *http.Request) {
		ServeWs(true)(pool,w,r)
	})
	http.HandleFunc("/ws/sender",func(w http.ResponseWriter, r *http.Request) {
		ServeWs(false)(pool,w,r)
	})
}
