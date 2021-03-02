package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Replace with your actual key. This is just a dummy

	// You can schedule a message to be sent at a later date. The format for the
	// date is as follows dd-mm-yyyy
	res, err := arkeselClient.Message("some message").From("Nana").To("23324XXXXXXX").Schedule("26-01-2021 03:29 AM")
	if err != nil {
		log.Fatal("could not schedule your message: ", err)
	}
	fmt.Println(res)

}
