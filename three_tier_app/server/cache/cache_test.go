package cache

import (
	"bytes"
	"errors"
	"io"
	"testing"
)

// MockHTTPRequester implements the HTTPRequester interface for testing.
type MockHTTPRequester struct {
	ExpectedMethod string
	ExpectedURL    string
	ExpectedBody   io.Reader
	Response       string
	Err            error
}

func (m *MockHTTPRequester) MakeHTTPRequest(method string, url string, body io.Reader) (string, error) {
	if m.ExpectedMethod != method || m.ExpectedURL != url {
		return "", errors.New("unexpected method or URL")
	}

	if m.ExpectedBody != nil {
		// Compare the expected and actual bodies
		expectedBodyBytes, _ := io.ReadAll(m.ExpectedBody)
		actualBodyBytes, _ := io.ReadAll(body)
		if !bytes.Equal(expectedBodyBytes, actualBodyBytes) {
			return "", errors.New("unexpected request body")
		}
	}

	return m.Response, m.Err
}

func TestGetCachedQueryResult(t *testing.T) {
	// Mock HTTPRequester for testing
	mockHTTP := &MockHTTPRequester{
		ExpectedMethod: "GET",
		ExpectedURL:    "http://h2memdb:8083",
		Response:       `{"result":"cached result"}`,
		Err:            nil,
	}

	cache := NewCache(mockHTTP)
	result, err := cache.GetCachedQueryResult("SELECT * FROM table")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedResult := "{\"result\":\"cached result\"}"
	if result != expectedResult {
		t.Errorf("Expected result %q, got %q", expectedResult, result)
	}
}

func TestUpdateCachedQueryResult(t *testing.T) {
	// Mock HTTPRequester for testing
	mockHTTP := &MockHTTPRequester{
		ExpectedMethod: "POST",
		ExpectedURL:    "http://h2memdb:8083",
		Response:       `{"status":"success"}`,
		Err:            nil,
	}

	cache := NewCache(mockHTTP)
	err := cache.UpdateCachedQueryResult("SELECT * FROM table", "new result")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestClearCachedQueryResult(t *testing.T) {
	// Mock HTTPRequester for testing
	mockHTTP := &MockHTTPRequester{
		ExpectedMethod: "POST",
		ExpectedURL:    "http://h2memdb:8083",
		Err:            nil,
	}

	cache := NewCache(mockHTTP)
	err := cache.ClearCachedQueryResult()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
