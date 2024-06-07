package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wuddup.com/models"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request){
	roomsJson, err := json.Marshal(models.AllRooms)
	if err != nil {
		fmt.Println("Error Marshalling JSON")
		return
	}

	w.Write(roomsJson)
}

func CreateNewRoom(w http.ResponseWriter, r *http.Request){
	var newRoom models.Room
	err :=json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		fmt.Println("error decoding JSON from request body")
		return
	}
	newRoom.AddRoomId()
	var clients []models.Client
	newRoom.Clients = clients
	fmt.Println("new room: ", newRoom)
	models.AllRooms = append(models.AllRooms, newRoom)
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(newRoom.Id))

}