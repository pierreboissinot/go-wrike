package wrike

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() (*http.ServeMux, *httptest.Server, *Client) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// mux is the HTTP request multiplexer used with the test server.
	mux := http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	client := NewClient(nil, "")
	client.SetBaseURL(server.URL)

	return mux, server, client
}

// teardown closes the test HTTP server.
func teardown(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}
