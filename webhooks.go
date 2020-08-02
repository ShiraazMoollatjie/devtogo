package devtogo

import (
	"fmt"
	"time"
)

// Webhooks will return a list of webhooks they have previously registered.
func (c *Client) Webhooks() (Webhooks, error) {
	var res Webhooks
	err := c.get(c.baseURL+"/webhooks", &res)

	return res, err
}

// Webhook will return a single webhook given its id.
func (c *Client) Webhook(id int) (*Webhook, error) {
	var res Webhook
	err := c.get(c.baseURL+fmt.Sprintf("/webhooks/%d", id), &res)

	return &res, err
}

type Webhooks []Webhook

type Webhook struct {
	TypeOf    string    `json:"type_of"`
	ID        int       `json:"id"`
	Source    string    `json:"source"`
	TargetURL string    `json:"target_url"`
	Events    []string  `json:"events"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}
