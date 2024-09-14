package negotiations

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	paintstructs "github.com/MaxRubel/zoot-server-2/paintWS/structs"
	"github.com/gorilla/websocket"
)

type OutgoingMessage struct {
	Type string      `json:"type"`
	To   string      `json:"to"`
	From string      `json:"from"`
	Room string      `json:"room"`
	Data interface{} `json:"data"`
}

func retreiveClient(clientId string, roomId string) paintstructs.Client {
	room, err := paintstructs.GetRoom(roomId)

	if err != nil {
		log.Println("no room", err)
		return paintstructs.Client{}
	}

	client, ok := room.Clients[clientId]

	if !ok {
		log.Println("client not found in room", err)
		return paintstructs.Client{}
	}

	return client
}

func HandleOffer(msg paintstructs.IncomingMessage) {
	roomId := msg.Room
	senderId := msg.From
	recipient := msg.To
	offer := msg.Data

	client := retreiveClient(recipient, roomId)

	outMsg := OutgoingMessage{
		Type: "offer",
		To:   recipient,
		From: senderId,
		Room: roomId,
		Data: offer,
	}

	jsonMsg, err := json.Marshal(outMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func HandleAnswer(msg paintstructs.IncomingMessage) {
	roomId := msg.Room
	senderId := msg.From
	recipient := msg.To
	answer := msg.Data

	client := retreiveClient(recipient, roomId)

	outMsg := OutgoingMessage{
		Type: "answer",
		To:   recipient,
		From: senderId,
		Room: roomId,
		Data: answer,
	}

	jsonMsg, err := json.Marshal(outMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func HandleIceCandidate(msg paintstructs.IncomingMessage) {
	roomId := msg.Room
	senderId := msg.From
	recipient := msg.To
	iceCandidateData := msg.Data

	client := retreiveClient(recipient, roomId)

	outMsg := OutgoingMessage{
		Type: "iceCandidate",
		To:   recipient,
		From: senderId,
		Room: roomId,
		Data: iceCandidateData,
	}

	jsonMsg, err := json.Marshal(outMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func LeaveRoom(clientId string, roomId string) error {

	room := paintstructs.AllRooms[roomId]

	if room == nil {
		return errors.New("oops, room is nil")
	}
	room.RemoveClient(clientId)

	for _, client := range room.Clients {
		outMsg := OutgoingMessage{
			Type: "someoneLeft",
			To:   client.ClientId,
			From: clientId,
			Room: roomId,
			Data: nil,
		}

		jsonMsg, err := json.Marshal(outMsg)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return errors.New("error unmarshaling json")
		}
		client.Conn.WriteMessage(websocket.TextMessage, jsonMsg)
	}
	return nil
}

func BounceBack(conn *websocket.Conn) {
	outMsg := OutgoingMessage{
		Type: "bounceBack",
		To:   "",
		From: "",
		Room: "",
		Data: nil,
	}
	jsonMsg, err := json.Marshal(outMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, jsonMsg)
}

func SendInitialRoomData(msg paintstructs.IncomingMessage) {
	fmt.Println("got init room data")
	roomId := msg.Room
	senderId := msg.From
	recipient := msg.To
	newRoomData := msg.Data

	client := retreiveClient(recipient, roomId)

	outMsg := OutgoingMessage{
		Type: "initalJoinData",
		To:   client.ClientId,
		From: senderId,
		Room: roomId,
		Data: newRoomData,
	}

	jsonMsg, err := json.Marshal(outMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println("sending new room data message")
	client.Conn.WriteMessage(websocket.TextMessage, jsonMsg)

}
