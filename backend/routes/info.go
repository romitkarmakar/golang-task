package routes

import (
	"fmt"
	"net/http"
)

func InfoRoute(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"A Simple Broadcasting Service built using Go")
}
