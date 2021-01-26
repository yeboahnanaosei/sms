package arkesel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// baseURL is the base URL for Arkesel's HTTP API
const baseURL string = "https://sms.arkesel.com"

// Sender IDs should not be longer than this value
const maxSenderIDLength int = 11

func (a *Arkesel) doBasicParameterChecks() error {
	if a.params.Get("from") == "" {
		return errors.New("no sender ID set. you can call From() to set it")
	}
	if len(a.params.Get("from")) > maxSenderIDLength {
		return fmt.Errorf("your sender ID is too long. maximum of %d characters allowed", maxSenderIDLength)
	}
	if a.params.Get("to") == "" {
		return errors.New("no recipient set. call the To method to set the recipient")
	}
	if a.params.Get("sms") == "" {
		return errors.New("no message body. call the Message method to set the body of the sms")
	}
	return nil
}

// doValidation performs the neccessary checks to make sure that a request can
// be made to Arkesel.
func (a *Arkesel) doValidation(action string) error {
	// The API key is required for all requests
	if a.params.Get("api_key") == "" {
		return errors.New("no api key set")
	}

	switch action {
	case "send-sms":
		return a.doBasicParameterChecks()
	case "check-balance", "subscribe-us":
		// do nothing
		return nil
	case "schedule":
		if err := a.doBasicParameterChecks(); err != nil {
			return err
		}
		if a.params.Get("schedule") == "" {
			return errors.New("no date and time set")
		}
	default:
		return errors.New("internal error. unknown endpoint")
	}
	return nil
}

// Send sends the message to the recipient.
func (a *Arkesel) Send() (map[string]interface{}, error) {
	a.params.Set("action", "send-sms")
	if err := a.doValidation("send-sms"); err != nil {
		return nil, fmt.Errorf("arkesel: sending failed: %v", err)
	}

	res, err := http.Get(fmt.Sprintf("%s/sms/api?%s", baseURL, a.params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("arkesel: failed to send sms: %v", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("arkesel: failed to read response from server: %v", err)
	}

	output := map[string]interface{}{}
	if err = json.Unmarshal(resBody, &output); err != nil {
		return nil, fmt.Errorf("arkesel: failed to read response from server: %v", err)
	}
	return output, nil
}

// GetBalance retrieves your balance on your account.
func (a *Arkesel) GetBalance() (map[string]interface{}, error) {
	a.params.Set("action", "check-balance")
	a.params.Set("response", "json")

	if err := a.doValidation("check-balance"); err != nil {
		return nil, fmt.Errorf("arkesel: cannot get balance: %v", err)
	}

	res, err := http.Get(fmt.Sprintf("%s/sms/api?%s", baseURL, a.params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("arkesel: cannot get balance: %v", err)
	}
	defer res.Body.Close()

	jsonBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("arkesel: could not read response from server: %v", err)
	}

	output := map[string]interface{}{}
	err = json.Unmarshal(jsonBody, &output)
	if err != nil {
		return nil, fmt.Errorf("arkesel: error unmarshalling response from server: %v", err)
	}
	return output, nil
}

// Schedule schedules a message to be delivered at the set date and time.
//
// This is the time the message is scheduled to be sent to the number(s).
// Set your schedule in the format dd-mm-yyyy hh:mm AM/PM
// Eg. 31-12-2021 12:30 PM
// This means 31st December 2021 at 12:30pm
//
// Example successful response looks like this:
//
//	map[string]interface{}{
//		"code": "ok",
//		"main_balance": 0.818,
//		"message": "SMS Scheduled successfully.",
//		"sms_balance": 336,
//		"user": "Nana Yeboah",
//	}
//
// Example error response looks like this:
//
// 	map[string]interface{}{
//		"code": 109,
//		"message": "Invalid Schedule Time"
//	}
func (a *Arkesel) Schedule(schedule string) (map[string]interface{}, error) {
	a.params.Set("action", "send-sms")
	a.params.Set("schedule", schedule)

	if err := a.doValidation("schedule"); err != nil {
		return nil, fmt.Errorf("arkesel: cannot schedule your message: %v", err)
	}

	res, err := http.Get(fmt.Sprintf("%s/sms/api?%s", baseURL, a.params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("arkesel: cannot get balance: %v", err)
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response from server %v", err)
	}

	output := map[string]interface{}{}
	if err := json.Unmarshal(responseBody, &output); err != nil {
		return nil, fmt.Errorf("could not unmarshal response from server: %v", err)
	}

	return output, nil
}

// Contact represents the contact details of a person.
//
// This is a contact you would like to save on your Arkesel account.
// The fields PhoneBook and PhoneNumber are required. The rest are optional.
type Contact struct {
	PhoneBook   string
	PhoneNumber string // Ensure the number is in the international format. Eg. 23324XXXXXXX
	FirstName   string
	LastName    string
	Email       string
	Company     string
}

func (c Contact) encode() string {
	return fmt.Sprintf("%s=%s&%s=%s&%s=%s&%s=%s&%s=%s&%s=%s",
		"phone_book",
		c.PhoneBook,
		"phone_number",
		c.PhoneNumber,
		"first_name",
		c.FirstName,
		"last_name",
		c.LastName,
		"email",
		c.Email,
		"company",
		c.Company,
	)
}

// SaveContact is used to save customer phone numbers or personal contacts to a
// contact group.
//
// The contact to be saved must have the fields PhoneBook and PhoneNumber set.
// These two are required. The rest are optional.
//
// NOTE: The contact group must be present on your account in order to add contact details
func (a *Arkesel) SaveContact(c Contact) (map[string]interface{}, error) {
	a.params.Set("action", "subscribe-us")

	if err := a.doValidation("subscribe-us"); err != nil {
		return nil, fmt.Errorf("arkesel: cannot save contact: %v", err)
	}
	if c.PhoneBook == "" || c.PhoneNumber == "" {
		return nil, errors.New("arkesel: cannot save contact. missing phone book or phone number")
	}

	res, err := http.Get(fmt.Sprintf("%s/contacts/api?%s&%s", baseURL, a.params.Encode(), c.encode()))
	if err != nil {
		return nil, fmt.Errorf("arkesel: saving contact failed: %v", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("arkesel: could not read resposne from server %v", err)
	}

	output := map[string]interface{}{}
	err = json.Unmarshal(resBody, &output)
	if err != nil {
		return nil, fmt.Errorf("arkesel: could not unmarshal response from server: %v", err)
	}
	return output, nil
}
