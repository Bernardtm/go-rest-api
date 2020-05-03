package main

import (
	"fmt"
	"go-rest-api/controllers"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	fmt.Printf("Starting server...\n\r")
	
	handler := cors.New(cors.Options{
		AllowedMethods:   []string { "GET", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		}).Handler(controllers.Include())
		
	port := ":9000"
	fmt.Printf("Server running: localhost%s\n\r", port)

	http.ListenAndServe(port, handler)

}
