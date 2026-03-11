package lemicraft

import (
	"errors"
	"fmt"
)

// UserByDiscord represents a Lemicraft user fetched by their Discord ID.
type UserByDiscord struct {
	DiscordID       string `json:"discord_id"`
	DiscordUsername string `json:"discord_username"`
	MinecraftNick   string `json:"minecraft_nick"`
	MinecraftUUID   string `json:"minecraft_uuid"`
	Whitelisted     bool   `json:"whitelisted"`
	NickSource      string `json:"nick_source"`
}

// APIError represents a non-2xx HTTP response returned by the Lemicraft API.
// Use errors.As(err, &apiErr) to inspect StatusCode and Message.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("lemicraft API error: status %d — %s", e.StatusCode, e.Message)
}

// Is allows errors.Is(err, lemicraft.ErrNotFound) to work transparently
// when the underlying error is an *APIError with StatusCode 404.
func (e *APIError) Is(target error) bool {
	return target == ErrNotFound && e.StatusCode == 404
}

// ErrNotFound is returned (via errors.Is) when the API responds with HTTP 404.
var ErrNotFound = errors.New("lemicraft: not found")
