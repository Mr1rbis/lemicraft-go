package lemicraft

import "net/http"

// authTransport wraps an http.RoundTripper and injects an Authorization header
// on every outgoing request. Replacing this single type is enough to support any
// auth scheme (Bearer token, API key, HMAC, etc.) across all endpoints at once.
type authTransport struct {
	token     string
	transport http.RoundTripper
}

// RoundTrip clones the request, attaches the Authorization header, and delegates
// to the underlying transport so that the original request is never mutated.
func (a *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", "Bearer "+a.token)
	return a.transport.RoundTrip(clone)
}
