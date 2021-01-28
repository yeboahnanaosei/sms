package fayasms

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var endpoints = map[string]string{
	"send":     "https://devapi.fayasms.com/send",
	"messages": "https://devapi.fayasms.com/messages",
	"balance":  "https://devapi.fayasms.com/balance",
	"estimate": "https://devapi.fayasms.com/estimate",
	"senders":  "https://devapi.fayasms.com/senders",
	"new_id":   "https://devapi.fayasms.com/senders/new",
}

// exec is what does the actual work. It executes the actual http request by
// fetching the endpoint and making a request to that endpoint. It also
// performs the checks required before making the request.
func (f *FayaSMS) exec(endpoint string) (response string, err error) {
	// Fetch the endpoint we need to make the request to
	endpnt, ok := endpoints[endpoint]

	if !ok {
		return response, errors.New("fayasms: unknown endpoint targetted")
	}

	// Check if the mandatory fields have been set for this FayaSMS instance.
	// The mandatory fields are required for all endpoints
	if err = f.checkMandatoryFields(mandatoryFields); err != nil {
		return response, err
	}

	// Now check the fields required for the particular endpoint only
	if err = f.checkConditionalFields(endpoint, conditionalFields); err != nil {
		return response, err
	}

	req, err := http.NewRequest(http.MethodPost, endpnt, strings.NewReader(f.payload.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return response, err
	}

	if f.ctx == nil {
		f.ctx = context.Background()
	}
	req = req.WithContext(f.ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	return string(data), nil
}

// Send sends the message to the recipient or recipients you've set
func (f *FayaSMS) Send() (response string, err error) {
	return f.exec("send")
}

// GetEstimate lets you know the number of units it will cost you to send the message.
// This is determined by length of your message body and the number of recipients.
func (f *FayaSMS) GetEstimate() (response string, err error) {
	return f.exec("estimate")
}

// GetBalance returns your current balance on FayaSMS
func (f *FayaSMS) GetBalance() (response string, err error) {
	return f.exec("balance")
}

// RequestSenderID makes a request to FayaSMS for a new sender id
// senderID is the sender id you are requesting for.
// desc is a description of the sender id. What will you use the id for.
// The description is used in the approval process
func (f *FayaSMS) RequestSenderID(senderID, desc string) (response string, err error) {
	f.payload.Set("Name", senderID)
	f.payload.Set("Description", desc)
	return f.exec("new_id")
}

// RetrieveMessages returns all the messages you've sent using your AppKey and AppSecret
func (f *FayaSMS) RetrieveMessages() (response string, err error) {
	return f.exec("messages")
}

// RetrieveMessage retrieves a particular message you've sent whose id is messageID
func (f *FayaSMS) RetrieveMessage(messageID string) (response string, err error) {
	f.payload.Set("MessageId", messageID)
	f.extra = true
	return f.exec("messages")
}
