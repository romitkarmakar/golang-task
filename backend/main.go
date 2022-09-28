package main

import (
	"fmt"
	"net/http"

	"github.com/sswastik02/PublicRoom/routes"
)



func main() {
	routes.SetupRoutes()
	fmt.Println("Starting Broadcasting Server")
	http.ListenAndServe(":8000",nil)
}