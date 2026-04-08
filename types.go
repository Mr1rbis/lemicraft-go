package lemicraft

import (
	"errors"
	"fmt"
	"time"
)

// ── Error types ──────────────────────────────────────────────────────────────

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

// ── Players ──────────────────────────────────────────────────────────────────

// PlayerShort contains basic player info (from /api/players)
type PlayerShort struct {
	Name      string  `json:"name"`
	UUID      *string `json:"uuid"`
	ServiceID int     `json:"serviceId"` // 0=Mojang, 1=Ely.by
	Banned    bool    `json:"banned"`
}

// PlayerFull contains detailed player info (from /api/players/{nick})
type PlayerFull struct {
	Name         string     `json:"name"`
	UUID         *string    `json:"uuid"`
	ServiceID    int        `json:"serviceId"`
	Banned       bool       `json:"banned"`
	BanReason    *string    `json:"banReason"`
	BanPermanent bool       `json:"banPermanent"`
	BanEndsAt    *time.Time `json:"banEndsAt"`
	AvatarURL    string     `json:"avatarUrl"`
	SkinURL      string     `json:"skinUrl"`
}

// PlayerPlan contains player statistics from Plan plugin
type PlayerPlan struct {
	Registered     *time.Time `json:"registered"`
	LastSeen       *time.Time `json:"lastSeen"`
	Online         bool       `json:"online"`
	Playtime       *int64     `json:"playtime"`       // milliseconds
	ActivePlaytime *int64     `json:"activePlaytime"` // milliseconds
	PlaytimeStr    string     `json:"playtimeStr"`
	Sessions       *int       `json:"sessions"`
	LongestSession *int64     `json:"longestSession"` // milliseconds
	Deaths         *int       `json:"deaths"`
	MobKills       *int       `json:"mobKills"`
	PlayerKills    *int       `json:"playerKills"`
	Ping           *float64   `json:"ping"`
}

// PlayersListResponse is the response from GET /api/players
type PlayersListResponse struct {
	Players []PlayerShort `json:"players"`
	Total   int           `json:"total"`
}

// ── Launcher ─────────────────────────────────────────────────────────────────

// LauncherVersion contains launcher release info (from /api/launcher/version)
type LauncherVersion struct {
	Version     string    `json:"version"`
	DownloadURL string    `json:"downloadUrl"`
	FileName    string    `json:"fileName"`
	FileSize    int64     `json:"fileSize"`
	ReleaseDate time.Time `json:"releaseDate"`
}

// ModpackVersion contains modpack info (from /api/launcher/modpack/version)
type ModpackVersion struct {
	Success     bool   `json:"success"`
	Version     string `json:"version"`
	FileName    string `json:"fileName"`
	DownloadURL string `json:"downloadUrl"`
	FileSize    int64  `json:"fileSize"`
	Name        string `json:"name"`
}

// LauncherNewsItem contains launcher news (from /api/launcher/news)
type LauncherNewsItem struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Preview         string    `json:"preview"`
	Content         string    `json:"content"`
	ImageURL        *string   `json:"imageUrl"`
	AuthorName      string    `json:"authorName"`
	AuthorRole      string    `json:"authorRole"`
	AuthorAvatarURL string    `json:"authorAvatarUrl"`
	PublishedAt     time.Time `json:"publishedAt"`
	Category        string    `json:"category"` // general, update, event, maintenance, announcement
	URL             string    `json:"url"`
}

// LauncherNewsResponse is the response from GET /api/launcher/news
type LauncherNewsResponse struct {
	Success bool               `json:"success"`
	Items   []LauncherNewsItem `json:"items"`
	Total   int                `json:"total"`
}

// ── Content ──────────────────────────────────────────────────────────────────

// SiteNewsItem contains site news (from /api/news)
type SiteNewsItem struct {
	ID              string    `json:"id"`
	HTML            string    `json:"html"`
	Images          []string  `json:"images"`
	AuthorName      string    `json:"authorName"`
	AuthorRole      string    `json:"authorRole"`
	AuthorAvatarURL *string   `json:"authorAvatarUrl"`
	PublishedAt     time.Time `json:"publishedAt"`
	URL             string    `json:"url"`
}

// NewsResponse is the response from GET /api/news
type NewsResponse struct {
	Items   []SiteNewsItem `json:"items"`
	HasMore bool           `json:"hasMore"`
	LastID  *string        `json:"lastId"`
	Stale   bool           `json:"stale"`
}

// GalleryImage represents a gallery image
type GalleryImage struct {
	ID       int    `json:"id"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

// GalleryResponse is the response from GET /api/gallery
type GalleryResponse struct {
	Images []GalleryImage `json:"images"`
}

// Comment represents a post comment
type Comment struct {
	ID              int       `json:"id"`
	AuthorDiscordID string    `json:"authorDiscordId"`
	AuthorName      string    `json:"authorName"`
	Text            string    `json:"text"`
	CreatedAt       time.Time `json:"createdAt"`
}

// Post represents a community post
type Post struct {
	ID              int       `json:"id"`
	AuthorDiscordID string    `json:"authorDiscordId"`
	AuthorName      string    `json:"authorName"`
	Text            string    `json:"text"`
	Edited          bool      `json:"edited"`
	CreatedAt       time.Time `json:"createdAt"`
	Media           []string  `json:"media"`
	VotesFor        int       `json:"votesFor"`
	VotesAgainst    int       `json:"votesAgainst"`
	UserVote        *int      `json:"userVote"` // 1, -1, or null
	Comments        []Comment `json:"comments"`
	CommentCount    int       `json:"commentCount"`
}

// PostsResponse is the response from GET /api/posts
type PostsResponse struct {
	Items   []Post `json:"items"`
	HasMore bool   `json:"hasMore"`
	LastID  *int   `json:"lastId"`
}

// ── Community ────────────────────────────────────────────────────────────────

// Petition represents a server petition
type Petition struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Author       string    `json:"author"`
	Status       string    `json:"status"` // active, closed, blocked
	VotesFor     int       `json:"votes_for"`
	VotesAgainst int       `json:"votes_against"`
	UserVote     *int      `json:"userVote"` // 1, -1, or null
	CreatedAt    time.Time `json:"created_at"`
}

// PetitionsResponse is the response from GET /api/petitions
type PetitionsResponse struct {
	Petitions []Petition `json:"petitions"`
}

// CourtCase represents a court case
type CourtCase struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Defendant   string    `json:"defendant"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // open, closed
	Author      string    `json:"author"`
	Judge       *string   `json:"judge"`
	Verdict     *string   `json:"verdict"`
	CreatedAt   time.Time `json:"created_at"`
}

// CourtCaseResponse is the response from GET /api/court/{id}
type CourtCaseResponse struct {
	Case CourtCase `json:"case"`
}

// CourtCasesResponse is the response from GET /api/court
type CourtCasesResponse struct {
	Cases []CourtCase `json:"cases"`
}

// CourtMessage represents a message in a court case
type CourtMessage struct {
	ID        int       `json:"id"`
	CaseID    int       `json:"caseId"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// CourtMessagesResponse is the response from GET /api/court/messages/{id}
type CourtMessagesResponse struct {
	Messages []CourtMessage `json:"messages"`
}
