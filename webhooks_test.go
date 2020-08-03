package devtogo

import (
	"encoding/json"
	"io/ioutil"
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

func TestDeleteWebhook(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/webhooks/12345", r.URL.Path)
		require.Equal(t, http.MethodDelete, r.Method)
		require.Equal(t, "myApiKey", r.Header.Get("api-key"))

		w.WriteHeader(http.StatusOK)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	err := client.DeleteWebhook(12345)
	require.NoError(t, err)
}

func TestCreateWebhook(t *testing.T) {
	var res Webhook
	b := unmarshalGoldenFileBytes(t, "webhook.json", &res)
	whr := CreateWebhookReq{
		Source:    "DEV",
		Events:    []string{"inserted", "updated"},
		TargetURL: "my targetURL",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/webhooks", r.URL.Path)
		require.Equal(t, http.MethodPost, r.Method)
		require.Equal(t, "myApiKey", r.Header.Get("api-key"))
		require.Equal(t, "application/json", r.Header.Get("Content-Type"))

		rb, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)

		var car webhookReq
		require.NoError(t, json.Unmarshal(rb, &car))
		require.Equal(t, webhookReq{Webhook: whr}, car)

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))

	client := NewClient(withBaseURL(ts.URL), WithApiKey("myApiKey"))
	articles, err := client.CreateWebhook(whr)
	require.NoError(t, err)
	require.Equal(t, res, articles)
}
