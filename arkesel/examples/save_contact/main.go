package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your api key
	arkeselClient := arkesel.New("ZG8gbm8gZXZpbA==") // <- Replace with your actual key. This is just a dummy

	// You can also save contacts to your Arkesel account. To do that you
	// first need to create a contact and then save that contact. Note that
	// both PhoneBook and PhoneNumber fields MUST be set. The rest of the
	// fields are not required (they are optional).
	contact := arkesel.Contact{
		PhoneBook:   "Family",            // Required. Should exist on your account
		PhoneNumber: "23324XXXXXXX",      // Required. Should be international format
		FirstName:   "FirstName",         // Optional
		LastName:    "LastName",          // Optional
		Email:       "person@email.com",  // Optional
		Company:     "Contact's Company", // Optional
	}
	res, err := arkeselClient.SaveContact(contact)
	if err != nil {
		log.Fatal("failed to save contact: ", err)
	}

	fmt.Println(res)
}
