package models

import "github.com/gorilla/websocket"

type Client struct {
	Id string `json:"id"`
	Ws *websocket.Conn
}