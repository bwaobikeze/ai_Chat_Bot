package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockQuery implements the queryInterface for testing.
type MockQuery struct{}

func (m *MockQuery) IsValidQuery(query string) bool {
	return query == "VALID_QUERY"
}

func (m *MockQuery) ExecuteQuery(query string) (string, int) {
	return "MockedResult", http.StatusOK
}

// MockCache implements the cacheInterface for testing.
type MockCache struct{}

func (m *MockCache) GetCachedQueryResult(query string) (string, error) {
	return "", nil
}

func (m *MockCache) UpdateCachedQueryResult(query string, result string) error {
	return nil
}

func (m *MockCache) ClearCachedQueryResult() error {
	return nil
}

func TestHandleRequestValidQuery(t *testing.T) {
	controller := NewController(&MockQuery{}, &MockCache{})

	// Create a test request with a valid query
	reqBody := `{"Query": "VALID_QUERY"}`
	req := httptest.NewRequest("POST", "/handle-request", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	controller.HandleRequest(w, req)

	// Check the response
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var serverMessage serverResponse
	err := json.NewDecoder(resp.Body).Decode(&serverMessage)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if serverMessage.Message != "query execution successful" {
		t.Errorf("Expected success message, got %q", serverMessage.Message)
	}

	if serverMessage.Data != "MockedResult" {
		t.Errorf("Expected result %q, got %q", "MockedResult", serverMessage.Data)
	}
}

func TestHandleRequestInvalidQuery(t *testing.T) {
	controller := NewController(&MockQuery{}, &MockCache{})

	// Create a test request with an invalid query
	reqBody := `{"Query": "INVALID_QUERY"}`
	req := httptest.NewRequest("POST", "/handle-request", strings.NewReader(reqBody))
	w := httptest.NewRecorder()

	controller.HandleRequest(w, req)

	// Check the response
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var serverMessage serverResponse
	err := json.NewDecoder(resp.Body).Decode(&serverMessage)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if serverMessage.Message != "invalid cypher query" {
		t.Errorf("Expected invalid query message, got %q", serverMessage.Message)
	}
}

func TestHandleNotFound(t *testing.T) {
	controller := NewController(&MockQuery{}, &MockCache{})

	// Create a test request for not found scenario
	req := httptest.NewRequest("GET", "/invalid-path", nil)
	w := httptest.NewRecorder()

	controller.HandleNotFound(w, req)

	// Check the response
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
	}

	var serverMessage serverResponse
	err := json.NewDecoder(resp.Body).Decode(&serverMessage)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if serverMessage.Message != "invalid path" {
		t.Errorf("Expected invalid path message, got %q", serverMessage.Message)
	}
}
