package paintstructs

import "github.com/gorilla/websocket"

type Client struct {
	ClientId string          `json:"clientId"`
	Conn     *websocket.Conn `json:"-"`
}
