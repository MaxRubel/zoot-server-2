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

func (r *Room) AddToRoom(id string, conn *websocket.Conn) error {
	for i := range r.Clients {
		if r.Clients[i].Id == id {
			return errors.New("client already in room")
		}
	}
	r.Clients = append(r.Clients, Client{
		Id: id,
		Ws: conn,
	})
	fmt.Println("client added to room")
	fmt.Println("clients in room: ", r.Clients)
	return nil
}

func (r *Room) RemoveClient(id string) (int, error) {
	for i := range r.Clients {
		if r.Clients[i].Id == id {
			fmt.Println("removing client from room ")
			r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
			return len(r.Clients), nil
		}
		message := "5&" + id + "&&"
		r.BroadcastMessage(message)
	}
	return 0, fmt.Errorf("client with ID %s not found in the room", id)
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

func (r *Room) Negotiate(senderId string, roomID, receiverId string, data string) error {
	fmt.Println("negotiating...")
	r.BroadcastMessage("3" + "&" + roomID + "&" + senderId + "&" + receiverId + "&" + data)
	return nil
}

func (r *Room) Delete() {
	for i := range AllRooms {
		if AllRooms[i].Id == r.Id {
			AllRooms = append(AllRooms[:i], AllRooms[i+1:]...)
			return
		}
	}
}

func (r *Room) FlattenArray() string {
	var strArr []string

	for i := range r.Clients {
		fmt.Println("client id:", r.Clients[i].Id)
		strArr = append(strArr, r.Clients[i].Id)
	}
	flattenedStr := strings.Join(strArr, "&")
	fmt.Println("clients in room: ", flattenedStr)
	return flattenedStr
}

func (r *Room) AddRoomId() {
	r.Id = uuid.New().String()
}
