package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	lemicraft "github.com/Mr1rbis/lemicraft-go"
)

func main() {
	// Get API token from command line or environment
	tokenFlag := flag.String("token", "", "Lemicraft API token (from https://lemicraft.ru/settings)")
	flag.Parse()

	if *tokenFlag == "" {
		log.Fatal("Please provide API token: -token=your-token-here")
	}

	// ── Create authenticated client ──────────────────────────────────────────
	client := lemicraft.New(*tokenFlag)

	// ── Optional: customize with options ─────────────────────────────────────
	// client := lemicraft.New(
	// 	*tokenFlag,
	// 	lemicraft.WithBaseURL("https://staging.lemicraft.ru/api"),
	// 	lemicraft.WithHTTPClient(&http.Client{Timeout: 15 * time.Second}),
	// )

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ── Example 1: Get player list ──────────────────────────────────────────
	fmt.Println("\n=== Players ===")
	players, err := client.Players.List(ctx, nil)
	if err != nil {
		handleError(err)
		return
	}
	fmt.Printf("Total players: %d\n", players.Total)
	if len(players.Players) > 0 {
		fmt.Printf("First player: %s (UUID: %v)\n", players.Players[0].Name, players.Players[0].UUID)
	}

	// ── Example 2: Get specific player ──────────────────────────────────────
	fmt.Println("\n=== Player Profile ===")
	playerProfile, err := client.Players.GetByNick(ctx, "Notch")
	if err != nil {
		fmt.Printf("Note: %v (expected if player doesn't exist)\n", err)
	} else {
		fmt.Printf("Player: %s\n", playerProfile.Name)
		fmt.Printf("Banned: %v\n", playerProfile.Banned)
		fmt.Printf("Avatar URL: %s\n", playerProfile.AvatarURL)
	}

	// ── Example 3: Get launcher version (public endpoint) ────────────────────
	fmt.Println("\n=== Launcher Version ===")
	version, err := client.Launcher.GetVersion(ctx)
	if err != nil {
		fmt.Printf("Note: %v\n", err)
	} else {
		fmt.Printf("Version: %s\n", version.Version)
		fmt.Printf("Download: %s\n", version.DownloadURL)
		fmt.Printf("Released: %s\n", version.ReleaseDate.Format(time.RFC3339))
	}

	// ── Example 4: Get site news ────────────────────────────────────────────
	fmt.Println("\n=== Site News ===")
	news, err := client.News.GetNews(ctx, nil)
	if err != nil {
		fmt.Printf("Note: %v\n", err)
	} else {
		fmt.Printf("News items: %d\n", len(news.Items))
		if len(news.Items) > 0 {
			fmt.Printf("First news: %s\n", news.Items[0].AuthorName)
			fmt.Printf("Has more: %v\n", news.HasMore)
		}
	}

	// ── Example 5: Get petitions ────────────────────────────────────────────
	fmt.Println("\n=== Petitions ===")
	petitions, err := client.Petitions.List(ctx, nil)
	if err != nil {
		fmt.Printf("Note: %v\n", err)
	} else {
		fmt.Printf("Petitions: %d\n", len(petitions.Petitions))
	}

	// ── Example 6: Get court cases ──────────────────────────────────────────
	fmt.Println("\n=== Court Cases ===")
	cases, err := client.Court.List(ctx, nil)
	if err != nil {
		fmt.Printf("Note: %v\n", err)
	} else {
		fmt.Printf("Court cases: %d\n", len(cases.Cases))
	}

	fmt.Println("\n=== All examples completed ===")
}

func handleError(err error) {
	// 1. Check for not found — most common API error
	if errors.Is(err, lemicraft.ErrNotFound) {
		log.Fatal("Resource not found (404)")
	}

	// 2. Extract full API error details
	var apiErr *lemicraft.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.StatusCode {
		case 400:
			log.Fatalf("Bad request: %s", apiErr.Message)
		case 401:
			log.Fatalf("Unauthorized (invalid token): %s", apiErr.Message)
		case 403:
			log.Fatalf("Forbidden: %s", apiErr.Message)
		case 404:
			log.Fatalf("Not found: %s", apiErr.Message)
		case 429:
			log.Fatalf("Rate limited: %s", apiErr.Message)
		case 500, 502, 503, 504:
			log.Fatalf("Server error (%d): %s", apiErr.StatusCode, apiErr.Message)
		default:
			log.Fatalf("API error (%d): %s", apiErr.StatusCode, apiErr.Message)
		}
	}

	// 3. Network / transport errors
	log.Fatalf("Request error: %v", err)
}
