package lemicraft

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultBaseURL = "https://lemicraft.ru/api"

// Client is the root API client. All resource services hang off it.
// Create one with New() and access endpoints via its service fields.
// An API token is required to access most endpoints.
//
// Example:
//
//	client := lemicraft.New("your-api-token")
//	players, err := client.Players.List(ctx, nil)
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string

	// Services
	Players   *PlayersService
	Launcher  *LauncherService
	News      *NewsService
	Petitions *PetitionsService
	Court     *CourtService
}

// New creates a Client with the provided API token (required).
// The token can be obtained from https://lemicraft.ru/settings
func New(token string, opts ...Option) *Client {
	if token == "" {
		panic("lemicraft: API token is required")
	}

	c := &Client{
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{},
		token:      token,
	}

	for _, opt := range opts {
		opt(c)
	}

	// Wire up all services after options have been applied
	c.Players = &PlayersService{client: c}
	c.Launcher = &LauncherService{client: c}
	c.News = &NewsService{client: c}
	c.Petitions = &PetitionsService{client: c}
	c.Court = &CourtService{client: c}

	return c
}

// do executes an HTTP request and decodes a successful JSON body into dst.
// It automatically adds the Authorization header and returns *APIError for non-2xx status.
func (c *Client) do(req *http.Request, dst any) error {
	// Add Authorization header
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

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

// doUnauth executes an HTTP request without authentication (for public endpoints like /launcher/*).
func (c *Client) doUnauth(req *http.Request, dst any) error {
	req.Header.Set("Accept", "application/json")

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
