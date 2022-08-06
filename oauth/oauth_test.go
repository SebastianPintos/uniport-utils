package oauth

import (
	"io"
	"net/http"
	url2 "net/url"
	"strings"
	"testing"
)

func TestAuthenticateRequest(t *testing.T) {
	r := io.NopCloser(strings.NewReader("{'test'{'msg':'hi'}}")) // r type is io.ReadCloser
	url, _ := url2.Parse("http://localhost:3333/oauth/access_token?access_token=abc123")

	request := http.Request{Method: "GET", Body: r, URL: url}
	AuthenticateRequest(&request)
}
