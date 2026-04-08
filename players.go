package lemicraft

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// PlayersService handles all /api/players/* endpoints.
type PlayersService struct {
	client *Client
}

// ListOptions contains optional parameters for listing players
type ListOptions struct {
	Search *string
}

// List fetches all players from the whitelist with optional search.
//
// GET /api/players
func (s *PlayersService) List(ctx context.Context, opts *ListOptions) (*PlayersListResponse, error) {
	endpoint := s.client.baseURL + "/players"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	// Add query parameters if provided
	if opts != nil && opts.Search != nil {
		q := req.URL.Query()
		q.Set("search", *opts.Search)
		req.URL.RawQuery = q.Encode()
	}

	var result PlayersListResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetByNick fetches detailed info about a player by their nickname.
//
// GET /api/players/{nick}
func (s *PlayersService) GetByNick(ctx context.Context, nick string) (*PlayerFull, error) {
	endpoint := fmt.Sprintf("%s/players/%s", s.client.baseURL, url.PathEscape(nick))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var player PlayerFull
	if err := s.client.do(req, &player); err != nil {
		return nil, err
	}

	return &player, nil
}

// GetByDiscord retrieves detailed information about a player based on their Discord ID.
//
// GET /api/users/discord/{discordid}
func (s *PlayersService) GetByDiscord(ctx context.Context, discordid string) (*PlayerByDiscordId, error) {
	endpoint := fmt.Sprintf("%s/users/discord/%s", s.client.baseURL, url.PathEscape(discordid))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var player PlayerByDiscordId
	if err := s.client.do(req, &player); err != nil {
		return nil, err
	}

	return &player, nil
}

// GetPlan fetches player statistics from the Plan plugin.
//
// GET /api/plan/{nick}
func (s *PlayersService) GetPlan(ctx context.Context, nick string) (*PlayerPlan, error) {
	endpoint := fmt.Sprintf("%s/plan/%s", s.client.baseURL, url.PathEscape(nick))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var plan PlayerPlan
	if err := s.client.do(req, &plan); err != nil {
		return nil, err
	}

	return &plan, nil
}

// GetAvatar fetches a player's avatar (128x128 PNG).
// Returns the raw image data as []byte.
//
// GET /api/avatar/{nick}
func (s *PlayersService) GetAvatar(ctx context.Context, nick string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/avatar/%s", s.client.baseURL, url.PathEscape(nick))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to fetch avatar",
		}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: reading response body: %w", err)
	}

	return data, nil
}

// GetSkin fetches a player's skin (64x64 PNG).
// Returns the raw image data as []byte.
//
// GET /api/skin/{nick}
func (s *PlayersService) GetSkin(ctx context.Context, nick string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s/skin/%s", s.client.baseURL, url.PathEscape(nick))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: http request failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message:    "failed to fetch skin",
		}
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: reading response body: %w", err)
	}

	return data, nil
}
