package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	lemicraft "github.com/Mr1rbis/lemicraft-go"
)

const exampleNick = "MrIrbis"
const exampleDiscordid = "363220369869504512"

func main() {
	tokenFlag := flag.String("token", "", "Lemicraft API token (from https://lemicraft.ru/settings)")
	flag.Parse()

	if *tokenFlag == "" {
		log.Fatal("❌ Please provide API token: -token=your-token-here")
	}

	// Create authenticated client
	client := lemicraft.New(*tokenFlag)
	ctx := context.Background()

	fmt.Println("🎮 LemiCraft Go Library - Примеры для всех эндпоинтов\n")

	// ────────────────────────────────────────────────────────────────────────
	// PLAYERS SERVICE
	// ────────────────────────────────────────────────────────────────────────

	fmt.Println("=" + stringRepeat("=", 69))
	fmt.Println("📋 PLAYERS SERVICE")
	fmt.Println("=" + stringRepeat("=", 69))

	// 1. List all players
	exampleListPlayers(ctx, client)

	// 2. Search players
	exampleSearchPlayers(ctx, client)

	// 3. Get player profile
	exampleGetPlayerProfile(ctx, client)

	// 4. Get user by discordID
	exampleGetPlayerByDiscordId(ctx, client)

	// 5. Get player statistics
	exampleGetPlayerStats(ctx, client)

	// 6. Download avatar
	exampleDownloadAvatar(ctx, client)

	// 7. Download skin
	exampleDownloadSkin(ctx, client)

	// ────────────────────────────────────────────────────────────────────────
	// LAUNCHER SERVICE (Public endpoints)
	// ────────────────────────────────────────────────────────────────────────

	fmt.Println("\n" + "=" + stringRepeat("=", 69))
	fmt.Println("🚀 LAUNCHER SERVICE (публичные эндпоинты)")
	fmt.Println("=" + stringRepeat("=", 69))

	// 8. Get launcher version
	exampleGetLauncherVersion(ctx, client)

	// 9. Get modpack version
	exampleGetModpackVersion(ctx, client)

	// 10. Get launcher news
	exampleGetLauncherNews(ctx, client)

	// ────────────────────────────────────────────────────────────────────────
	// NEWS SERVICE
	// ────────────────────────────────────────────────────────────────────────

	fmt.Println("\n" + "=" + stringRepeat("=", 69))
	fmt.Println("📰 NEWS SERVICE")
	fmt.Println("=" + stringRepeat("=", 69))

	// 11. Get site news
	exampleGetSiteNews(ctx, client)

	// 12. Get gallery
	exampleGetGallery(ctx, client)

	// 13. Get posts
	exampleGetPosts(ctx, client)

	// 14. Get single post
	exampleGetSinglePost(ctx, client)

	// ────────────────────────────────────────────────────────────────────────
	// PETITIONS SERVICE
	// ────────────────────────────────────────────────────────────────────────

	fmt.Println("\n" + "=" + stringRepeat("=", 69))
	fmt.Println("✋ PETITIONS SERVICE")
	fmt.Println("=" + stringRepeat("=", 69))

	// 15. Get petitions
	exampleGetPetitions(ctx, client)

	// ────────────────────────────────────────────────────────────────────────
	// COURT SERVICE
	// ────────────────────────────────────────────────────────────────────────

	fmt.Println("\n" + "=" + stringRepeat("=", 69))
	fmt.Println("⚖️  COURT SERVICE")
	fmt.Println("=" + stringRepeat("=", 69))

	// 16. Get court cases
	exampleGetCourtCases(ctx, client)

	// 17. Get court case
	exampleGetCourtCase(ctx, client)

	// 18. Get court messages
	exampleGetCourtMessages(ctx, client)

	fmt.Println("\n✅ Все примеры завершены!")
}

// ============================================================================
// PLAYERS SERVICE EXAMPLES
// ============================================================================

func exampleListPlayers(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 1: Получить список всех игроков")
	fmt.Println("   GET /api/players")

	players, err := client.Players.List(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Игроков найдено: %d\n", players.Total)
	if len(players.Players) > 0 {
		for i, p := range players.Players {
			if i >= 3 {
				fmt.Printf("   ... и еще %d\n", len(players.Players)-3)
				break
			}
			fmt.Printf("   - %s (UUID: %v, Забанен: %v)\n", p.Name, p.UUID, p.Banned)
		}
	}
}

func exampleSearchPlayers(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 2: Поиск игроков по нику")
	fmt.Println("   GET /api/players?search=Steve")

	search := "Steve"
	players, err := client.Players.List(ctx, &lemicraft.ListOptions{
		Search: &search,
	})
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Найдено игроков: %d\n", players.Total)
	for _, p := range players.Players {
		fmt.Printf("   - %s\n", p.Name)
	}
}

func exampleGetPlayerProfile(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 3: Получить профиль игрока")
	fmt.Println("   GET /api/players/{nick}")

	player, err := client.Players.GetByNick(ctx, exampleNick)
	if err != nil {
		// Это нормально если игрока нет - он может не существовать
		if errors.Is(err, lemicraft.ErrNotFound) {
			fmt.Printf("   ✗ Игрок %s не найден (404)\n", exampleNick)
		} else {
			printError(err)
		}
		return
	}

	fmt.Printf("   ✓ Ник: %s\n", player.Name)
	fmt.Printf("   - UUID: %v\n", player.UUID)
	fmt.Printf("   - Забанен: %v\n", player.Banned)
	if player.BanReason != nil {
		fmt.Printf("   - Причина бана: %s\n", *player.BanReason)
	}
	fmt.Printf("   - Аватар: %s\n", player.AvatarURL)
	fmt.Printf("   - Скин: %s\n", player.SkinURL)
}

func exampleGetPlayerByDiscordId(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 4: Получить профиль игрока")
	fmt.Println("   GET /api/users/discord/{discordid}")

	player, err := client.Players.GetByDiscord(ctx, exampleDiscordid)
	if err != nil {
		// Это нормально если игрока нет - он может не существовать
		if errors.Is(err, lemicraft.ErrNotFound) {
			fmt.Printf("   ✗ Игрок %s не найден (404)\n", exampleDiscordid)
		} else {
			printError(err)
		}
		return
	}

	fmt.Printf("   ✓ DiscordID: %s\n", player.DiscordId)
	fmt.Printf("   - Nick: %v\n", player.MinecraftNick)
	fmt.Printf("   - DiscordUsername: %v\n", player.DiscordUsername)
	fmt.Printf("   - UUID: %s\n", player.MinecraftUuid)
	fmt.Printf("   - Nick source: %s\n", player.NickSource)
	fmt.Printf("   - В вайтлисте: %v\n", player.Whitelisted)

}

func exampleGetPlayerStats(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 5: Получить статистику игрока (Plan)")
	fmt.Println("   GET /api/plan/{nick}")

	plan, err := client.Players.GetPlan(ctx, exampleNick)
	if err != nil {
		if errors.Is(err, lemicraft.ErrNotFound) {
			fmt.Printf("   ✗ Статистика для %s не найдена (игрок может быть неактивен)\n", exampleNick)
		} else {
			printError(err)
		}
		return
	}

	fmt.Printf("   ✓ Статистика игрока:\n")
	fmt.Printf("   - Онлайн: %v\n", plan.Online)
	fmt.Printf("   - Время игры: %s\n", plan.PlaytimeStr)
	if plan.Playtime != nil {
		fmt.Printf("   - Общее время (мс): %d\n", *plan.Playtime)
	}
	if plan.Sessions != nil {
		fmt.Printf("   - Сессии: %d\n", *plan.Sessions)
	}
	if plan.MobKills != nil {
		fmt.Printf("   - Убийств мобов: %d\n", *plan.MobKills)
	}
	if plan.Deaths != nil {
		fmt.Printf("   - Смертей: %d\n", *plan.Deaths)
	}
}

func exampleDownloadAvatar(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 6: Скачать аватар игрока")
	fmt.Println("   GET /api/avatar/{nick}")

	avatar, err := client.Players.GetAvatar(ctx, exampleNick)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Аватар скачан: %d байт\n", len(avatar))
	// В реальном коде можно сохранить:
	// ioutil.WriteFile("avatar.png", avatar, 0644)
}

func exampleDownloadSkin(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 7: Скачать скин игрока")
	fmt.Println("   GET /api/skin/{nick}")

	skin, err := client.Players.GetSkin(ctx, exampleNick)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Скин скачан: %d байт\n", len(skin))
	// В реальном коде можно сохранить:
	// ioutil.WriteFile("skin.png", skin, 0644)
}

// ============================================================================
// LAUNCHER SERVICE EXAMPLES (Public endpoints)
// ============================================================================

func exampleGetLauncherVersion(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 8: Получить версию лаунчера")
	fmt.Println("   GET /api/launcher/version (публичный)")

	version, err := client.Launcher.GetVersion(ctx)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Версия: %s\n", version.Version)
	fmt.Printf("   - Скачать: %s\n", version.DownloadURL)
	fmt.Printf("   - Файл: %s\n", version.FileName)
	fmt.Printf("   - Размер: %d MB\n", version.FileSize/1024/1024)
	fmt.Printf("   - Дата: %s\n", version.ReleaseDate.Format("2006-01-02"))
}

func exampleGetModpackVersion(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 9: Получить версию модпака")
	fmt.Println("   GET /api/launcher/modpack/version (публичный)")

	modpack, err := client.Launcher.GetModpackVersion(ctx)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Модпак: %s\n", modpack.Name)
	fmt.Printf("   - Версия: %s\n", modpack.Version)
	fmt.Printf("   - Скачать: %s\n", modpack.DownloadURL)
	fmt.Printf("   - Размер: %d MB\n", modpack.FileSize/1024/1024)
}

func exampleGetLauncherNews(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 10: Получить новости для лаунчера")
	fmt.Println("   GET /api/launcher/news?limit=10&category=update (публичный)")

	limit := 10
	category := "update"
	news, err := client.Launcher.GetNews(ctx, &lemicraft.NewsOptions{
		Limit:    &limit,
		Category: &category,
	})
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Новостей найдено: %d\n", news.Total)
	for i, item := range news.Items {
		if i >= 3 {
			fmt.Printf("   ... и еще %d новостей\n", news.Total-3)
			break
		}
		fmt.Printf("   - [%s] %s (автор: %s)\n", item.Category, item.Title, item.AuthorName)
	}
}

// ============================================================================
// NEWS SERVICE EXAMPLES
// ============================================================================

func exampleGetSiteNews(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 11: Получить новости сайта")
	fmt.Println("   GET /api/news")

	news, err := client.News.GetNews(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Новостей найдено: %d\n", len(news.Items))
	fmt.Printf("   - Есть еще страницы: %v\n", news.HasMore)
	if len(news.Items) > 0 {
		fmt.Printf("   - Автор первой новости: %s (%s)\n",
			news.Items[0].AuthorName, news.Items[0].AuthorRole)
		// Для получения следующей страницы:
		// if news.HasMore && news.LastID != nil {
		//     nextNews, _ := client.News.GetNews(ctx, news.LastID)
		// }
	}
}

func exampleGetGallery(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 12: Получить галерею")
	fmt.Println("   GET /api/gallery")

	gallery, err := client.News.GetGallery(ctx)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Изображений в галерее: %d\n", len(gallery.Images))
	for i, img := range gallery.Images {
		if i >= 3 {
			fmt.Printf("   ... и еще %d изображений\n", len(gallery.Images)-3)
			break
		}
		fmt.Printf("   - ID %d: %s (%s)\n", img.ID, img.Filename, img.URL)
	}
}

func exampleGetPosts(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 13: Получить посты")
	fmt.Println("   GET /api/posts")

	posts, err := client.News.GetPosts(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Постов найдено: %d\n", len(posts.Items))
	fmt.Printf("   - Есть еще страницы: %v\n", posts.HasMore)
	if len(posts.Items) > 0 {
		post := posts.Items[0]
		fmt.Printf("   - Автор: %s\n", post.AuthorName)
		fmt.Printf("   - Текст: %s...\n", truncate(post.Text, 50))
		fmt.Printf("   - Лайков: %d, Дизлайков: %d\n", post.VotesFor, post.VotesAgainst)
		fmt.Printf("   - Комментариев: %d (показано %d превью)\n", post.CommentCount, len(post.Comments))
	}
}

func exampleGetSinglePost(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 14: Получить один пост со всеми комментариями")
	fmt.Println("   GET /api/posts/{id}")

	// Сначала получим ID поста из списка
	posts, err := client.News.GetPosts(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	if len(posts.Items) == 0 {
		fmt.Println("   ✗ Нет постов для примера")
		return
	}

	postID := posts.Items[0].ID
	post, err := client.News.GetPost(ctx, postID)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Пост #%d загружен\n", postID)
	fmt.Printf("   - Автор: %s\n", post.AuthorName)
	fmt.Printf("   - Комментариев: %d (ВСЕ, а не превью)\n", len(post.Comments))
	for i, comment := range post.Comments {
		if i >= 2 {
			fmt.Printf("   ... и еще %d комментариев\n", len(post.Comments)-2)
			break
		}
		fmt.Printf("     * %s: %s\n", comment.AuthorName, truncate(comment.Text, 40))
	}
}

// ============================================================================
// PETITIONS SERVICE EXAMPLES
// ============================================================================

func exampleGetPetitions(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 15: Получить петиции")
	fmt.Println("   GET /api/petitions?status=active")

	petitions, err := client.Petitions.List(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Активных петиций: %d\n", len(petitions.Petitions))
	for i, p := range petitions.Petitions {
		if i >= 3 {
			fmt.Printf("   ... и еще %d\n", len(petitions.Petitions)-3)
			break
		}
		fmt.Printf("   - #%d: %s (автор: %s)\n", p.ID, p.Title, p.Author)
		fmt.Printf("     За: %d, Против: %d\n", p.VotesFor, p.VotesAgainst)
	}

	// Пример с фильтром по статусу
	fmt.Println("\n   Пример с фильтром (status=closed):")
	status := "closed"
	closedPetitions, err := client.Petitions.List(ctx, &lemicraft.PetitionsListOptions{
		Status: &status,
	})
	if err != nil {
		printError(err)
		return
	}
	fmt.Printf("   - Закрытых петиций: %d\n", len(closedPetitions.Petitions))
}

// ============================================================================
// COURT SERVICE EXAMPLES
// ============================================================================

func exampleGetCourtCases(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 16: Получить все судебные дела")
	fmt.Println("   GET /api/court")

	cases, err := client.Court.List(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Всего дел: %d\n", len(cases.Cases))
	for i, courtCase := range cases.Cases {
		if i >= 3 {
			fmt.Printf("   ... и еще %d\n", len(cases.Cases)-3)
			break
		}
		fmt.Printf("   - Дело #%d: %s\n", courtCase.ID, courtCase.Title)
		fmt.Printf("     Обвиняемый: %s, Статус: %s\n", courtCase.Defendant, courtCase.Status)
	}
}

func exampleGetCourtCase(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 17: Получить конкретное судебное дело")
	fmt.Println("   GET /api/court/{id}")

	// Сначала получим ID из списка
	cases, err := client.Court.List(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	if len(cases.Cases) == 0 {
		fmt.Println("   ✗ Нет дел для примера")
		return
	}

	caseID := cases.Cases[0].ID
	courtCase, err := client.Court.Get(ctx, caseID)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Дело #%d загружено\n", caseID)
	fmt.Printf("   - Название: %s\n", courtCase.Title)
	fmt.Printf("   - Описание: %s\n", truncate(courtCase.Description, 60))
	fmt.Printf("   - Обвиняемый: %s\n", courtCase.Defendant)
	fmt.Printf("   - Истец: %s\n", courtCase.Author)
	if courtCase.Judge != nil {
		fmt.Printf("   - Судья: %s\n", *courtCase.Judge)
	}
	fmt.Printf("   - Статус: %s\n", courtCase.Status)
}

func exampleGetCourtMessages(ctx context.Context, client *lemicraft.Client) {
	fmt.Println("\n📍 Пример 18: Получить сообщения (аргументы) в судебном деле")
	fmt.Println("   GET /api/court/messages/{id}")

	// Сначала получим ID из списка
	cases, err := client.Court.List(ctx, nil)
	if err != nil {
		printError(err)
		return
	}

	if len(cases.Cases) == 0 {
		fmt.Println("   ✗ Нет дел для примера")
		return
	}

	caseID := cases.Cases[0].ID
	messages, err := client.Court.GetMessages(ctx, caseID)
	if err != nil {
		printError(err)
		return
	}

	fmt.Printf("   ✓ Аргументов в деле #%d: %d\n", caseID, len(messages.Messages))
	for i, msg := range messages.Messages {
		if i >= 3 {
			fmt.Printf("   ... и еще %d\n", len(messages.Messages)-3)
			break
		}
		fmt.Printf("   - %s (%s):\n", msg.Author, msg.CreatedAt.Format("2006-01-02 15:04"))
		fmt.Printf("     %s\n", truncate(msg.Message, 60))
	}
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

func printError(err error) {
	var apiErr *lemicraft.APIError
	if errors.As(err, &apiErr) {
		fmt.Printf("   ✗ Ошибка API (%d): %s\n", apiErr.StatusCode, apiErr.Message)
		return
	}

	fmt.Printf("   ✗ Ошибка: %v\n", err)
}

func truncate(s string, length int) string {
	if len(s) > length {
		return s[:length] + "..."
	}
	return s
}

// String multiplication helper for Go (since Go doesn't have built-in)
func stringRepeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
