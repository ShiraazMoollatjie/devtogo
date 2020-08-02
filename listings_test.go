package devtogo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListing(t *testing.T) {
	var res Listing
	b := unmarshalGoldenFileBytes(t, "listing.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/listings/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.Listing(167919)
	require.NoError(t, err)
	require.Equal(t, &res, article)
}

func TestListings(t *testing.T) {
	var res Listings
	b := unmarshalGoldenFileBytes(t, "listings.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/listings?", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.Listings(Defaults())
	require.NoError(t, err)
	require.Equal(t, res, article)
}
func TestListingsByCategory(t *testing.T) {
	var res Listings
	b := unmarshalGoldenFileBytes(t, "listings.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/listings/category/cfp", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.ListingsByCategory(ListingCategoryCFP)
	require.NoError(t, err)
	require.Equal(t, res, article)
}

func TestCreateListing(t *testing.T) {
	var res Listing
	b := unmarshalGoldenFileBytes(t, "create_listing.json", &res)
	testListing := CreateListingReq{
		Tags:           []string{"events"},
		OrganizationID: 1,
		Action:         ListingActionBump,
		Category:       ListingCategoryCFP,
		BodyMarkdown:   "Awesome conference",
		Title:          "ACME Conference",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/listings", r.URL.Path)
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "myApiKey", r.Header.Get("api-key"))
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		var car listingReq
		require.NoError(t, json.Unmarshal(rb, &car))
		require.Equal(t, listingReq{Listing: testListing}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	listing, err := client.CreateListing(testListing)
	require.NoError(t, err)
	require.Equal(t, &res, listing)
}

func TestUpdateListing(t *testing.T) {
	var res Listing
	b := unmarshalGoldenFileBytes(t, "create_listing.json", &res)
	testListing := CreateListingReq{
		Tags:           []string{"events"},
		OrganizationID: 1,
		Action:         ListingActionPublish,
		Category:       ListingCategoryEducation,
		BodyMarkdown:   "Awesome conference",
		Title:          "ACME Conference",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/listings", r.URL.Path)
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "myApiKey", r.Header.Get("api-key"))
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		var car listingReq
		require.NoError(t, json.Unmarshal(rb, &car))
		require.Equal(t, listingReq{Listing: testListing}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	listing, err := client.CreateListing(testListing)
	require.NoError(t, err)
	require.Equal(t, &res, listing)
}
