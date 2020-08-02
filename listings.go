package devtogo

import "fmt"

// Listing will retrieve a listing based on the provided id.
func (c *Client) Listing(id int) (*Listing, error) {
	var res Listing
	err := c.get(fmt.Sprintf(c.baseURL+"/listings/%d", id), &res)

	return &res, err
}

// Listings will retrieve listings. "Listings" are classified ads that users create on DEV.
func (c *Client) Listings(args Arguments) (Listings, error) {
	var res Listings
	qp := args.toQueryParams().Encode()
	err := c.get(c.baseURL+"/listings?"+qp, &res)

	return res, err
}

// ListingsByCategory will retrieve listings belonging to the provided category
func (c *Client) ListingsByCategory(category ListingCategory) (Listings, error) {
	var res Listings
	err := c.get(fmt.Sprintf(c.baseURL+"/listings/category/%s", category), &res)

	return res, err
}

type ListingCategory string

const (
	ListingCategoryCFP       ListingCategory = "cfp"
	ListingCategoryForHire   ListingCategory = "forhire"
	ListingCategoryCollabs   ListingCategory = "collabs"
	ListingCategoryEducation ListingCategory = "education"
	ListingCategoryJobs      ListingCategory = "jobs"
	ListingCategoryMentors   ListingCategory = "mentors"
	ListingCategoryProducts  ListingCategory = "products"
	ListingCategorymentees   ListingCategory = "mentees"
	ListingCategoryForSale   ListingCategory = "forsale"
	ListingCategoryEvents    ListingCategory = "events"
	ListingCategoryMisc      ListingCategory = "misc"
)

type Listings []Listing

type Listing struct {
	TypeOf        string          `json:"type_of"`
	ID            int             `json:"id"`
	Title         string          `json:"title"`
	Slug          string          `json:"slug"`
	BodyMarkdown  string          `json:"body_markdown"`
	TagList       string          `json:"tag_list"`
	Tags          []string        `json:"tags"`
	Category      ListingCategory `json:"category"`
	ProcessedHTML string          `json:"processed_html"`
	Published     bool            `json:"published"`
	User          struct {
		Name            string      `json:"name"`
		Username        string      `json:"username"`
		TwitterUsername interface{} `json:"twitter_username"`
		GithubUsername  string      `json:"github_username"`
		WebsiteURL      interface{} `json:"website_url"`
		ProfileImage    string      `json:"profile_image"`
		ProfileImage90  string      `json:"profile_image_90"`
	} `json:"user"`
}
