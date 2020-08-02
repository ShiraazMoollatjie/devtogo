package devtogo

import (
	"fmt"
	"time"
)

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

// CreateListing creates a listing if creating the user or the organization on which behalf the user is creating for has enough creadits.
func (c *Client) CreateListing(req CreateListingReq) (*Listing, error) {
	var res Listing
	err := c.post(c.baseURL+"/listings", listingReq{Listing: req}, &res)

	return &res, err
}

// UpdateListing updates a listing if creating the user or the organization on which behalf the user is creating for has enough creadits.
func (c *Client) UpdateListing(id int, req CreateListingReq) (*Listing, error) {
	var res Listing
	err := c.put(c.baseURL+fmt.Sprintf("/listings/%d", id), listingReq{Listing: req}, &res)

	return &res, err
}

type listingReq struct {
	Listing CreateListingReq `json:"listing"`
}

type CreateListingReq struct {
	Title             string          `json:"title"`
	BodyMarkdown      string          `json:"body_markdown"`
	Category          ListingCategory `json:"category"`
	Tags              []string        `json:"tags"`
	TagList           string          `json:"tag_list"`
	ExpiresAt         time.Time       `json:"expiresAt"`
	ContactViaConnect bool            `json:"contact_via_connect"`
	Location          string          `json:"location"`
	OrganizationID    int             `json:"organization_id"`
	Action            ListingAction   `json:"action"`
}

type ListingAction string

const (
	ListingActionDraft     ListingAction = "draft"
	ListingActionBump      ListingAction = "bump"
	ListingActionPublish   ListingAction = "publish"
	ListingActionUnpublish ListingAction = "unpublish"
)

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
	User          User            `json:"user"`
	Organization  Organization    `json:"organization"`
}
