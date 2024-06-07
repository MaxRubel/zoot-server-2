package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"wuddup.com/routes"
)

func main() {
	r := routes.Router()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", corsHandler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
