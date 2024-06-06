package utils

import "fmt"

func PrintIncomingWs(tp string, senderId string, recepientId string){
	fmt.Println("type is: ", tp)
	fmt.Println("senderId is: ", senderId)
	fmt.Println("receiverId is: ", recepientId)
}