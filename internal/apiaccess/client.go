package apiaccess

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient defines the interface for HTTP operations
type HTTPClient interface {
	Get(ctx context.Context, url string) (*http.Response, error)
	Post(ctx context.Context, url string, data interface{}) (*http.Response, error)
}

// DefaultHTTPClient implements HTTPClient with timeout and error handling
type DefaultHTTPClient struct {
	client  *http.Client
	timeout time.Duration
}

// NewHTTPClient creates a new HTTP client with specified timeout
func NewHTTPClient(timeout time.Duration) HTTPClient {
	return &DefaultHTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		timeout: timeout,
	}
}

// Get performs a GET request with timeout and error handling
func (c *DefaultHTTPClient) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request for %s: %w", url, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "home-control-center/1.0")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET request to %s failed: %w", url, err)
	}

	return resp, nil
}

// Post performs a POST request with JSON data, timeout and error handling
func (c *DefaultHTTPClient) Post(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	var body io.Reader
	
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON data for POST to %s: %w", url, err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request for %s: %w", url, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "home-control-center/1.0")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST request to %s failed: %w", url, err)
	}

	return resp, nil
}

// readResponseBody reads and closes the response body, returning the content and any error
func readResponseBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	
	return body, nil
}

// checkHTTPStatus checks if the HTTP status code indicates success
func checkHTTPStatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := readResponseBody(resp)
		return fmt.Errorf("HTTP request failed with status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}