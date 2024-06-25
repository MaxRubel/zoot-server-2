package utils

import (
	"fmt"

	"github.com/MaxRubel/zoot-server-2/models"
)

func PingEveryone(msg string) {
	// msg := models.AllRoomsJSON()
	for i := range models.AllRooms {
		clients := models.AllRooms[i].Clients
		for x := range clients {
			fmt.Println("sending message to: ", clients[x].Id)
			err := clients[x].Ws.WriteMessage(1, []byte(msg))
			if err != nil {
				fmt.Printf("Error sending message to client: %v\n", err)
			}
		}
	}
}
