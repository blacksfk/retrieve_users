package github

import (
	"testing"
)

func TestRetrieveUser(t *testing.T) {
	// expected user login
	expected := "blacksfk"

	// try to retrieve a known good user
	actual, e := RetrieveUser(expected)

	if e != nil {
		t.Fatal(e)
	}

	if actual.Login != expected {
		t.Fatalf("Expected: %s. Actual: %s", expected, actual.Login)
	}
}
