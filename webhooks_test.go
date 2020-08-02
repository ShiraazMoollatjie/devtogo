package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebhooks(t *testing.T) {
	var res Webhooks
	b := unmarshalGoldenFileBytes(t, "webhooks.json", &res)

	tests := []struct {
		name                string
		arguments           Arguments
		expectedQueryParams string
	}{
		{"No params", Defaults(), ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, "/webhooks", r.URL.String())
				w.WriteHeader(http.StatusOK)
				w.Write(b)
			}))

			client := NewClient(withBaseURL(ts.URL))
			articles, err := client.Webhooks()
			require.NoError(t, err)
			require.Equal(t, res, articles)
		})
	}
}

func TestWebhook(t *testing.T) {
	var res Webhook
	b := unmarshalGoldenFileBytes(t, "webhook.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/webhooks/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.Webhook(167919)
	require.NoError(t, err)
	require.Equal(t, &res, article)
}
