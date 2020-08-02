package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookupUser(t *testing.T) {
	var res UserProfile
	b := unmarshalGoldenFileBytes(t, "user.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/users/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	up, err := client.LookupUser("167919")
	require.NoError(t, err)
	require.Equal(t, &res, up)
}

func TestLookupMe(t *testing.T) {
	var res UserProfile
	b := unmarshalGoldenFileBytes(t, "user.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/users/me", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	up, err := client.Me()
	require.NoError(t, err)
	require.Equal(t, &res, up)
}
