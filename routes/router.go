package routes

import (
	"github.com/MaxRubel/zoot-server-2/ws"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ws", ws.WsHandler)
	r.HandleFunc("/rooms", GetAllRooms).Methods("GET")
	r.HandleFunc("/rooms", CreateNewRoom).Methods("POST")

	return r
}
