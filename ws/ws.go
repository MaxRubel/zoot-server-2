package ws

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MaxRubel/zoot-server-2/db"
	"github.com/MaxRubel/zoot-server-2/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var WaitingRoom models.Room

func init() {
	WaitingRoom.Clients = make(map[string]models.Client)
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error creating ws connection")
		return
	}
	defer conn.Close()

	//-----LISTENER-----//
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				fmt.Println("Client closed WebSocket connection and you did not catch it")
			}
			return
		}

		split := strings.Split(string(msg), "&")
		if len(split) < 4 {
			continue
		}

		msgType := split[0]
		roomId := split[1]
		senderId := split[2]
		recepientId := split[3]
		data := split[4]

		room, _ := models.AllRooms.FindRoom(roomId)
		db.IncrementWsCount()
		// debugging: display the incoming data:
		// utils.PrintIncomingWs(msgType, roomId, senderId, recepientId)

		//------FUNCTIONS------//
		switch msgType {

		case "0":
			fmt.Println("Received test message")
			room.BroadcastMessage("0Server received your message!")
		case "1":
			err := room.AddClient(senderId, conn)
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.WriteMessage(1, []byte("4&"+room.GetAllIds()))
			WaitingRoom.BroadcastRoomsUpdate()
		case "2":
			room.Negotiate(senderId, recepientId, data)
		case "3":
			// fmt.Println("leaving room")
			room.RemoveClient(senderId)
			conn.Close()
			WaitingRoom.BroadcastRoomsUpdate()
			return
		case "4":
			room.ClearClientArray()
		case "5":
			room.BroadcastMessage("0" + room.GetAllIds())
		case "6":
			WaitingRoom.AddClient(senderId, conn)
			conn.WriteMessage(1, []byte("7&"+string(models.AllRoomsJSON())))
			db.IncrementWsCount()
		case "7":
			conn.Close()
			WaitingRoom.RemoveClient(senderId)
		}
	}
}
