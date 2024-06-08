package models

import "github.com/gorilla/websocket"

type Client struct {
	Id string `json:"id"`
	Ws *websocket.Conn
}

func (c *Client) BroadcastRooms(roomsInfo []byte){
	c.Ws.WriteMessage(1, roomsInfo)
}