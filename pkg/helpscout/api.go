package helpscout

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	client = &http.Client{
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

func (hs *HelpScout) getJSON(path string, query url.Values, target interface{}) error {
	if hs.apiKey == "" {
		return fmt.Errorf("HELPSCOUT_API_KEY is not set")
	}

	request, _ := http.NewRequest(http.MethodGet, "", nil)
	request.URL = hs.apiURL.ResolveReference(&url.URL{
		Path:     path,
		RawQuery: query.Encode(),
	})
	request.Header.Add("User-Agent", userAgent)
	request.SetBasicAuth(hs.apiKey, "X")

	response, err := client.Do(request)
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return fmt.Errorf("request send: %w", err)
	}
	defer response.Body.Close()

	if v := response.Header.Get("X-RateLimit-Remaining"); v != "" {
		hs.rateLimitRemain, _ = strconv.Atoi(v)
	}

	if v := response.Header.Get("X-RateLimit-Reset"); v != "" {
		hs.rateLimitReset, _ = strconv.Atoi(v)
	}

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

// RateLimit returns the number of requests remaining in the current rate limit window
// and duration until the current rate limit window expires.
func (hs *HelpScout) RateLimit() (int, time.Duration) {
	return hs.rateLimitRemain, time.Duration(hs.rateLimitReset) * time.Second
}
