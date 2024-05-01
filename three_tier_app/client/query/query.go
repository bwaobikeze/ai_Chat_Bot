package query

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ubsss/pp-uu444/client/parser"
	"github.com/antlr4-go/antlr/v4"
)

// response struct for server request
type serverResponse struct {
	Message string `json:",omitempty"`
	Data    string `json:",omitempty"`
}

type query struct{}

// cypher listener struct
type CypherListener struct {
	*parser.BaseCypherListener
}

type CustomSyntaxError struct {
	line, column int
	msg          string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *CustomSyntaxError) Error() string {
	return c.msg
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	newError := &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	}
	c.Errors = append(c.Errors, newError)
}

// validate cypher query
func (q *query) IsValidQuery(query string) bool {
	fmt.Println("validating cypher query")
	lexerErrors := &CustomErrorListener{}
	parserErrors := &CustomErrorListener{}
	// set up input stream
	inputStream := antlr.NewInputStream(query)

	// create lexer based on input
	lexer := parser.NewCypherLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// create parser
	parser := parser.NewCypherParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)
	// parser.BuildParseTrees = true

	// Parse request
	antlr.ParseTreeWalkerDefault.Walk(&CypherListener{}, parser.OC_Cypher())

	return len(lexerErrors.Errors) < 1 && len(parserErrors.Errors) < 1
}

// execute cypher query
func (q *query) ExecuteQuery(query string) (string, int) {
	fmt.Println("calling server with query")
	values := map[string]string{"query": query}
	json_data, err := json.Marshal(values)
	if err != nil {
		fmt.Printf("unable to create request body: %v\n", err.Error())
		return "", http.StatusInternalServerError
	}
	resp, err := http.Post("http://server:8082", "application/json",
		bytes.NewBuffer(json_data))
	if err != nil {
		fmt.Printf("error making request: %v\n", err.Error())
		return "", http.StatusInternalServerError
	}
	defer resp.Body.Close()
	var res serverResponse
	json.NewDecoder(resp.Body).Decode(&res)
	return res.Data, http.StatusOK
}

func NewQuery() query {
	return query{}
}
