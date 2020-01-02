package main

import (
	"io/ioutil"
	"net/http"
)

func request(link string) (*string, error) {
	response, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	content := string(body)

	return &content, nil
}
