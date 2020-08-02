package devtogo

import (
	"errors"
	"fmt"
)

// AllComments will return the all top level comments with their nested comments as threads.
func (c *Client) AllComments(articleID int) (Comments, error) {
	if articleID == 0 {
		return nil, errors.New("articleID cannot be empty")
	}

	var res Comments
	err := c.get(c.baseURL+fmt.Sprintf("/comments?a_id=%d", articleID), &res)

	return res, err
}

// CommentWithReplies will retrieve a comment as well as his descendants comments for the provided id.
func (c *Client) CommentWithReplies(id int) (*Comment, error) {
	var res Comment
	err := c.get(c.baseURL+fmt.Sprintf("/comments/%d", id), &res)

	return &res, err
}

type Comments []Comment

type Comment struct {
	TypeOf   string `json:"type_of"`
	IDCode   string `json:"id_code"`
	BodyHTML string `json:"body_html"`
	User     struct {
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		TwitterUsername interface{} `json:"twitter_username"`
		GithubUsername  string      `json:"github_username"`
		WebsiteURL      interface{} `json:"website_url"`
		ProfileImage    string      `json:"profile_image"`
		ProfileImage90  string      `json:"profile_image_90"`
	} `json:"user"`
	Children []Comment `json:"children"`
}
