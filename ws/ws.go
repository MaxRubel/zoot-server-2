package ws

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"wuddup.com/models"
)

func writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var room models.Room

func WsHandler(w http.ResponseWriter, r *http.Request) {
	writeCORS(w)
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	defer conn.Close()

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
			// fmt.Println("error reading incoming message:", err)
			continue
		}

		conv := string(msg)

		split := strings.Split(conv, "&")

		if len(split) < 4 {
			continue
		}

		tp := split[0]
		senderId := split[1]
		recepientId := split[2]
		data := split[3]

		// utils.PrintIncomingWs(tp, senderId, recepientId)

		switch tp {
		case "0":
			fmt.Println("Received test message")
			room.BroadcastMessage("0Server received your message!")
		case "1":
			room.AddToRoom(senderId, conn)
			clientIdString, err := room.FlattenArray()
			if err != nil {
				fmt.Println("error making string array of client ids")
			}
			conn.WriteMessage(1, []byte("4&"+clientIdString))
		case "2":
			err := room.Negotiate(senderId, recepientId, data)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "3":
			fmt.Println("leaving room")
			err = room.RemoveFromRoom(senderId)
			if err !=nil {
				fmt.Println(err)
				continue
			}
			conn.Close()
			return
		case "4":
			room.ClearClientArray()
		case "5":
			msg, err := room.FlattenArray()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("client return: ", msg)
			room.BroadcastMessage("0"+msg)
		}
	}
}
