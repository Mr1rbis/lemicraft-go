package lemicraft

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultBaseURL = "https://lemicraft.ru/api"

// Client is the root API client. All resource services hang off it.
// Create one with New() and access endpoints via its service fields:
//
//	client := lemicraft.New()
//	user, err := client.Users.GetByDiscordID(ctx, "123456789")
type Client struct {
	baseURL    string
	httpClient *http.Client

	// Services — add new resource services here as the API grows.
	Users *UsersService
}

// New creates a Client with sensible defaults. Customize it with Option functions.
func New(opts ...Option) *Client {
	c := &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	// Wire up all services after options have been applied.
	c.Users = &UsersService{client: c}

	return c
}

// do executes an HTTP request and decodes a successful JSON body into dst.
// It returns *APIError for any non-2xx status code.
func (c *Client) do(req *http.Request, dst any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("lemicraft: http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("lemicraft: reading response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	if dst != nil {
		if err := json.Unmarshal(body, dst); err != nil {
			return fmt.Errorf("lemicraft: decoding response: %w", err)
		}
	}

	return nil
}
