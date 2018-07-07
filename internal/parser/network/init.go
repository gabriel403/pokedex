package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"pokedex/internal/log"
)

func CreateGETRequest(url string) (*http.Request, error) {
	return createRequest(url, "", "GET")
}

func createRequest(url string, body string, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func UnmarshallResponse(r *http.Response, request interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func UnmarshallBody(body []byte, request interface{}) error {
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
