package devto_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	NewClient(withBaseURL(ts.URL))
	fmt.Println(ts.URL)
	assert.True(t, strings.Contains(ts.URL, "127.0.0.1"))
}
