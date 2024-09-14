package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	paintWs "github.com/MaxRubel/zoot-server-2/paintWS"
	"github.com/MaxRubel/zoot-server-2/routes"
	"github.com/MaxRubel/zoot-server-2/ws"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ws", ws.WsHandler)
	r.HandleFunc("/rooms", routes.GetAllRooms).Methods("GET")
	r.HandleFunc("/rooms", routes.CreateNewRoom).Methods("POST")
	r.HandleFunc("/paintws", paintWs.HandleWebSocket)

	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	origin1 := os.Getenv("ORIGIN1")
	origin2 := os.Getenv("ORIGIN2")

	r := Router()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{origin1, origin2}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)

	fmt.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":"+port, corsHandler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
