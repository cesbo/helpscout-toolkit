package main

import (
	"encoding/json"

	"github.com/cesbo/helpscout-toolkit/internal/helpscout"
)

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(args map[string]interface{}) (*Response, error) {
	resp := &Response{}

	collections, err := helpscout.ListCollections()
	if err != nil {
		resp.StatusCode = 500
		resp.Body = err.Error()
		return resp, err
	}

	data, err := json.Marshal(collections)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = err.Error()
		return resp, err
	}

	resp.StatusCode = 200
	resp.Body = string(data)

	return resp, nil
}
