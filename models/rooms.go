package models

import (
	"errors"
)

type Rooms []Room

var AllRooms Rooms

func (r *Rooms) FindRoom(id string) (*Room, error) {
	for i := range *r {
		if (*r)[i].Id == id {
			return &(*r)[i], nil
		}
	}
	return nil, errors.New("error-no room found")
}

