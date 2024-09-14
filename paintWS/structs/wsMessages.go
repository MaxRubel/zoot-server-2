package paintstructs

type InitialData struct {
	Host     bool   `json:"host"`
	RoomId   string `json:"roomId"`
	ClientId string `json:"clientId"`
}

type OutgoingMessage struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    InitialData `json:"data"`
}

type IncomingMessage struct {
	Type string      `json:"type"`
	To   string      `json:"to"`
	From string      `json:"from"`
	Room string      `json:"room"`
	Data interface{} `json:"data"`
}
