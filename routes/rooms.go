package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wuddup.com/models"
	"wuddup.com/ws"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	roomsJson, err := json.Marshal(models.AllRooms)
	fmt.Println("serving number of rooms: ", len(models.AllRooms))
	if err != nil {
		fmt.Println("Error Marshalling JSON")
		return
	}

	w.Write(roomsJson)
}

func CreateNewRoom(w http.ResponseWriter, r *http.Request) {
	var newRoom models.Room
	err := json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		fmt.Println("error decoding JSON from request body")
		return
	}

	id := newRoom.Create()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
	ws.WaitingRoom.BroadcastRoomsUpdate()
}
