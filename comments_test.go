package devtogo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommentWithReplies(t *testing.T) {
	var res Comment
	b := unmarshalGoldenFileBytes(t, "comment.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/comments/167919", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.CommentWithReplies(167919)
	require.NoError(t, err)
	require.Equal(t, &res, article)
}

func TestAllComments(t *testing.T) {
	var res Comments
	b := unmarshalGoldenFileBytes(t, "comments.json", &res)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "/comments?a_id=167919", r.URL.String())
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}))
	client := NewClient(withBaseURL(ts.URL))
	article, err := client.AllComments(167919)
	require.NoError(t, err)
	require.Equal(t, res, article)
}
