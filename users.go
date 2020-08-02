package devtogo

import "fmt"

// LookupUser will retrieve single user, either by id or by the user's username.
func (c *Client) LookupUser(id string) (*UserProfile, error) {
	var res UserProfile
	err := c.get(c.baseURL+fmt.Sprintf("/users/%s", id), &res)

	return &res, err
}

// Me will retrieve the user profile for the configured authentication key.
func (c *Client) Me() (*UserProfile, error) {
	var res UserProfile
	err := c.get(c.baseURL+"/users/me", &res)

	return &res, err
}

// User represents a user from the dev.to api.
type User struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	ProfileImage    string `json:"profile_image"`
	ProfileImage90  string `json:"profile_image_90"`
}

type UserProfile struct {
	TypeOf          string `json:"type_of"`
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Summary         string `json:"summary"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteURL      string `json:"website_url"`
	Location        string `json:"location"`
	JoinedAt        string `json:"joined_at"`
	ProfileImage    string `json:"profile_image"`
}
