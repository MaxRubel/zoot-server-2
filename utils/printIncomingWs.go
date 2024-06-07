package utils

import "fmt"

func PrintIncomingWs(
	tp string,
	roomID string,
	senderId string,
	recepientId string) {
	fmt.Println("type is: ", tp)
	fmt.Println("room id is: ", roomID)
	fmt.Println("senderId is: ", senderId)
	fmt.Println("receiverId is: ", recepientId)
}
