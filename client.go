// Package extremeiplookup contains an eXTReMe IP Lookup API client.
package extremeiplookup

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// API response statuses.
const (
	StatusFail    = "fail"
	StatusSuccess = "success"
)

const defaultBaseURL = "https://extreme-ip-lookup.com"

// Client an eXTReMe IP Lookup API client.
type Client struct {
	apiKey  string
	baseURL *url.URL

	HTTPClient *http.Client
}

// NewClient creates a new Client.
func NewClient(apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 5 * time.Second},
	}
}

// Lookup gets all geolocation information about an IP address.
func (c Client) Lookup(ctx context.Context, ip string) (*IPInfo, error) {
	endpoint := c.baseURL.JoinPath("json", ip)

	if c.apiKey != "" {
		query := endpoint.Query()
		query.Set("key", c.apiKey)
		endpoint.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(data))
	}

	defer func() { _ = resp.Body.Close() }()

	var apiResp IPInfo
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	if apiResp.Status != StatusSuccess {
		return nil, fmt.Errorf("%s: %s", apiResp.Status, apiResp.Message)
	}

	return &apiResp, nil
}
