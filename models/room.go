package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Room struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Clients []Client `json:"clients"`
}

func (r *Room) Create() string {
	r.Clients = []Client{}
	r.AddRoomId()
	AllRooms = append(AllRooms, *r)
	return r.Id
}

func (r *Room) AddClient(id string, conn *websocket.Conn) error {
	if r == nil {
		return errors.New("error this room does not exist")
	}
	for i := range r.Clients {
		if r.Clients[i].Id == id {
			return errors.New("error: client already in room")
		}
	}
	r.Clients = append(r.Clients, Client{
		Id: id,
		Ws: conn,
	})
	fmt.Println("client added to room")
	fmt.Println("number of clients in room: ", len(r.Clients))
	return nil
}

func (r *Room) RemoveClient(id string) {
	for i := range r.Clients {
		if r.Clients[i].Id == id {
			r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
			fmt.Println("removed client from room.  ", len(r.Clients), "  left in room")
			break
		}
	}
	message := "5&" + id + "&&"
	r.BroadcastMessage(message)
	if len(r.Clients) == 0 {
		fmt.Println("room is empty, deleting")
		r.Delete()
	}
}

func (r *Room) BroadcastMessage(msg string) {
	for i := range r.Clients {
		r.Clients[i].Ws.WriteMessage(1, []byte(msg))
	}
}

func (r *Room) ClearClientArray() {
	var newArray []Client
	r.Clients = newArray
	fmt.Println("cleared client array", newArray)
}

func (r *Room) Negotiate(senderId string, receiverId string, data string) {
	fmt.Println("negotiating...")
	r.BroadcastMessage("3" + "&" + r.Id + "&" + senderId + "&" + receiverId + "&" + data)
}

func (r *Room) Delete() {
	for i := range AllRooms {
		if AllRooms[i].Id == r.Id {
			AllRooms = append(AllRooms[:i], AllRooms[i+1:]...)
			return
		}
	}
}

func (r *Room) GetAllIds() string {
	var strArr []string

	for i := range r.Clients {
		strArr = append(strArr, r.Clients[i].Id)
	}
	flattenedStr := strings.Join(strArr, "&")
	return flattenedStr
}

func (r *Room) AddRoomId() {
	r.Id = uuid.New().String()
}
