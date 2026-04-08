package lemicraft

import (
	"context"
	"fmt"
	"net/http"
)

// NewsService handles content endpoints: news, gallery, posts.
type NewsService struct {
	client *Client
}

// GetNews fetches site news from Discord with full HTML rendering.
// Supports cursor-based pagination.
//
// GET /api/news
func (s *NewsService) GetNews(ctx context.Context, beforeID *string) (*NewsResponse, error) {
	endpoint := s.client.baseURL + "/news"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	// Add query parameters if provided
	if beforeID != nil {
		q := req.URL.Query()
		q.Set("before", *beforeID)
		req.URL.RawQuery = q.Encode()
	}

	var result NewsResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGallery fetches approved screenshots from the gallery.
//
// GET /api/gallery
func (s *NewsService) GetGallery(ctx context.Context) (*GalleryResponse, error) {
	endpoint := s.client.baseURL + "/gallery"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var result GalleryResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// PostsOptions contains optional parameters for fetching posts
type PostsOptions struct {
	Before *int // Load posts older than this ID (for pagination)
}

// GetPosts fetches community posts with pagination support.
//
// GET /api/posts
func (s *NewsService) GetPosts(ctx context.Context, opts *PostsOptions) (*PostsResponse, error) {
	endpoint := s.client.baseURL + "/posts"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	// Add query parameters if provided
	if opts != nil && opts.Before != nil {
		q := req.URL.Query()
		q.Set("before", fmt.Sprint(*opts.Before))
		req.URL.RawQuery = q.Encode()
	}

	var result PostsResponse
	if err := s.client.do(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPost fetches a single post by ID with all comments.
//
// GET /api/posts/{id}
func (s *NewsService) GetPost(ctx context.Context, id int) (*Post, error) {
	endpoint := fmt.Sprintf("%s/posts/%d", s.client.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("lemicraft: building request: %w", err)
	}

	var post Post
	if err := s.client.do(req, &post); err != nil {
		return nil, err
	}

	return &post, nil
}
