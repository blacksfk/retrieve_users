package github

import (
	"encoding/json"
	"github.com/blacksfk/retrieve_users"
	"io"
	"net/http"
)

const (
	URL = "https://api.github.com/users/"
)

// Retrieve a user from the API via their username.
func RetrieveUser(username string) (*retrieve_users.User, error) {
	// send the request
	res, e := http.Get(URL + username)

	if e != nil {
		return e
	}

	// read and close the body
	defer res.Body.Close()
	bytes, e := io.ReadAll(res.Body)

	if e != nil {
		return e
	}

	// decode the body into a native user type
	user := &retrieve_users.User{}
	e = json.Unmarshal(bytes, user)

	return user, e
}
