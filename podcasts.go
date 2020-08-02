package devtogo

// PodcaseEpisodes returns podcast episodes that are according to https://docs.dev.to/api/#operation/getPodcastEpisodes.
func (c *Client) PodcastEpisodes(args Arguments) (PodcastEpisodes, error) {
	var res PodcastEpisodes
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/podcast_episodes?"+qp, &res)

	return res, err
}

type PodcastEpisodes []struct {
	TypeOf    string `json:"type_of"`
	ClassName string `json:"class_name"`
	ID        int    `json:"id"`
	Path      string `json:"path"`
	Title     string `json:"title"`
	ImageURL  string `json:"image_url"`
	Podcast   struct {
		Title    string `json:"title"`
		Slug     string `json:"slug"`
		ImageURL string `json:"image_url"`
	} `json:"podcast"`
}
