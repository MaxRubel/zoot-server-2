package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/MaxRubel/zoot-server-2/db"
	"github.com/gorilla/websocket"
)

type Room struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Clients map[string]Client `json:"clients"`
}

func (r *Room) AddClient(id string, conn *websocket.Conn) error {
	if r == nil {
		conn.WriteMessage(1, []byte("6&&&"))
		return errors.New("error this room does not exist")
	}

	if _, exists := r.Clients[id]; exists {
		return errors.New("this client is already in this room")
	}

	r.Clients[id] = Client{Ws: conn}
	return nil
}

func (r *Room) RemoveClient(id string) {
	if r == nil {
		return
	}
	delete(r.Clients, id)
	message := "5&" + id + "&&"
	r.BroadcastMessage(message)
	if len(r.Clients) == 0 {
		r.Delete()
	}
}

func (r *Room) BroadcastMessage(msg string) {
	if r == nil {
		return
	}
	for i := range r.Clients {
		r.Clients[i].Ws.WriteMessage(1, []byte(msg))
	}
	db.IncrementWsCount()
}

func (r *Room) ClearClientArray() {
	if r == nil {
		return
	}
	var newMap map[string]Client
	r.Clients = newMap
}

func (r *Room) Negotiate(senderId string, receiverId string, data string) {
	if r == nil {
		return
	}
	r.Clients[receiverId].Ws.WriteMessage(1, []byte("3"+"&"+r.Id+"&"+senderId+"&"+receiverId+"&"+data))
	db.IncrementWsCount()
}

func (r *Room) Delete() {
	if r == nil {
		return
	}
	delete(AllRooms, r.Id)
}

func (r *Room) GetAllIds() string {
	if r == nil {
		return ""
	}
	var strArr []string
	for key := range r.Clients {
		strArr = append(strArr, key)
	}
	flattenedStr := strings.Join(strArr, "&")
	return flattenedStr
}

func (r *Room) BroadcastRoomsUpdate() {
	if r == nil {
		return
	}
	msg := AllRoomsJSON()
	for i := range r.Clients {
		r.Clients[i].Ws.WriteMessage(1, []byte("7&"+string(msg)))
	}
	db.IncrementWsCount()
}

func AllRoomsJSON() []byte {
	roomsJson, err := json.Marshal(AllRooms)
	// fmt.Println("serving number of rooms: ", len(AllRooms))
	if err != nil {
		fmt.Println("Error Marshalling JSON")
		return nil
	}
	return roomsJson
}
