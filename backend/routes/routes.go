package routes

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/",InfoRoute)
}
