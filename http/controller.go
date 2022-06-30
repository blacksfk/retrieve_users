package http

import (
	"net/http"

	"github.com/blacksfk/retrieve_users"
	"github.com/blacksfk/retrieve_users/github"
	"github.com/blacksfk/uf"
)

func RetrieveUsers(w http.ResponseWriter, r *http.Request) error {
	var usernames []string

	// decode the body into the array
	e := uf.DecodeBodyJSON(r, &usernames)

	if e != nil {
		// something went wrong
		return e
	}

	if usernames == nil {
		return uf.BadRequest("usernames is null")
	}

	users := retrieve_users.Users{}

	// loop over the users and retrieve each one
	for _, username := range usernames {
		// check if the user already exists
		if users.Exists(username) {
			// skip
			continue
		}

		// retrieve the username from the API
		user, e := github.RetrieveUser(username)

		if e != nil {
			// error occurred retrieving user.
			// assume github API is rock solid and
			// that the user doesn't exist
			continue
		}

		// append the user
		users = users.SortedInsert(user)
	}

	return uf.SendJSON(w, users)
}
