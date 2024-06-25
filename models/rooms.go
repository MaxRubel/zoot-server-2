package models

import (
	"errors"
)

type Rooms map[string]Room

var AllRooms = make(Rooms)

func (r Rooms) FindRoom(id string) (*Room, error) {
	if room, exists := r[id]; exists {
		return &room, nil
	}
	return nil, errors.New("error: no room found")
}

// func (r Rooms) RemoveClient(id string) {
// 	fmt.Println("deleting client")
// 	for _, room := range r {
// 		for clientID, client := range room.Clients {
// 			if client.Id == id {
// 				delete(room.Clients, clientID)
// 				return
// 			}
// 		}
// 	}
// }
