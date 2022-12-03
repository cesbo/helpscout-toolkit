package helpscout

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	userAgent = "helpscout-toolkit/0.1"
)

var (
	apiURL, _ = url.Parse("https://docsapi.helpscout.net/v1/")
	apiKey    = os.Getenv("HELPSCOUT_API_KEY")
	client    = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:    3 * time.Second,
			MaxIdleConns:           0,
			MaxIdleConnsPerHost:    4,
			MaxConnsPerHost:        0,
			IdleConnTimeout:        90 * time.Second,
			ResponseHeaderTimeout:  2 * time.Second,
			ExpectContinueTimeout:  1 * time.Second,
			MaxResponseHeaderBytes: 2 * 1024,
			ForceAttemptHTTP2:      true,
		},
	}
)

func getJSON(path string, query url.Values, target interface{}) error {
	if apiKey == "" {
		return fmt.Errorf("HELPSCOUT_API_KEY is not set")
	}

	request, _ := http.NewRequest(http.MethodGet, "", nil)
	request.URL = apiURL.ResolveReference(&url.URL{
		Path:     path,
		RawQuery: query.Encode(),
	})
	request.Header.Add("User-Agent", userAgent)
	request.SetBasicAuth(apiKey, "X")

	response, err := client.Do(request)
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return fmt.Errorf("request send: %w", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != 200 {
		return fmt.Errorf("request response: %s", response.Status)
	}

	// Decode response body
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("response decode: %w", err)
	}

	return nil
}
