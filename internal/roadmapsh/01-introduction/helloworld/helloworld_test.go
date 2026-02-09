package helloworld

import (
	"fmt"
	"os/user"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := helloWorld()

	currentUser, err := user.Current()
	expected := fmt.Sprintf("Hello %s!", currentUser.Username)

	if expected != got && err != nil {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}
