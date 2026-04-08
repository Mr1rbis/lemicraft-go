package lemicraft

import "net/http"

// Option is a functional option for configuring a Client.
// This pattern keeps the constructor clean and makes it trivial to add new
// configuration knobs without breaking existing call sites.
type Option func(*Client)

// WithBaseURL overrides the default API base URL.
// Useful for pointing at a staging or local server during development.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient replaces the default http.Client.
// Use this to set timeouts, custom TLS config, or a pre-configured client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}
