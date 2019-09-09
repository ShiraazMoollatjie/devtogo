package devtogo

import "net/http"

// Client makes all the API calls to dev.to.
type Client struct {
	baseURL string
	c       *http.Client
}

// Option allows the client to be configured with different options.
type Option func(*Client)

func withBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// NewClient creates a dev.to client with the provided options.
func NewClient(opts ...Option) *Client {
	res := &Client{
		baseURL: "http://dev.to/api",
	}
	for _, o := range opts {
		o(res)
	}

	return res
}
