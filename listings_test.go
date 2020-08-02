package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListing(t *testing.T) {
	var res Listing
	b := unmarshalGoldenFileBytes(t, "listing.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/listings/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.Listing(167919)
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestListings(t *testing.T) {
	var res Listings
	b := unmarshalGoldenFileBytes(t, "listings.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/listings?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.Listings(Defaults())
	assert.NoError(t, err)
	assert.Equal(t, res, article)
}
func TestListingsByCategory(t *testing.T) {
	var res Listings
	b := unmarshalGoldenFileBytes(t, "listings.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/listings/category/cfp", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.ListingsByCategory(ListingCategoryCFP)
	assert.NoError(t, err)
	assert.Equal(t, res, article)
}
