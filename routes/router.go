package routes

import (
	"github.com/gorilla/mux"
	"wuddup.com/ws"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	
	r.HandleFunc("/ws", ws.WsHandler)
	
	r.HandleFunc("/rooms", GetAllRooms).Methods("GET")
	r.HandleFunc("/rooms", CreateNewRoom).Methods("POST")

	return r
}