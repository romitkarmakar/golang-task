package routes

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/",InfoRoute)
	http.HandleFunc("/ws",ServeWs)
}
