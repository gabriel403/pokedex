package integrationtests

import (
	"net/http/httptest"
)

var (
	server    *httptest.Server
	serverURL string
)

func init() {
	server = httptest.NewServer(api.Handlers())
	serverURL = server.URL
}
