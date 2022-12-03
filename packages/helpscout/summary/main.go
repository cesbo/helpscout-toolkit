package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/cesbo/helpscout-toolkit/pkg/helpscout"
)

var (
	apiKey = os.Getenv("HELPSCOUT_API_KEY")
)

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Headers    http.Header `json:"headers,omitempty"`
	Body       string      `json:"body,omitempty"`
}

func Main(args map[string]interface{}) (*Response, error) {
	hs := helpscout.NewHelpScout(apiKey)
	resp := &Response{}

	summary, err := hs.GetSummary()
	if err != nil {
		resp.StatusCode = 500
		resp.Body = err.Error()
		return resp, err
	}

	body, err := json.Marshal(summary)
	if err != nil {
		resp.StatusCode = 500
		resp.Body = err.Error()
		return resp, err
	}

	resp.StatusCode = 200
	resp.Headers = http.Header{
		"Content-Type": []string{"application/json"},
	}
	resp.Body = string(body)

	return resp, nil
}
