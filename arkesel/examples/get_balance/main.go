package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Replace with your actual key. This is just a dummy

	response, err := arkeselClient.GetBalance()
	if err != nil {
		log.Fatalf("could not get balance: %v", err)
	}
	fmt.Println("Balance is:", response["balance"])

}
