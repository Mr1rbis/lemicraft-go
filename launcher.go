package lemicraft

import (
	"context"
	"fmt"
	"net/http"
)

// LauncherService handles all /api/launcher/* endpoints.
// Note: These endpoints do not require authentication.
type LauncherService struct {
	client *Client
}

// GetVersion fetches the latest launcher version.
//
// GET /api/launcher/version
// This endpoint does not require authentication.
func (s *LauncherService) GetVersion(ctx context.Context) (*LauncherVersion, error) {
	endpoint := s.client.baseURL + "/launcher/version"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var version LauncherVersion
	// Use doUnauth since this endpoint is public
	if err := s.client.doUnauth(req, &version); err != nil {
		return nil, err
	}

	return &version, nil
}

// GetModpackVersion fetches the latest modpack version.
//
// GET /api/launcher/modpack/version
// This endpoint does not require authentication.
func (s *LauncherService) GetModpackVersion(ctx context.Context) (*ModpackVersion, error) {
	endpoint := s.client.baseURL + "/launcher/modpack/version"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var modpack ModpackVersion
	// Use doUnauth since this endpoint is public
	if err := s.client.doUnauth(req, &modpack); err != nil {
		return nil, err
	}

	return &modpack, nil
}

// NewsOptions contains optional parameters for fetching launcher news
type NewsOptions struct {
	Limit    *int    // Maximum 50
	Category *string // general, update, event, maintenance, announcement
}

// GetNews fetches news items for the launcher.
//
// GET /api/launcher/news
// This endpoint does not require authentication.
func (s *LauncherService) GetNews(ctx context.Context, opts *NewsOptions) (*LauncherNewsResponse, error) {
	endpoint := s.client.baseURL + "/launcher/news"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	// Add query parameters if provided
	if opts != nil {
		q := req.URL.Query()
		if opts.Limit != nil {
			q.Set("limit", fmt.Sprint(*opts.Limit))
		}
		if opts.Category != nil {
			q.Set("category", *opts.Category)
		}
		req.URL.RawQuery = q.Encode()
	}

	var result LauncherNewsResponse
	// Use doUnauth since this endpoint is public
	if err := s.client.doUnauth(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
