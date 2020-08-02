package devtogo

// Followers will retrieve a list of the followers that you have.
func (c *Client) Followers(args Arguments) (Followers, error) {
	var res Followers
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/followers?"+qp, &res)

	return res, err
}

type Followers []struct {
	TypeOf       string `json:"type_of"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}
