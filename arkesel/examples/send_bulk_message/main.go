package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Replace with your actual key. This is just a dummy

	// Send a bulk message to several recipients. To do this just pass the numbers
	// as a list of comma separated values. Please not that as at now the Arkesel
	// only allows you to send message to 100 recipients at a time
	res, err := arkeselClient.Message("bulk message").From("Nana").To("23324XXXXXXX,23326XXXXXXX").Send()
	if err != nil {
		log.Fatalf("sending sms failed: %v", err)
	}
	fmt.Println(res)
}
