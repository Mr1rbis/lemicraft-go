package lemicraft

import (
	"context"
	"fmt"
	"net/http"
)

// CourtService handles court-related endpoints.
type CourtService struct {
	client *Client
}

// ListOptions contains optional parameters for listing court cases
type CourtListOptions struct {
	Status *string // open, closed
}

// List fetches all court cases.
//
// GET /api/court
func (s *CourtService) List(ctx context.Context, opts *CourtListOptions) (*CourtCasesResponse, error) {
	endpoint := s.client.baseURL + "/court"

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

	var result CourtCasesResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Get fetches a single court case by ID.
//
// GET /api/court/{id}
func (s *CourtService) Get(ctx context.Context, id int) (*CourtCase, error) {
	endpoint := fmt.Sprintf("%s/court/%d", s.client.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var result CourtCaseResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result.Case, nil
}

// GetMessages fetches all messages (arguments) in a court case.
//
// GET /api/court/messages/{id}
func (s *CourtService) GetMessages(ctx context.Context, caseID int) (*CourtMessagesResponse, error) {
	endpoint := fmt.Sprintf("%s/court/messages/%d", s.client.baseURL, caseID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var result CourtMessagesResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
