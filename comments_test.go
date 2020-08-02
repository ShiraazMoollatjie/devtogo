package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentWithReplies(t *testing.T) {
	var res Comment
	b := unmarshalGoldenFileBytes(t, "comment.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/comments/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.CommentWithReplies(167919)
	assert.NoError(t, err)
	assert.Equal(t, &res, article)
}

func TestAllComments(t *testing.T) {
	var res Comments
	b := unmarshalGoldenFileBytes(t, "comments.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/comments?a_id=167919", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.AllComments(167919)
	assert.NoError(t, err)
	assert.Equal(t, res, article)
}
