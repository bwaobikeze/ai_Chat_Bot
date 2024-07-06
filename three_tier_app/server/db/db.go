package db

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DB interface {
    ExecuteDBQuery(query string) (*neo4j.EagerResult, error)
}

type db struct {
	url      string
	username string
	password string
}

func (d *db) ExecuteDBQuery(query string) (*neo4j.EagerResult, error) {
	fmt.Println("executing query")
	// get driver
	driver, err := neo4j.NewDriverWithContext(d.url, neo4j.BasicAuth(d.username, d.password, ""))
	if err != nil {
		fmt.Printf("error connecting to db: %v\n", err.Error())
		return nil, err
	}
	defer driver.Close(context.TODO())
	result, err := neo4j.ExecuteQuery(context.TODO(), driver,
		query,
		nil, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		fmt.Printf("error executing query: %v\n", err.Error())
		return nil, err
	}

	return result, nil
}

func NewDB(url string, username string, password string) db {
	return db{
		url:      url,
		username: username,
		password: password,
	}
}
