package arkesel

import (
	"testing"
)

func TestDoBasicParameterChecks(t *testing.T) {
	aksel := New("apikey")
	if err := aksel.doBasicParameterChecks(); aksel.params.Get("from") == "" && err == nil {
		t.Logf("failed expected %v but got %v", "no sender id set", err)
	}

	aksel.From("veryverylongsenderid")
	if err := aksel.doBasicParameterChecks(); len(aksel.params.Get("from")) > maxSenderIDLength && err == nil {
		t.Logf("failed expected %v but got %v", "sender id too long", err)
	}

	aksel.From("somesender")
	if err := aksel.doBasicParameterChecks(); aksel.params.Get("to") == "" && err == nil {
		t.Logf("failed. expected %v but got %v", "no recipient set", err)
	}

	aksel.To("024XXXXXXX")
	if err := aksel.doBasicParameterChecks(); aksel.params.Get("sms") == "" && err == nil {
		t.Logf("failed. expected %v but got %v", "no message body", err)
	}

	aksel.Message("some message")
	if err := aksel.doBasicParameterChecks(); err != nil {
		t.Logf("failed. expected %v but got %v", nil, err)
	}
}

func TestDoValidation(t *testing.T) {
	a := New("")
	if err := a.doValidation("send-sms"); err == nil {
		t.Logf("failed: expected %v but got %v", "no api key set", err)
	}

	a.SetAPIKey("apikey").From("senderid").To("024XXXXXXX")
	if err := a.doValidation("send-sms"); err == nil {
		t.Logf("failed: expected %v but got %v", "no message body set", nil)
	}

	a.SetAPIKey("newapikey").To("024XXXXXXX").Message("some message")
	if err := a.doValidation("send-sms"); err == nil {
		t.Logf("failed: expected %v but got %v", "no sender id set", nil)
	}

	a.SetAPIKey("newapikey").From("senderid").Message("some message")
	if err := a.doValidation("send-sms"); err == nil {
		t.Logf("failed: expected %v but got %v", "no recipient set", nil)
	}

	if err := a.doValidation("check-balance"); err != nil {
		t.Logf("validating `check-balance` failed: expected %v but got %v", nil, err)
	}

	if err := a.doValidation("subscribe-us"); err != nil {
		t.Logf("validating `subscribe-us` failed: expected %v but got %v", nil, err)
	}

	a.SetAPIKey("apikey").From("senderid").To("024XXXXXXX").From("senderid").Message("some message")
	if err := a.doValidation("schedule"); err == nil {
		t.Logf("failed: expected %v but got %v", "no date and time set", err)
	}

	a.params.Set("schedule", "26-01-2021 03:23 AM")
	a.SetAPIKey("apikey").From("senderid").To("024XXXXXXX").Message("some message")
	if err := a.doValidation("schedule"); err != nil {
		t.Logf("failed: expected %v but got %v", nil, err)
	}

	if err := a.doValidation("unknown endpoint"); err == nil {
		t.Logf("validation failed: expected %v but got %v", "unknown endpoint", nil)
	}
}

func TestSend(t *testing.T) {
	a := New("apikey")
	a.From("senderid").To("024XXXXXXX").Message("message to send")
	_, err := a.Send()
	if err != nil {
		t.Errorf("test failed: expected %v but got %v", nil, err)
	}
}

func TestEncode(t *testing.T) {
	tt := []struct {
		c   Contact
		out string
	}{
		{
			Contact{
				PhoneBook:   "friends",
				PhoneNumber: "024XXXXXXX",
			},
			"phone_book=friends&phone_number=024XXXXXXX&first_name=&last_name=&email=&company=",
		},
		{
			Contact{
				PhoneBook:   "",
				PhoneNumber: "024XXXXXXX",
				Company:     "companyname",
			},
			"phone_book=&phone_number=024XXXXXXX&first_name=&last_name=&email=&company=companyname",
		},
	}

	for _, data := range tt {
		out := data.c.encode()
		if out != data.out {
			t.Logf("failed: expecting %s but got %s", data.out, out)
		}
	}
}
