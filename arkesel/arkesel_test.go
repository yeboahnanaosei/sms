package arkesel

import (
	"testing"
)

func TestNew(t *testing.T) {
	tt := []struct{in, out string}{
		{"key1", "key1"},
		{"key2", "key2"},
	}

	for _, data := range tt {
		a := New(data.in)
		if a.params.Get("api_key") != data.out {
			t.Logf("test failed: expected %s but got %s", data.in, data.out)
		}
	}
}

func TestMessage(t *testing.T) {
	arksel := New("apiKey")

	tt := []struct{in, out string}{
		{"message1", "message1"},
		{"anothermessage", "anothermessage"},
	}

	for _, data := range tt {
		arksel.Message(data.in)
		if arksel.params.Get("sms") != data.out {
			t.Logf("failed: expected %s but got %s", data.in, data.out)
		}
	}
}

func TestFrom(t *testing.T) {
	arksel := New("apiKey")

	tt := []struct{in, out string}{
		{"sender1", "sender1"},
		{"anotherID", "anotherID"},
	}

	for _, data := range tt {
		arksel.From(data.in)
		if arksel.params.Get("from") != data.out {
			t.Logf("failed: expected %s but got %s", data.in, data.out)
		}
	}
}


func TestTo(t *testing.T) {
	arksel := New("apiKey")

	tt := []struct{in, out string}{
		{"23323XXXXXXX", "23323XXXXXXX"},
		{"026XXXXXXX", "026XXXXXXX"},
	}

	for _, data := range tt {
		arksel.To(data.in)
		if arksel.params.Get("to") != data.out {
			t.Logf("failed: expected %s but got %s", data.in, data.out)
		}
	}
}
func TestSetAPIKey(t *testing.T) {
	arksel := New("apiKey")

	tt := []struct{in, out string}{
		{"key1", "key1"},
		{"keyabcd==", "keyabcd=="},
	}

	for _, data := range tt {
		arksel.SetAPIKey(data.in)
		if arksel.params.Get("api_key") != data.out {
			t.Logf("failed: expected %s but got %s", data.in, data.out)
		}
	}
}