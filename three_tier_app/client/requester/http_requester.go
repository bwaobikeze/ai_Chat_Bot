package requester

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HTTPRequester struct{}

func (h *HTTPRequester) MakeHTTPRequest(method string, url string, body io.Reader) (string, error) {
	fmt.Println("making http call")
	result := ""
	if len(method) < 1 {
		fmt.Println("no method provided")
		return result, errors.New("no method provided")
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("error making http request: %v\n", err.Error())
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error executing request: %v\n", err.Error())
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		fmt.Println("status was ok")
	} else {
		fmt.Println("status was not ok")
	}
	if method == "GET" {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("error parsing request response body: %v\n", err.Error())
			return result, err
		}
		result = string(bodyBytes)
	}
	return result, nil
}

func NewHTTPRequester() HTTPRequester {
	return HTTPRequester{}
}