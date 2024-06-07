package ws

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"wuddup.com/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	//LISTENER
	for {
		_, msg, err := conn.ReadMessage()
		if websocket.IsCloseError(err, websocket.CloseGoingAway) {
			fmt.Println("Closing WebSocket connection")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("Error writing close message:", err)
			}
			break
		}
		if err != nil {
			continue
		}

		conv := string(msg)
		split := strings.Split(conv, "&")
		if len(split) < 4 {
			continue
		}
		tp := split[0]
		roomId := split[1]
		senderId := split[2]
		recepientId := split[3]
		data := split[4]
		room, _ := models.AllRooms.FindRoom(roomId)

		// utils.PrintIncomingWs(tp, roomId, senderId, recepientId)

		switch tp {

		case "0":
			fmt.Println("Received test message")
			room.BroadcastMessage("0Server received your message!")

		case "1":
			room.AddToRoom(senderId, conn)
			clientIdString := room.FlattenArray()
			conn.WriteMessage(1, []byte("4&"+clientIdString))

		case "2":
			err := room.Negotiate(senderId, roomId, recepientId, data)
			if err != nil {
				fmt.Println(err)
				continue
			}

		case "3":
			fmt.Println("leaving room")
			clientsLeft, err := room.RemoveClient(senderId)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if clientsLeft == 0 {
				fmt.Println("room is empty, deleting")
				room, err := models.AllRooms.FindRoom(roomId)
				if err != nil {
					fmt.Println(err)
				}
				room.Delete()
			}

			conn.Close()
			return

		case "4":
			room.ClearClientArray()

		case "5":
			msg := room.FlattenArray()
			room.BroadcastMessage("0" + msg)
		}
	}
}
