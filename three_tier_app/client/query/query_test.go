package query

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidQueryValid(t *testing.T) {
	q := NewQuery()
	validQuery := "MATCH (n) RETURN n"
	isValid := q.IsValidQuery(validQuery)

	if !isValid {
		t.Errorf("Expected valid query, got invalid")
	}
}

func TestIsValidQueryInvalid(t *testing.T) {
	q := NewQuery()
	invalidQuery := "INVALID QUERY"
	isValid := q.IsValidQuery(invalidQuery)

	if isValid {
		t.Errorf("Expected invalid query, got valid")
	}
}

func TestExecuteQueryServerError(t *testing.T) {
	// Create a test server that returns an error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message": "Internal Server Error"}`))
	}))
	defer server.Close()

	q := NewQuery()
	queryResult, statusCode := q.ExecuteQuery("MATCH (n) RETURN n")

	if statusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, statusCode)
	}

	if queryResult != "" {
		t.Errorf("Expected empty result, got %q", queryResult)
	}
}
