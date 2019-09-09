package devtogo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestGetArticles(t *testing.T) {
	p := filepath.Join("testdata", "articles.json")
	b, err := ioutil.ReadFile(p)
	var res Articles
	err = json.Unmarshal(b, &res)
	assert.NoError(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/articles", r.URL.Path)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
	}))
	client := NewClient(withBaseURL(ts.URL))
	articles, err := client.GetArticles()
	assert.NoError(t, err)
	assert.Equal(t, res, articles)
}
