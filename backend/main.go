package main

import (
	"fmt"
	"net/http"

	"github.com/sswastik02/PublicRoom/routes"
)



func main() {
	routes.SetupRoutes()
	fmt.Println("Starting Broadcasting Server")
	err := http.ListenAndServe(":8000",nil)

	if(err != nil){
		fmt.Printf("\nError : %v",err)
	}
}
