package paintstructs

import (
	"errors"
	"sync"
)

var (
	AllRooms  map[string]*Room
	roomMutex sync.RWMutex
)

func init() {
	AllRooms = make(map[string]*Room)
}

func AddRoom(roomId string) *Room {

	roomMutex.Lock()
	defer roomMutex.Unlock()

	newRoom := Room{
		RoomId:  roomId,
		Clients: make(map[string]Client),
	}

	AllRooms[roomId] = &newRoom
	return &newRoom
}

func GetRoom(roomId string) (*Room, error) {
	if AllRooms[roomId] == nil {
		return nil, errors.New("error getting room, that room doesn't exist or is nil")
	}
	return AllRooms[roomId], nil
}

func DeleteRoom(roomId string) error {
	roomMutex.Lock()
	defer roomMutex.Unlock()
	delete(AllRooms, roomId)

	return nil
}
