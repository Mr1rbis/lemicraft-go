# LemiCraft Go Library
![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
Официальная Go-библиотека для работы с [LemiCraft API](https://lemicraft.ru/api-reference).
## 📦 Установка
```bash
go get github.com/Mr1rbis/lemicraft-go
```
Требуется **Go 1.21+**.
## 🚀 Быстрый старт
```go
package main
import (
    "context"
    "log"
    lemicraft "github.com/Mr1rbis/lemicraft-go"
)
func main() {
    // Получите API токен на https://lemicraft.ru/settings
    client := lemicraft.New("your-api-token")
    ctx := context.Background()
    // Получить список игроков
    players, err := client.Players.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Total players: %d\n", players.Total)
}
```
## 🔐 Авторизация
API токен **обязателен** для использования библиотеки.
Получить токен можно в [личном кабинете](https://lemicraft.ru/settings).
```go
// Создание клиента с токеном
client := lemicraft.New("your-api-token")
// Опционально: кастомизация базового URL (для тестирования)
client := lemicraft.New("token", 
    lemicraft.WithBaseURL("https://staging.lemicraft.ru/api"),
)
// Опционально: кастомный HTTP клиент с таймаутом
import "net/http"
import "time"
customHTTPClient := &http.Client{
    Timeout: 15 * time.Second,
}
client := lemicraft.New("token",
    lemicraft.WithHTTPClient(customHTTPClient),
)
```
## 📚 API Эндпоинты
### Игроки (Players)
#### Получить список игроков
```go
// Получить всех игроков
players, err := client.Players.List(ctx, nil)
// С поиском по нику
search := "Steve"
players, err := client.Players.List(ctx, &lemicraft.ListOptions{
    Search: &search,
})
// Использование
for _, p := range players.Players {
    log.Printf("%s (UUID: %v, Banned: %v)\n", p.Name, p.UUID, p.Banned)
}
```
#### Получить профиль игрока
```go
player, err := client.Players.GetByNick(ctx, "Notch")
if err != nil {
    log.Fatal(err)
}
log.Printf("Ник: %s\n", player.Name)
log.Printf("Забанен: %v\n", player.Banned)
log.Printf("Аватар: %s\n", player.AvatarURL)
log.Printf("Скин: %s\n", player.SkinURL)
```
#### Получить статистику игрока (Plan)
```go
plan, err := client.Players.GetPlan(ctx, "Notch")
if err != nil {
    log.Fatal(err)
}
log.Printf("Онлайн: %v\n", plan.Online)
log.Printf("Время игры: %s\n", plan.PlaytimeStr)
log.Printf("Сессии: %d\n", *plan.Sessions)
log.Printf("Убийства мобов: %d\n", *plan.MobKills)
```
#### Скачать аватар и скин
```go
// Скачать аватар (128x128 PNG)
avatar, err := client.Players.GetAvatar(ctx, "Notch")
if err != nil {
    log.Fatal(err)
}
ioutil.WriteFile("avatar.png", avatar, 0644)
// Скачать скин (64x64 PNG)
skin, err := client.Players.GetSkin(ctx, "Notch")
if err != nil {
    log.Fatal(err)
}
ioutil.WriteFile("skin.png", skin, 0644)
```
### Лаунчер (Launcher) - Публичные эндпоинты
Эти эндпоинты **не требуют** авторизации.
#### Версия лаунчера
```go
version, err := client.Launcher.GetVersion(ctx)
if err != nil {
    log.Fatal(err)
}
log.Printf("Последняя версия: %s\n", version.Version)
log.Printf("Скачать: %s\n", version.DownloadURL)
log.Printf("Размер: %d MB\n", version.FileSize / 1024 / 1024)
log.Printf("Дата: %s\n", version.ReleaseDate.Format("2006-01-02"))
```
#### Версия модпака
```go
modpack, err := client.Launcher.GetModpackVersion(ctx)
if err != nil {
    log.Fatal(err)
}
log.Printf("Модпак: %s\n", modpack.Name)
log.Printf("Версия: %s\n", modpack.Version)
log.Printf("Скачать: %s\n", modpack.DownloadURL)
```
#### Новости для лаунчера
```go
// Получить последние новости
news, err := client.Launcher.GetNews(ctx, nil)
// С параметрами
limit := 20
category := "update"
news, err := client.Launcher.GetNews(ctx, &lemicraft.NewsOptions{
    Limit:    &limit,
    Category: &category,
})
for _, item := range news.Items {
    log.Printf("[%s] %s - %s\n", item.Category, item.Title, item.AuthorName)
}
```
### Контент (News/Content)
#### Новости сайта
```go
// Получить первую страницу новостей
news, err := client.News.GetNews(ctx, nil)
if err != nil {
    log.Fatal(err)
}
// Пагинация
for _, item := range news.Items {
    log.Printf("HTML: %s\n", item.HTML)
    log.Printf("Автор: %s\n", item.AuthorName)
}
// Получить следующую страницу
if news.HasMore && news.LastID != nil {
    nextNews, err := client.News.GetNews(ctx, news.LastID)
}
```
#### Галерея
```go
gallery, err := client.News.GetGallery(ctx)
if err != nil {
    log.Fatal(err)
}
for _, img := range gallery.Images {
    log.Printf("ID: %d, Файл: %s, URL: %s\n", img.ID, img.Filename, img.URL)
}
```
#### Посты сообщества
```go
// Получить список постов (с первыми 3 комментариями в каждом)
posts, err := client.News.GetPosts(ctx, nil)
if err != nil {
    log.Fatal(err)
}
// Пагинация
for posts.HasMore && posts.LastID != nil {
    posts, err = client.News.GetPosts(ctx, &lemicraft.PostsOptions{
        Before: posts.LastID,
    })
    if err != nil {
        break
    }
}
// Получить один пост со ВСЕМИ комментариями
post, err := client.News.GetPost(ctx, 42)
if err != nil {
    log.Fatal(err)
}
log.Printf("Автор: %s\n", post.AuthorName)
log.Printf("Текст: %s\n", post.Text)
log.Printf("Лайков: %d, Дизлайков: %d\n", post.VotesFor, post.VotesAgainst)
for _, comment := range post.Comments {
    log.Printf("  %s: %s\n", comment.AuthorName, comment.Text)
}
```
### Сообщество (Community)
#### Петиции
```go
// Получить активные петиции (по умолчанию)
petitions, err := client.Petitions.List(ctx, nil)
// Фильтр по статусу
status := "closed"
petitions, err := client.Petitions.List(ctx, &lemicraft.PetitionsListOptions{
    Status: &status,
})
for _, p := range petitions.Petitions {
    log.Printf("Заголовок: %s\n", p.Title)
    log.Printf("Автор: %s\n", p.Author)
    log.Printf("За: %d, Против: %d\n", p.VotesFor, p.VotesAgainst)
    log.Printf("Статус: %s\n", p.Status)
}
```
#### Судебные дела
```go
// Получить все открытые дела
cases, err := client.Court.List(ctx, nil)
log.Printf("Дел: %d\n", len(cases.Cases))
// Получить конкретное дело
courtCase, err := client.Court.Get(ctx, 1)
if err != nil {
    log.Fatal(err)
}
log.Printf("Название: %s\n", courtCase.Title)
log.Printf("Обвиняемый: %s\n", courtCase.Defendant)
log.Printf("Истец: %s\n", courtCase.Author)
log.Printf("Статус: %s\n", courtCase.Status)
// Получить все аргументы сторон в деле
messages, err := client.Court.GetMessages(ctx, 1)
if err != nil {
    log.Fatal(err)
}
for _, msg := range messages.Messages {
    log.Printf("%s (%s): %s\n", msg.Author, msg.CreatedAt.Format("2006-01-02"), msg.Message)
}
```
## ⚠️ Обработка ошибок
### Типичные ошибки API
```go
import "errors"
user, err := client.Players.GetByNick(ctx, "NonExistentPlayer")
if err != nil {
    // Проверка конкретной ошибки "не найдено"
    if errors.Is(err, lemicraft.ErrNotFound) {
        log.Println("Игрок не найден (404)")
        return
    }
    // Проверка любой ошибки API
    var apiErr *lemicraft.APIError
    if errors.As(err, &apiErr) {
        switch apiErr.StatusCode {
        case 400:
            log.Printf("Неправильный запрос: %s\n", apiErr.Message)
        case 401:
            log.Fatal("Не авторизован - проверьте токен")
        case 403:
            log.Fatal("Доступ запрещён")
        case 429:
            log.Fatal("Слишком много запросов - попробуйте позже")
        case 500, 502, 503, 504:
            log.Printf("Ошибка сервера (%d): %s\n", apiErr.StatusCode, apiErr.Message)
        default:
            log.Printf("Ошибка API (%d): %s\n", apiErr.StatusCode, apiErr.Message)
        }
        return
    }
    // Сетевые ошибки, таймауты и прочее
    log.Printf("Ошибка запроса: %v\n", err)
}
```
## 📊 Структура типов данных
### Players
```go
type PlayerShort struct {
    Name      string  // Ник игрока
    UUID      *string // UUID или nil
    ServiceID int     // 0 = Mojang, 1 = Ely.by
    Banned    bool    // Забанен ли
}
type PlayerFull struct {
    Name        string    // Ник
    UUID        *string   // UUID
    ServiceID   int       // Сервис авторизации
    Banned      bool      // Забанен
    BanReason   *string   // Причина бана
    BanPermanent bool    // Перманентный ли бан
    BanEndsAt   *time.Time // Когда закончится бан
    AvatarURL   string    // URL аватара
    SkinURL     string    // URL скина
}
type PlayerPlan struct {
    Registered     *time.Time // Первый вход
    LastSeen       *time.Time // Последний визит
    Online         bool       // Онлайн ли сейчас
    Playtime       *int64     // Общее время в мс
    ActivePlaytime *int64     // Активное время в мс (без AFK)
    PlaytimeStr    string     // Время в читаемом формате
    Sessions       *int       // Количество сессий
    LongestSession *int64     // Самая долгая сессия в мс
    Deaths         *int       // Смертей
    MobKills       *int       // Убийств мобов
    PlayerKills    *int       // Убийств игроков
    Ping           *float64   // Средний пинг
}
```
### Posts
```go
type Post struct {
    ID              int
    AuthorDiscordID string
    AuthorName      string
    Text            string
    Edited          bool
    CreatedAt       time.Time
    Media           []string  // URLs к изображениям
    VotesFor        int       // Лайки
    VotesAgainst    int       // Дизлайки
    UserVote        *int      // Ваш голос (1, -1, или nil)
    Comments        []Comment
    CommentCount    int
}
```
## 🎯 Best Practices
### 1. Используйте контекст с таймаутом
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
players, err := client.Players.List(ctx, nil)
```
### 2. Переиспользуйте клиент
```go
// ✅ Правильно - один клиент для всех запросов
client := lemicraft.New("token")
players, _ := client.Players.List(ctx, nil)
news, _ := client.News.GetNews(ctx, nil)
cases, _ := client.Court.List(ctx, nil)
// ❌ Неправильно - не создавайте новый клиент каждый раз
for i := 0; i < 100; i++ {
    client := lemicraft.New("token") // ❌ Плохо!
}
```
### 3. Проверяйте ошибки правильно
```go
// ✅ Правильно
if errors.Is(err, lemicraft.ErrNotFound) {
    // ...
}
var apiErr *lemicraft.APIError
if errors.As(err, &apiErr) {
    // ...
}
// ❌ Неправильно - не сравнивайте строки
if err != nil && strings.Contains(err.Error(), "404") {
    // Плохой способ проверки
}
```
### 4. Пагинация
```go
// Получить первую страницу
posts, err := client.News.GetPosts(ctx, nil)
if err != nil {
    log.Fatal(err)
}
// Пока есть еще страницы
for posts.HasMore && posts.LastID != nil {
    posts, err = client.News.GetPosts(ctx, &lemicraft.PostsOptions{
        Before: posts.LastID,
    })
    if err != nil {
        break
    }
}
```
## 📊 Лимиты и квоты
- **Без авторизации** (только публичные эндпоинты типа `/launcher/*`): 60 запросов/мин по IP
- **С авторизацией**: 300 запросов/мин по API ключу
## 🤝 Вклад
Приветствуются пул-реквесты и issue репорты на GitHub.
## 📄 Лицензия
MIT License - см. LICENSE файл
## 🔗 Ссылки
- [LemiCraft Website](https://lemicraft.ru)
- [API Reference](https://lemicraft.ru/api-reference)
- [Settings (API Keys)](https://lemicraft.ru/settings)
---
**Последнее обновление:** 2025-01-15  
**Версия API:** 1.0
