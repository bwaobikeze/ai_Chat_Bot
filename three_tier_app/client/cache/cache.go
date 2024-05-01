package cache

import (
	"bytes"
	"fmt"
	"io"
)

// HTTPRequester is an interface for creating HTTP requests.
type HTTPRequester interface {
	MakeHTTPRequest(method string, url string, body io.Reader) (string, error)
}

type Cache struct {
	http HTTPRequester
}

// get cached query result
func (c *Cache) GetCachedQueryResult(query string) (string, error) {
	fmt.Println("calling cache with query")
	method := "GET"
	url := "http://h2memdb:8083"
	bodyString := fmt.Sprintf(`{
		"query": "%s"
	}`, query)
	bodyBytes := bytes.NewBuffer([]byte(bodyString))
	result, err := c.http.MakeHTTPRequest(method, url, bodyBytes)
	if err != nil {
		fmt.Printf("error getting cached query result: %v\n", err.Error())
		return "", err
	}
	return result, nil
}

// update cached query result
func (c *Cache) UpdateCachedQueryResult(query string, result string) error {
	fmt.Println("updating cached query result")
	method := "POST"
	url := "http://h2memdb:8083"
	bodyString := fmt.Sprintf(`{
		"query": "%s",
		"result": %s
	}`, query, result)
	body := []byte(bodyString)
	bodyBytes := bytes.NewBuffer([]byte(body))
	_, err := c.http.MakeHTTPRequest(method, url, bodyBytes)
	if err != nil {
		fmt.Printf("error updating cached query result: %v\n", err.Error())
		return err
	}
	return nil
}

// clear cached query result
func (c *Cache) ClearCachedQueryResult() error {
	fmt.Println("clearing cached queries")
	method := "POST"
	url := "http://h2memdb:8083"
	_, err := c.http.MakeHTTPRequest(method, url, nil)
	if err != nil {
		fmt.Printf("error clearing cached query results: %v\n", err.Error())
		return err
	}
	return nil
}

func NewCache(http HTTPRequester) Cache {
	return Cache{http: http}
}
