package models

type Rooms []Room
var AllRooms Rooms

func (r *Rooms)AddRoom(room Room){
	*r = append(*r, room)
}

func (r *Rooms) DeleteRoomByID(roomID string) {
    for i, room := range *r {
        if room.Id == roomID {
            *r = append((*r)[:i], (*r)[i+1:]...)
            return
        }
    }
}


