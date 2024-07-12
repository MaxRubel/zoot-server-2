package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaxRubel/zoot-server-2/models"
	"github.com/MaxRubel/zoot-server-2/utils"
	"github.com/MaxRubel/zoot-server-2/ws"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all da rooms")
	roomsJson, err := json.Marshal(models.AllRooms)
	// fmt.Println("serving number of rooms: ", len(models.AllRooms))
	if err != nil {
		fmt.Println("Error Marshalling JSON")
		return
	}

	w.Write(roomsJson)
}

func CreateNewRoom(w http.ResponseWriter, r *http.Request) {
	var newRoom models.Room
	id := utils.CreateId()

	newRoom.Clients = make(map[string]models.Client)
	newRoom.Id = id

	err := json.NewDecoder(r.Body).Decode(&newRoom)
	if err != nil {
		fmt.Println("error decoding JSON from request body")
		return
	}

	models.AllRooms[id] = newRoom

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
	ws.WaitingRoom.BroadcastRoomsUpdate()
}
