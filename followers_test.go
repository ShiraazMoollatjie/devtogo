package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFollowers(t *testing.T) {
	var res Followers
	b := unmarshalGoldenFileBytes(t, "followers.json", &res)

	tests := []struct {
		name                string
		arguments           Arguments
		expectedQueryParams string
	}{
		{"No params", Defaults(), ""},
		{"Page param", Arguments{"page": "1"}, "page=1"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/followers?"+test.expectedQueryParams, r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			client := NewClient(withBaseURL(ts.URL))
			followers, err := client.Followers(test.arguments)
			require.NoError(t, err)
			require.Equal(t, res, followers)
		})
	}
}
