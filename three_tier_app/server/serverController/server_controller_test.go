package serverController

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/stretchr/testify/assert"
)

// MockDB implements the dbInterface for testing.
type MockDB struct {
	ExecuteDBQueryResult *neo4j.EagerResult
	ExecuteDBQueryError  error
}

func (m *MockDB) ExecuteDBQuery(query string) (*neo4j.EagerResult, error) {
	return m.ExecuteDBQueryResult, m.ExecuteDBQueryError
}

// MockCache implements the cacheInterface for testing.
type MockCache struct {
	GetCachedQueryResultResult string
	GetCachedQueryResultError  error
	UpdateCachedQueryResultError error
	ClearCachedQueryResultError  error
}

func (m *MockCache) GetCachedQueryResult(query string) (string, error) {
	return m.GetCachedQueryResultResult, m.GetCachedQueryResultError
}

func (m *MockCache) UpdateCachedQueryResult(query string, result string) error {
	return m.UpdateCachedQueryResultError
}

func (m *MockCache) ClearCachedQueryResult() error {
	return m.ClearCachedQueryResultError
}

func TestHandleQuery(t *testing.T) {
	// Mocks
	mockDB := &MockDB{
		ExecuteDBQueryResult: &neo4j.EagerResult{},
		ExecuteDBQueryError:  nil,
	}

	mockCache := &MockCache{
		GetCachedQueryResultResult: "",
		GetCachedQueryResultError:  nil,
		UpdateCachedQueryResultError: nil,
		ClearCachedQueryResultError:  nil,
	}

	// Controller instance with mocks
	controller := NewController(mockDB, mockCache)

	rec := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, rec.Code, "Expected HTTP 200 OK status")

	// Test case 2: Error in decoding request body
	req := httptest.NewRequest("POST", "/query", bytes.NewBuffer([]byte("invalid json")))
	rec = httptest.NewRecorder()

	controller.HandleQuery(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code, "Expected HTTP 400 Bad Request status")
}

func TestHandleNotFound(t *testing.T) {
	mockDB := &MockDB{}
	mockCache := &MockCache{}

	controller := NewController(mockDB, mockCache)

	req := httptest.NewRequest("GET", "/invalid-path", nil)
	rec := httptest.NewRecorder()

	controller.HandleNotFound(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "Expected HTTP 404 Not Found status")
}
