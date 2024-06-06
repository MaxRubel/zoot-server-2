package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
)

type Room struct {
	Clients []Client
	HostJoined bool
}

func (r *Room) AddToRoom (id string, conn *websocket.Conn) error {
	if len(r.Clients) == 0 {
		r.HostJoined = true
	}
	for i := range r.Clients{
		if r.Clients[i].Id == id {
			return errors.New("client already in room")
		}
	}

	r.Clients = append(r.Clients, Client{
		Id: id,
		Ws: conn,
	})

	fmt.Println("client added to room")
	fmt.Println("host has joined:", r.HostJoined)
	fmt.Println("clients in room: ", r.Clients)
	return nil
}

func (r *Room) RemoveFromRoom(id string) error {
    for i := range r.Clients {
        if r.Clients[i].Id == id {
            r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("client with ID %s not found in the room", id)
}

func(r *Room) BroadcastMessage(msg string) {
	for i := range r.Clients {
		r.Clients[i].Ws.WriteMessage(1, []byte(msg))
	}
}

func (r *Room) ClearClientArray(){
	var newArray []Client
	r.Clients = newArray
	fmt.Println("cleared client array", newArray)
}

func (r *Room) Negotiate(senderId string, receiverId string, data string) error {
	fmt.Println("negotiating...")

	r.BroadcastMessage("3" + "&" + senderId + "&" + receiverId + "&" + data)
	return nil
}

func (r *Room) FlattenArray() (string, error) {
    var strArr []string

	for i := range r.Clients{
		fmt.Println("client id:", r.Clients[i].Id)
		strArr = append(strArr, r.Clients[i].Id)
	}

    flattenedStr := strings.Join(strArr, "&")

	fmt.Println( "clients in room: ", flattenedStr)

    return flattenedStr, nil
}