package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Fake API key

	// Send a message.
	res, err := arkeselClient.Message("bulk message").From("Nana").To("23324XXXXXXX,23326XXXXXXX").Send()
	if err != nil {
		log.Fatalf("sending sms failed: %v", err)
	}
	fmt.Println(res)


	// You can schedule a message to be sent at a later date
	res, err = arkeselClient.Message("some message").From("Nana").To("23324XXXXXXX").Schedule("26-01-2021 03:29 AM")
	if err != nil {
		log.Fatal("could not schedule your message: ", err)
	}
	fmt.Println(res)


	
	// Check account balance
	balance, err := arkeselClient.GetBalance()
	if err != nil {
		log.Fatalf("could not get balance: %v", err)
	}
	fmt.Println("Balance is:", balance)



	// You can also save contacts to your Arkesel account. To do that you
	// first need to create a contact and then save that contact. Note that
	// both PhoneBook and PhoneNumber fields MUST be set. The rest of the
	// fields are not required (they are optional).
	c := arkesel.Contact{
		PhoneBook:   "Family",            // Required. Should exist on your account
		PhoneNumber: "23324XXXXXXX",      // Required. Should be international format
		FirstName:   "FirstName",         // Optional
		LastName:    "LastName",          // Optional
		Email:       "person@email.com",  // Optional
		Company:     "Contact's Company", // Optional
	}
	res, err = arkeselClient.SaveContact(c)
	if err != nil {
		log.Fatal("failed to save contact: ", err)
	}

	fmt.Println(res)
}
