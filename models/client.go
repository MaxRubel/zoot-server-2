package models

import "github.com/gorilla/websocket"

type Client struct {
	Id string
	Ws *websocket.Conn
}