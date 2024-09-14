package paintstructs

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	RoomId  string
	Clients map[string]Client
	mu      sync.RWMutex
}

type ClientInfo struct {
	ClientId string `json:"clientId"`
}

func (r *Room) AddClient(clientId string, ws *websocket.Conn) error {
	if r == nil {
		return errors.New("nil pointer, room is not defined or accesible")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.Clients[clientId] = Client{
		ClientId: clientId,
		Conn:     ws,
	}

	// fmt.Printf("added client %s to room %s \n", clientId, r.RoomId)
	return nil
}

func (r *Room) RemoveClient(clientId string) error {
	if r == nil {
		return errors.New("nil pointer, room is not defined or accesible")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.Clients, clientId)

	// fmt.Printf("removed client %s from room %s \n", clientId, r.RoomId)
	return nil
}

func SendRoomAsJSON(conn *websocket.Conn, room *Room) error {

	if room == nil {
		return errors.New("no room found to get json data from")
	}

	room.mu.RLock()
	defer room.mu.RUnlock()

	roomCopy := struct {
		RoomId  string   `json:"roomId"`
		Clients []string `json:"clientIds"`
	}{
		RoomId:  room.RoomId,
		Clients: make([]string, 0, len(room.Clients)),
	}

	var array []string

	for _, client := range room.Clients {
		array = append(array, client.ClientId)
	}

	roomCopy.Clients = array

	jsonData, err := json.Marshal(roomCopy)
	if err != nil {
		log.Println("Error marshaling room to JSON:", err)
		return err
	}

	message := struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}{
		Type: "new_client_ids",
		Data: jsonData,
	}

	// Marshal the entire message
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling message to JSON:", err)
		return err
	}

	// Send the JSON data through the WebSocket connection
	err = conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		log.Println("Error sending room data:", err)
		return err
	}

	return nil
}
