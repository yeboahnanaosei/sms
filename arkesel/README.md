![Arkesel's logo](logo.png)

## Contents
- [Contents](#contents)
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
	- [**Send SMS - Single**](#send-sms---single)
	- [**Send SMS - Bulk**](#send-sms---bulk)
	- [**Check balance**](#check-balance)
	- [**Schedule message**](#schedule-message)
	- [**Save contact**](#save-contact)
- [Contributing](#contributing)


## Introduction
[back-to-top](#contents)

This is wrapper around [Arkesel's](https://arkesel.com/) SMS API.

Arkesel provides platforms for engagement and personnalized communication using
SMS, Email, Voice and USSD. Explore our apps for free and step up your communication game.

## Installation
[back-to-top](#contents)

```bash
go get github.com/yeboahnanaosei/sms/arkesel
```


## Usage

### **Send SMS - Single**
[back-to-top](#contents)


```go
package main

import (
    "fmt"

    "github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
    // Create your Arkesel client or instance by passing your API key
	// obtained from your Arkesel account. Below is a dummy key.
	client := arkesel.New("ZG8gbm8gZXZpbA==")

	// Send a message to a single recipient.
    res, err := client.Message("Your message").From("SenderID").To("23324XXXXXXX").Send()
	if err != nil {
		log.Fatalf("sending sms failed: %v", err)
    }

	fmt.Println(res)
}
```

### **Send SMS - Bulk**
[back-to-top](#contents)


```go
package main

import (
    "fmt"

    "github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
    // Create your Arkesel client or instance by passing your API key
	// obtained from your Arkesel account. Below is a dummy key.
    client := arkesel.New("ZG8gbm8gZXZpbA==")

    // NOTE: Arkesel limits mulitple recipients to a maximum of 100 per message
    // To send a message to multiple recipients, you need to pass a comma separated
    // list of numbers to the To() method.
    recipients := "23324XXXXXXX,23326XXXXXXX,23320XXXXXXX"
    res, err := client.Message("Your message").From("SenderID").To(recipients).Send()
    if err != nil {
		log.Fatalf("sending sms failed: %v", err)
    }

	fmt.Println(res)
}
```


### **Check balance**
[back-to-top](#contents)


```go
package main

import (
    "fmt"

    "github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
    // Create your Arkesel client or instance by passing your API key
	// obtained from your Arkesel account. Below is a dummy key.
	client := arkesel.New("ZG8gbm8gZXZpbA==")

	// Send a message.
	res, err := client.GetBalance()
	if err != nil {
		log.Fatalf("checking balance failed: %v", err)
    }

	fmt.Println(res)
}
```


### **Schedule message**
[back-to-top](#contents)


```go
package main

import (
    "fmt"

    "github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
    // Create your Arkesel client or instance by passing your API key
	// obtained from your Arkesel account. Below is a dummy key.
	client := arkesel.New("ZG8gbm8gZXZpbA==")

	// Send a message.
	res, err := client.Message("Your message").From("SenderID").To("23324XXXXXXX").Schedule("26-01-2021 12:00 AM")
	if err != nil {
		log.Fatalf("schedulling sms failed: %v", err)
    }

	fmt.Println(res)
}
```


### **Save contact**
[back-to-top](#contents)


```go
package main

import (
	"fmt"
	"log"

	"github.com/yeboahnanaosei/sms/arkesel"
)

func main() {
	// Create your Arkesel client or instance by passing your API key
	// obtained from your Arkesel account. Below is a dummy key.
	client := arkesel.New("ZG8gbm8gZXZpbA==")

	// You can also save contacts to your Arkesel account. To do that you
    // first need to create a contact. Note that both PhoneBook and PhoneNumber
    // fields MUST be set. The rest of the fields are not required (they are optional).
	c := arkesel.Contact{
		PhoneBook:   "Family",            // Required. Should exist on your account
		PhoneNumber: "23324XXXXXXX",      // Required. Should be international format
		FirstName:   "FirstName",         // Optional
		LastName:    "LastName",          // Optional
		Email:       "person@email.com",  // Optional
		Company:     "Contact's Company", // Optional
	}
	res, err := client.SaveContact(c)
	if err != nil {
		log.Fatal("failed to save contact: ", err)
	}

	fmt.Println(res)
}
```

## Contributing
Suggestions, improvements, pull requests etc.. are all welcomed.
Feel free to get in touch on twitter: @yeboahnanaosei