package paintWs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/MaxRubel/zoot-server-2/paintWS/negotiations"
	paintstructs "github.com/MaxRubel/zoot-server-2/paintWS/structs"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("receiving new websocket connect")
	// Parse initial data from URL
	initialData, err := parseInitialData(r, conn)
	if err != nil {
		log.Println("Error parsing initial data:", err)
		return
	}

	log.Printf("New connection: Host: %v, RoomId: %s, ClientId: %s", initialData.Host, initialData.RoomId, initialData.ClientId)

	room, err := paintstructs.GetRoom(initialData.RoomId)
	if err != nil {
		negotiations.BounceBack(conn)
	}

	paintstructs.SendRoomAsJSON(conn, room)

	cleanup := func() {
		log.Printf("Connection closed for client %s in room %s", initialData.ClientId, initialData.RoomId)
		room, err := paintstructs.GetRoom(initialData.RoomId)
		negotiations.LeaveRoom(initialData.ClientId, initialData.RoomId)

		if err != nil {
			fmt.Println(err)
			return
		}

		if room != nil {
			err := room.RemoveClient(initialData.ClientId)
			if err != nil {
				log.Printf("Error removing client from room: %v", err)
			}
			if len(room.Clients) == 0 {
				paintstructs.DeleteRoom(initialData.RoomId)
			}
		}
		conn.Close()
	}

	defer cleanup()

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var message paintstructs.IncomingMessage
		if err := json.Unmarshal(p, &message); err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		// Handle the message based on its type
		switch message.Type {
		case "offer":
			negotiations.HandleOffer(message)
		case "answer":
			negotiations.HandleAnswer(message)
		case "iceCandidate":
			negotiations.HandleIceCandidate(message)
		case "initalJoinData":
			negotiations.SendInitialRoomData(message)
		default:
			log.Printf("Received unknown message type: %s", message.Type)
		}
	}

}

func parseInitialData(r *http.Request, conn *websocket.Conn) (paintstructs.InitialData, error) {
	var initialData paintstructs.InitialData

	encodedData := r.URL.Query().Get("data")
	if encodedData == "" {
		return initialData, fmt.Errorf("no initial data provided")
	}

	decodedData, err := url.QueryUnescape(encodedData)
	if err != nil {
		return initialData, fmt.Errorf("error decoding data: %v", err)
	}

	err = json.Unmarshal([]byte(decodedData), &initialData)
	if err != nil {
		return initialData, fmt.Errorf("error parsing JSON: %v", err)
	}

	if initialData.Host {
		paintstructs.AddRoom(initialData.RoomId)
	}
	room := paintstructs.AllRooms[initialData.RoomId]

	room.AddClient(initialData.ClientId, conn)

	return initialData, nil
}
