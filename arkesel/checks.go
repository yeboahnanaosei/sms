package arkesel

import (
	"errors"
	"net/url"
)

// runChecks performs some basic checks to make sure that a request can be
// made to Arkesel
func (a *Arkesel) runChecks(action string) error {
	// For whatever reason, make sure the API key is always set
	if a.params.Get("api_key") == "" {
		return errors.New("no api key set")
	}

	switch action {
	case "send-sms":
		return sendSMSCheck(a.params)
	case "check-balance":
	case "subscribe-us":
		if err := sendSMSCheck(a.params); err != nil {
			return err
		}

		if a.params.Get("schedule") == "" {
			return errors.New("no date and time set")
		}

	}
	return nil
}

func sendSMSCheck(params url.Values) error {
	if params.Get("from") == "" {
		return errors.New("no api key set. you can call SetAPIKey to set it")
	}

	if params.Get("to") == "" {
		return errors.New("no recipient set. call the To method to set the recipient")
	}

	if params.Get("sms") == "" {
		return errors.New("no message body. call the Message method to set the body of the sms")
	}
	return nil
}
