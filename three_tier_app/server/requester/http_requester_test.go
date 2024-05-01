package requester

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeHTTPRequestGET(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test response"))
	}))
	defer server.Close()

	requester := NewHTTPRequester()
	result, err := requester.MakeHTTPRequest("GET", server.URL, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := "Test response"
	if result != expectedResult {
		t.Errorf("Expected result %q, got %q", expectedResult, result)
	}
}

func TestMakeHTTPRequestPOST(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test response"))
	}))
	defer server.Close()

	// Create a test request body
	body := bytes.NewBufferString(`{"key": "value"}`)

	requester := NewHTTPRequester()
	result, err := requester.MakeHTTPRequest("POST", server.URL, body)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := ""
	if result != expectedResult {
		t.Errorf("Expected result %q, got %q", expectedResult, result)
	}
}

func TestMakeHTTPRequestError(t *testing.T) {
	// Create a test HTTP server that always returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	requester := NewHTTPRequester()
	result, err := requester.MakeHTTPRequest("", server.URL, nil)

	if err == nil {
		t.Error("Expected an error, but got none")
	}

	if result != "" {
		t.Errorf("Expected an empty result, but got %q", result)
	}
}
