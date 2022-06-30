package retrieve_users

import (
	"time"
)

// Represents a github user.
type User struct {
	Login               string    `json:"login"`
	Id                  int64     `json:"id"`
	Node_id             string    `json:"node_id"`
	Avatar_url          string    `json:"avatar_url"`
	Gravatar_id         string    `json:"gravatar_id"`
	Url                 string    `json:"url"`
	Html_url            string    `json:"html_url"`
	Followers_url       string    `json:"followers_url"`
	Following_url       string    `json:"following_url"`
	Gists_url           string    `json:"gists_url"`
	Starred_url         string    `json:"starred_url"`
	Subscriptions_url   string    `json:"subscriptions_url"`
	Organizations_url   string    `json:"organizations_url"`
	Repos_url           string    `json:"repos_url"`
	Events_url          string    `json:"events_url"`
	Received_events_url string    `json:"received_events_url"`
	Type                string    `json:"type"`
	Site_admin          bool      `json:"site_admin"`
	Name                string    `json:"name"`
	Company             string    `json:"company"`
	Blog                string    `json:"blog"`
	Location            string    `json:"location"`
	Email               string    `json:"email"`
	Hireable            string    `json:"hireable"`
	Bio                 string    `json:"bio"`
	Twitter_username    string    `json:"twitter_username"`
	Public_repos        int64     `json:"public_repos"`
	Public_gists        int64     `json:"public_gists"`
	Followers           int64     `json:"followers"`
	Following           int64     `json:"following"`
	Created_at          time.Time `json:"created_at"`
	Updated_at          time.Time `json:"updated_at"`
}

// Returns true if a < b and false otherwise
func (a *User) Cmp(b *User) bool {
	return a.Name < b.Name
}
