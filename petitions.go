package lemicraft

import (
	"context"
	"fmt"
	"net/http"
)

// PetitionsService handles petition-related endpoints.
type PetitionsService struct {
	client *Client
}

// ListOptions contains optional parameters for listing petitions
type PetitionsListOptions struct {
	Status *string // active, closed, blocked (default: active)
}

// List fetches all petitions.
//
// GET /api/petitions
func (s *PetitionsService) List(ctx context.Context, opts *PetitionsListOptions) (*PetitionsResponse, error) {
	endpoint := s.client.baseURL + "/petitions"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	// Add query parameters if provided
	if opts != nil && opts.Status != nil {
		q := req.URL.Query()
		q.Set("status", *opts.Status)
		req.URL.RawQuery = q.Encode()
	}

	var result PetitionsResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
