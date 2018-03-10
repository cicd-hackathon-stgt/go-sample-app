package greeting

import "testing"

func TestGetGreeting(t *testing.T) {
	greeting := GetGreeting()
	if len(greeting) == 0 {
		t.Error("'greeting' has length 0")
	}
}
