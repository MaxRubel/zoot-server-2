package main

import (
	"fmt"
	"net/http"

	"wuddup.com/ws"
)

func main() {
	http.HandleFunc("/ws", ws.WsHandler)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
