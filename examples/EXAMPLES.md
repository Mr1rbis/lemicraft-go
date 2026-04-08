# LemiCraft Go - Примеры для каждого эндпоинта

Этот файл содержит полные примеры использования всех 17 эндпоинтов библиотеки.

## Запуск примеров

### Требования
- Go 1.21+
- API токен Lemicraft (получить можно на https://lemicraft.ru/settings)

### Быстрый старт

```bash
cd /path/to/LemicraftGo
go run ./examples/complete/main.go -token="your-api-token"
```

## Примеры по сервисам

### 📋 PLAYERS SERVICE (6 эндпоинтов)

#### 1️⃣ Список всех игроков
```go
players, err := client.Players.List(ctx, nil)
// GET /api/players
// Возвращает: PlayersListResponse
```
**Результат:** Список всех игроков из вайтлиста с базовой информацией

#### 2️⃣ Поиск игроков
```go
search := "Steve"
players, err := client.Players.List(ctx, &lemicraft.ListOptions{
    Search: &search,
})
// GET /api/players?search=Steve
```
**Результат:** Игроки, совпадающие с поиском (регистронезависимо)

#### 3️⃣ Профиль игрока
```go
player, err := client.Players.GetByNick(ctx, "Notch")
// GET /api/players/{nick}
// Возвращает: PlayerFull
```
**Результат:** Полная информация об игроке (UUID, статус бана, ссылки на аватар/скин)

#### 4️⃣ Статистика игрока (Plan)
```go
plan, err := client.Players.GetPlan(ctx, "Notch")
// GET /api/plan/{nick}
// Возвращает: PlayerPlan
```
**Результат:** Статистика из плагина Plan (время игры, сессии, убийства и т.д.)

#### 5️⃣ Аватар игрока
```go
avatar, err := client.Players.GetAvatar(ctx, "Notch")
// GET /api/avatar/{nick}
// Возвращает: []byte (PNG 128x128)
```
**Результат:** Бинарные данные PNG изображения головы игрока

#### 6️⃣ Скин игрока
```go
skin, err := client.Players.GetSkin(ctx, "Notch")
// GET /api/skin/{nick}
// Возвращает: []byte (PNG 64x64)
```
**Результат:** Бинарные данные PNG текстуры скина

---

### 🚀 LAUNCHER SERVICE (3 эндпоинта) - Публичные!

Эти эндпоинты **не требуют авторизации**, но в библиотеке она используется везде.

#### 7️⃣ Версия лаунчера
```go
version, err := client.Launcher.GetVersion(ctx)
// GET /api/launcher/version
// Возвращает: LauncherVersion
```
**Результат:** Информация о последней версии лаунчера с GitHub

#### 8️⃣ Версия модпака
```go
modpack, err := client.Launcher.GetModpackVersion(ctx)
// GET /api/launcher/modpack/version
// Возвращает: ModpackVersion
```
**Результат:** Информация о последней версии модпака LemiSborka

#### 9️⃣ Новости для лаунчера
```go
limit := 20
category := "update"
news, err := client.Launcher.GetNews(ctx, &lemicraft.NewsOptions{
    Limit:    &limit,
    Category: &category,
})
// GET /api/launcher/news?limit=20&category=update
// Возвращает: LauncherNewsResponse
```
**Результат:** Список новостей в упрощенном формате для лаунчера

**Доступные категории:** general, update, event, maintenance, announcement

---

### 📰 NEWS SERVICE (4 эндпоинта)

#### 🔟 Новости сайта
```go
news, err := client.News.GetNews(ctx, nil)
// GET /api/news
// Возвращает: NewsResponse
// Поддерживает пагинацию: news.LastID → beforeID
```
**Результат:** Новости из Discord с HTML-рендером и пагинацией

#### 1️⃣1️⃣ Галерея
```go
gallery, err := client.News.GetGallery(ctx)
// GET /api/gallery
// Возвращает: GalleryResponse
```
**Результат:** Список одобренных скриншотов из галереи

#### 1️⃣2️⃣ Посты сообщества
```go
posts, err := client.News.GetPosts(ctx, nil)
// GET /api/posts
// Возвращает: PostsResponse (с превью из 3 комментариев)
// Поддерживает пагинацию: posts.LastID → Before
```
**Результат:** Список постов с лайками/дизлайками и превью комментариев

#### 1️⃣3️⃣ Один пост со всеми комментариями
```go
post, err := client.News.GetPost(ctx, 42)
// GET /api/posts/{id}
// Возвращает: Post (все комментарии, не превью!)
```
**Результат:** Полный пост со всеми (!) комментариями

---

### ✋ PETITIONS SERVICE (1 эндпоинт)

#### 1️⃣4️⃣ Петиции сервера
```go
// Активные (по умолчанию)
petitions, err := client.Petitions.List(ctx, nil)

// С фильтром по статусу
status := "closed"
petitions, err := client.Petitions.List(ctx, &lemicraft.PetitionsListOptions{
    Status: &status,
})
// GET /api/petitions?status=active
// Возвращает: PetitionsResponse
```
**Результат:** Список петиций с голосами

**Доступные статусы:** active (по умолчанию), closed, blocked

---

### ⚖️ COURT SERVICE (3 эндпоинта)

#### 1️⃣5️⃣ Все судебные дела
```go
cases, err := client.Court.List(ctx, nil)
// GET /api/court
// Возвращает: CourtCasesResponse
```
**Результат:** Список всех судебных дел

#### 1️⃣6️⃣ Конкретное судебное дело
```go
courtCase, err := client.Court.Get(ctx, 1)
// GET /api/court/{id}
// Возвращает: CourtCase
```
**Результат:** Полная информация об одном деле (название, участники, статус, вердикт)

#### 1️⃣7️⃣ Сообщения (аргументы) в деле
```go
messages, err := client.Court.GetMessages(ctx, 1)
// GET /api/court/messages/{id}
// Возвращает: CourtMessagesResponse
```
**Результат:** Все аргументы и возражения сторон в судебном деле

---

## Обработка ошибок

### Пример обработки в примерах:

```go
player, err := client.Players.GetByNick(ctx, "Notch")
if err != nil {
    // Проверка 404
    if errors.Is(err, lemicraft.ErrNotFound) {
        log.Println("Игрок не найден")
        return
    }
    
    // Проверка других ошибок API
    var apiErr *lemicraft.APIError
    if errors.As(err, &apiErr) {
        switch apiErr.StatusCode {
        case 401:
            log.Fatal("Неверный токен")
        case 429:
            log.Fatal("Слишком много запросов")
        default:
            log.Printf("Ошибка (%d): %s\n", apiErr.StatusCode, apiErr.Message)
        }
    }
}
```

---

## Пагинация

Некоторые эндпоинты поддерживают пагинацию:

```go
// Первая страница
posts, _ := client.News.GetPosts(ctx, nil)

// Получить следующую страницу
for posts.HasMore && posts.LastID != nil {
    posts, _ = client.News.GetPosts(ctx, &lemicraft.PostsOptions{
        Before: posts.LastID,
    })
}
```

---

## Производительность и рекомендации

### ✅ Правильно
```go
// Один клиент на все запросы
client := lemicraft.New("token")

players, _ := client.Players.List(ctx, nil)
posts, _ := client.News.GetPosts(ctx, nil)
cases, _ := client.Court.List(ctx, nil)

// С контекстом и таймаутом
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

### ❌ Неправильно
```go
// Новый клиент для каждого запроса
for i := 0; i < 100; i++ {
    client := lemicraft.New("token") // ❌ Плохо!
}

// Без таймаута
players, _ := client.Players.List(context.Background(), nil) // ❌ Может зависнуть
```

---

## Количество запросов

При выполнении всех примеров будет сделано примерно **20+ HTTP запросов**.

**Лимиты:**
- С авторизацией: 300 запросов/мин по ключу
- Один запрос в примере: < 100ms (обычно)

---

## Структура примера

```
├── exampleListPlayers()        // Пример 1
├── exampleSearchPlayers()      // Пример 2
├── exampleGetPlayerProfile()   // Пример 3
├── exampleGetPlayerStats()     // Пример 4
├── exampleDownloadAvatar()     // Пример 5
├── exampleDownloadSkin()       // Пример 6
├── exampleGetLauncherVersion() // Пример 7
├── exampleGetModpackVersion()  // Пример 8
├── exampleGetLauncherNews()    // Пример 9
├── exampleGetSiteNews()        // Пример 10
├── exampleGetGallery()         // Пример 11
├── exampleGetPosts()           // Пример 12
├── exampleGetSinglePost()      // Пример 13
├── exampleGetPetitions()       // Пример 14
├── exampleGetCourtCases()      // Пример 15
├── exampleGetCourtCase()       // Пример 16
└── exampleGetCourtMessages()   // Пример 17
```

---

## Запуск отдельных примеров

Если вам нужен только конкретный пример, отредактируйте `main()` и закомментируйте ненужные вызовы:

```go
func main() {
    client := lemicraft.New("token")
    ctx := context.Background()
    
    // Только этот пример
    exampleListPlayers(ctx, client)
}
```

---

Для полной документации смотрите [README.md](../README.md)

