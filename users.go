package lemicraft

import (
	"context"
	"fmt"
	"net/http"
)

// UsersService handles all /api/users/* endpoints.
// Add new user-related methods here as the API grows.
type UsersService struct {
	client *Client
}

// GetByDiscordID fetches a Lemicraft user by their Discord snowflake ID.
//
//	GET /api/users/discord/<discordID>
func (s *UsersService) GetByDiscordID(ctx context.Context, discordID string) (*UserByDiscord, error) {
	url := fmt.Sprintf("%s/users/discord/%s", s.client.baseURL, discordID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	var user UserByDiscord
	if err := s.client.do(req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
