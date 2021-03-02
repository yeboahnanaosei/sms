package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Replace with your actual key. This is just a dummy

	// Send a message to a single recipient
	res, err := arkeselClient.From("SenderID").To("23324XXXXXXX").Message("Message to send").Send()
	if err != nil {
		log.Fatalf("sending sms failed: %v", err)
	}
	fmt.Println(res)
}
