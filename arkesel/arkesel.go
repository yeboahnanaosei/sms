// Package arkesel is a wrapper around Arkesel.com's sms api. With this package
// you can directly interact with Arkesel's platform directly from your go
// application
package arkesel

import (
	"net/url"
)

// Arkesel models one request to be sent to the ark
type Arkesel struct {
	params url.Values
}

// New returns an instance of Arkesel with which you can make request to the
// Arkesel platform
func New(APIKey string) *Arkesel {
	a := Arkesel{params: url.Values{}}
	a.params.Set("api_key", APIKey)
	return &a
}

// Message sets the message to be sent
//
// 1 paged message = 160 character, so you send a message with 200 characters,
// the messages will be 2 pages.
func (a *Arkesel) Message(msg string) *Arkesel {
	a.params.Set("sms", msg)
	return a
}

// From sets the senderID of the message.
//
// A Sender ID (from address) is the name or number that identifies the sender
// of an SMS message. Note that this field should be 11 characters max including
// space. Anything more than that will result in your messages failing.
func (a *Arkesel) From(senderID string) *Arkesel {
	a.params.Set("from", senderID)
	return a
}

// To sets the phone number of the recipient(s) of the message. This can be
// the phone number of one recipient or a comma separated list of several
// recipients. Ensure the number is in international format eg. 23324XXXXXXX
//
// Examples:
// One recipient: '23324XXXXXXX'.
// Multiple recipients: '23324XXXXXXX,233020XXXXXXX,233026XXXXXXXX'.
func (a *Arkesel) To(phone string) *Arkesel {
	a.params.Set("to", phone)
	return a
}

// SetAPIKey allows you to set a different api key
func (a *Arkesel) SetAPIKey(key string) *Arkesel {
	a.params.Set("api_key", key)
	return a
}
