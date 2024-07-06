package serverController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// response struct for server request
type serverResponse struct {
	Message string      `json:",omitempty"`
	Data    interface{} `json:",omitempty"`
}

// Request body
type requestBody struct {
	Query string `json:",omitempty"`
}

type dbInterface interface {
	ExecuteDBQuery(query string) (*neo4j.EagerResult, error)
}

type cacheInterface interface {
	GetCachedQueryResult(query string) (string, error)
	UpdateCachedQueryResult(query string, result string) error
	ClearCachedQueryResult() error
}

type controller struct {
	db dbInterface
	c  cacheInterface
}

// handle valid requests
func (c *controller) HandleQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==handling query==")
	w.Header().Add("content-type", "application/json")
	serverMessage := serverResponse{Message: "unable to decode request body for query", Data: nil}
	var request requestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Printf("unable to decode body: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&serverMessage)
		return
	}
	queryResult, err := c.db.ExecuteDBQuery(request.Query)
	if err != nil {
		fmt.Printf("error executing query: %v\n", err.Error())
		serverMessage.Message = "error executing query"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&serverMessage)
		return
	}
	if queryResult.Summary.StatementType() != neo4j.StatementTypeReadOnly {
		fmt.Println("Clearing cache for non-readonly query")
		err = c.c.ClearCachedQueryResult()
		if err != nil {
			fmt.Printf("error clearing cache: %v\n", err.Error())
		}
	}
	var output []map[string]any
	for _, record := range queryResult.Records {
		output = append(output, record.AsMap())
	}
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		fmt.Printf("error converting query output to json: %v\n", err.Error())
		serverMessage.Message = "error preparing query output"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&serverMessage)
		return
	}
	serverMessage.Message = "query execution successful"
	serverMessage.Data = string(jsonBytes)
	json.NewEncoder(w).Encode(&serverMessage)
	fmt.Println("==query handled==")
}

// handle not found
func (c *controller) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==invalid path==")
	serverMessage := serverResponse{Message: "invalid path", Data: nil}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&serverMessage)
	fmt.Println("==response sent==")
}

func NewController(db dbInterface, cache cacheInterface) controller {
	return controller{db: db, c: cache}
}
