package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// MockNeo4jDriver implements the neo4j.Driver interface for testing.
type MockNeo4jDriver struct {
	ExecuteQueryResult *neo4j.EagerResult
	ExecuteQueryError  error
	CloseError         error
}

func (m *MockNeo4jDriver) Session(accessMode neo4j.AccessMode, bookmarks ...string) (neo4j.Session, error) {
	return nil, errors.New("not implemented")
}

func (m *MockNeo4jDriver) Close() error {
	return m.CloseError
}

func (m *MockNeo4jDriver) VerifyExecuteQueryCalled(t *testing.T, expectedQuery string) {
	// Implement verification logic as needed
}

func TestExecuteDBQuery(t *testing.T) {

	// Create a DB instance with the mock driver
	db := NewDB("test","test","test")

	query := "MATCH (n) RETURN n"

	result, err := db.ExecuteDBQuery(query)

	assert.Error(t, err, "Expected error for a query execution error")
	assert.Nil(t, result, "Expected nil result for a query execution error")
}
