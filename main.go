package main

import (
	"fmt"
	"net/http"

	"github.com/MaxRubel/zoot-server-2/routes"
	"github.com/MaxRubel/zoot-server-2/ws"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ws", ws.WsHandler)
	r.HandleFunc("/rooms", routes.GetAllRooms).Methods("GET")
	r.HandleFunc("/create_room", routes.CreateNewRoom).Methods("POST")

	return r
}

func main() {
	r := Router()

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
