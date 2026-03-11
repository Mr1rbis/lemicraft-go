package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	lemicraft "github.com/Mr1rbis/lemicraft-go"
)

func main() {
	// ── Unauthenticated client (default) ──────────────────────────────────────
	client := lemicraft.New()

	// ── Authenticated client example (uncomment when you have a token) ────────
	// client := lemicraft.New(
	// 	lemicraft.WithAuthToken("your-api-token-here"),
	// )

	// ── Custom base URL + timeout example ─────────────────────────────────────
	// client := lemicraft.New(
	// 	lemicraft.WithBaseURL("https://staging.lemicraft.ru/api"),
	// 	lemicraft.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	// )

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	const discordID = "535868441433735188"

	user, err := client.Users.GetByDiscordID(ctx, discordID)
	if err != nil {
		handleError(err)
		return
	}

	fmt.Printf("User found:\n")
	fmt.Printf("  Discord ID:       %s\n", user.DiscordID)
	fmt.Printf("  Discord Username: %s\n", user.DiscordUsername)
	fmt.Printf("  Minecraft Nick:   %s\n", user.MinecraftNick)
	fmt.Printf("  Minecraft UUID:   %s\n", user.MinecraftUUID)
	fmt.Printf("  Whitelisted:      %v\n", user.Whitelisted)
	fmt.Printf("  Nick Source:      %s\n", user.NickSource)
}

func handleError(err error) {
	// 1. Проверяем конкретный sentinel — пользователь не найден.
	if errors.Is(err, lemicraft.ErrNotFound) {
		log.Fatal("пользователь не найден (404)")
	}

	// 2. Для остальных HTTP-ошибок достаём полный APIError со статусом и телом.
	var apiErr *lemicraft.APIError
	if errors.As(err, &apiErr) {
		switch {
		case apiErr.StatusCode == 401:
			log.Fatalf("не авторизован — передайте токен через WithAuthToken: %s", apiErr.Message)
		case apiErr.StatusCode == 403:
			log.Fatalf("доступ запрещён: %s", apiErr.Message)
		case apiErr.StatusCode >= 500:
			log.Fatalf("ошибка сервера (%d): %s", apiErr.StatusCode, apiErr.Message)
		default:
			log.Fatalf("API вернул ошибку (%d): %s", apiErr.StatusCode, apiErr.Message)
		}
	}

	// 3. Сетевые / таймаут / прочие ошибки транспорта.
	log.Fatalf("ошибка запроса: %v", err)
}
