package retrieve_users

// Represents a github user.
type User struct {
	Login               string `json:"login"`
	Id                  string `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"type"`
	Site_admin          string `json:"site_admin"`
	Name                string `json:"name"`
	Company             string `json:"company"`
	Blog                string `json:"blog"`
	Location            string `json:"location"`
	Email               string `json:"email"`
	Hireable            string `json:"hireable"`
	Bio                 string `json:"bio"`
	Twitter_username    string `json:"twitter_username"`
	Public_repos        string `json:"public_repos"`
	Public_gists        string `json:"public_gists"`
	Followers           string `json:"followers"`
	Following           string `json:"following"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
}

// Returns true if a < b and false otherwise
func (a *User) Cmp(b *User) bool {
	return a.Login < b.Login
}
