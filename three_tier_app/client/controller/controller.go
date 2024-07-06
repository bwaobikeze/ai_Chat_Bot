package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// response struct for server request
type serverResponse struct {
	Message string `json:",omitempty"`
	Data    string `json:",omitempty"`
}

// Request body
type requestBody struct {
	Query string `json:",omitempty"`
}

type queryInterface interface {
	IsValidQuery(query string) bool
	ExecuteQuery(query string) (string, int)
}

type cacheInterface interface {
	GetCachedQueryResult(query string) (string, error)
	UpdateCachedQueryResult(query string, result string) error
	ClearCachedQueryResult() error
}

type controller struct {
	q queryInterface
	c cacheInterface
}

// handle valid requests
func (c *controller) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==got request==")
	w.Header().Add("content-type", "application/json")
	serverMessage := serverResponse{Message: "unable to decode request body for query", Data: "{}"}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Printf("unable to decode body: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&serverMessage)
		return
	}
	validCypherQuery := c.q.IsValidQuery(request.Query)
	if !validCypherQuery {
		fmt.Println("invalid cypher query")
		serverMessage.Message = "invalid cypher query"
		json.NewEncoder(w).Encode(&serverMessage)
		return
	}
	cachedResult, err := c.c.GetCachedQueryResult(request.Query)
	if err != nil {
		fmt.Printf("error getting cached query result: %v\n", err.Error())
	} else {
		if len(cachedResult) < 1 {
			fmt.Println("cached result empty")
		} else {
			fmt.Println("cache result value found")
			serverMessage.Message = "successful query result retrieval from cache"
			serverMessage.Data = cachedResult
			json.NewEncoder(w).Encode(&serverMessage)
			return
		}
	}
	data, status := c.q.ExecuteQuery(request.Query)
	if status != http.StatusOK {
		fmt.Println("unsuccessful query execution")
		w.WriteHeader(http.StatusInternalServerError)
		serverMessage.Message = "query execution unsuccessful"
		serverMessage.Data = "{}"
	} else {
		fmt.Println("successful query execution")
		w.WriteHeader(http.StatusOK)
		serverMessage.Message = "query execution successful"
		serverMessage.Data = data
	}
	if len(serverMessage.Data) > 0 {
		fmt.Println("updating cache with query result")
		jsonBytes, err := json.Marshal(serverMessage.Data)
		if err != nil {
			fmt.Printf("error formatting response json for cache: %v", err.Error())
		} else {
			err = c.c.UpdateCachedQueryResult(request.Query, string(jsonBytes))
			if err != nil {
				fmt.Printf("error updating cache: %v\n", err.Error())
			}
		}

	}
	json.NewEncoder(w).Encode(&serverMessage)
	fmt.Println("==request processed==")
}

// handle not found
func (c *controller) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==invalid path==")
	serverMessage := serverResponse{Message: "invalid path", Data: "{}"}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&serverMessage)
	fmt.Println("==response sent==")
}

func NewController(query queryInterface, cache cacheInterface) controller {
	return controller{q: query, c: cache}
}
