# LemicraftGo

Go-клиент для [Lemicraft API](https://lemicraft.ru). Библиотека спроектирована с чётким разделением ответственности, поддержкой авторизации одним вызовом и лёгким добавлением новых эндпоинтов.

## Содержание

- [Установка](#установка)
- [Быстрый старт](#быстрый-старт)
- [Конфигурация клиента](#конфигурация-клиента)
- [Авторизация](#авторизация)
- [Обработка ошибок](#обработка-ошибок)
- [API Reference](#api-reference)
- [Структура проекта](#структура-проекта)
- [Расширение библиотеки](#расширение-библиотеки)

---

## Установка

```bash
go get github.com/Mr1rbis/lemicraft-go
```

Требуется **Go 1.21+**.

---

## Быстрый старт

```go
package main

import (
    "context"
    "fmt"
    "log"

    lemicraft "github.com/Mr1rbis/lemicraft-go"
)

func main() {
    client := lemicraft.New()

    ctx := context.Background()

    user, err := client.Users.GetByDiscordID(ctx, "535868441433735188")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(user.MinecraftNick) // DaPoT22
}
```

---

## Конфигурация клиента

Клиент создаётся через `lemicraft.New()` с функциональными опциями:

```go
import "net/http"

client := lemicraft.New(
    // Кастомный базовый URL (например, staging-окружение)
    lemicraft.WithBaseURL("https://staging.lemicraft.ru/api"),

    // Свой http.Client — для настройки таймаутов, TLS и т.д.
    lemicraft.WithHTTPClient(&http.Client{
        Timeout: 10 * time.Second,
    }),

    // Bearer-токен авторизации (см. раздел ниже)
    lemicraft.WithAuthToken("your-token-here"),
)
```

Опции можно комбинировать в любом порядке. Без опций клиент использует `https://lemicraft.ru/api` и стандартный `http.Client`.

---

## Авторизация

Авторизация подключается **одной опцией** и автоматически применяется ко **всем** текущим и будущим эндпоинтам:

```go
client := lemicraft.New(
    lemicraft.WithAuthToken("your-api-token"),
)
```

Токен передаётся через HTTP-заголовок `Authorization: Bearer <token>` на уровне транспорта (`http.RoundTripper`), что гарантирует его присутствие в каждом запросе без изменения кода сервисов.

---

## Обработка ошибок

Библиотека возвращает два вида ошибок:

| Тип | Когда возникает | Как проверить |
|-----|----------------|---------------|
| `*lemicraft.APIError` | Сервер вернул не-2xx статус | `errors.As(err, &apiErr)` |
| `lemicraft.ErrNotFound` | Сервер вернул 404 | `errors.Is(err, lemicraft.ErrNotFound)` |
| `error` (стандартный) | Сетевая ошибка, таймаут | `err != nil` |

### Пример полной обработки

```go
user, err := client.Users.GetByDiscordID(ctx, discordID)
if err != nil {
    // 404 — пользователь не найден
    if errors.Is(err, lemicraft.ErrNotFound) {
        fmt.Println("пользователь не найден")
        return
    }

    // Любая другая HTTP-ошибка — достаём код и тело ответа
    var apiErr *lemicraft.APIError
    if errors.As(err, &apiErr) {
        switch {
        case apiErr.StatusCode == 401:
            fmt.Println("не авторизован — передайте токен через WithAuthToken")
        case apiErr.StatusCode == 403:
            fmt.Println("доступ запрещён")
        case apiErr.StatusCode >= 500:
            fmt.Printf("ошибка сервера (%d): %s\n", apiErr.StatusCode, apiErr.Message)
        default:
            fmt.Printf("API error (%d): %s\n", apiErr.StatusCode, apiErr.Message)
        }
        return
    }

    // Сетевая ошибка / таймаут
    fmt.Printf("ошибка запроса: %v\n", err)
    return
}
```

> **Примечание:** `errors.Is(err, lemicraft.ErrNotFound)` работает автоматически благодаря методу `(*APIError).Is()` — отдельно проверять `StatusCode == 404` не нужно.

---

## API Reference

### `lemicraft.New(opts ...Option) *Client`

Создаёт новый клиент. Принимает опциональные настройки.

---

### `client.Users` — `*UsersService`

#### `GetByDiscordID(ctx context.Context, discordID string) (*UserByDiscord, error)`

Возвращает пользователя Lemicraft по его Discord ID.

```
GET https://lemicraft.ru/api/users/discord/<discordID>
```

**Параметры:**

| Параметр | Тип | Описание |
|----------|-----|----------|
| `ctx` | `context.Context` | Контекст для отмены/таймаута запроса |
| `discordID` | `string` | Discord snowflake ID пользователя |

**Возвращает:** `*UserByDiscord`, `error`

---

### Типы

#### `UserByDiscord`

```go
type UserByDiscord struct {
    DiscordID       string // Discord snowflake ID
    DiscordUsername string // Имя пользователя в Discord
    MinecraftNick   string // Ник в Minecraft
    MinecraftUUID   string // UUID Minecraft-аккаунта
    Whitelisted     bool   // Находится ли в вайтлисте сервера
    NickSource      string // Источник ника (например, "multilogin")
}
```

#### `APIError`

```go
type APIError struct {
    StatusCode int    // HTTP-статус код
    Message    string // Тело ответа от сервера
}
```

#### `ErrNotFound`

```go
var ErrNotFound = errors.New("lemicraft: not found")
```

Sentinel-ошибка для HTTP 404. Проверяется через `errors.Is`.

---

## Структура проекта

```
github.com/Mr1rbis/lemicraft-go
├── go.mod
├── README.md
├── client.go            # Client, New(), центральный метод do()
├── options.go           # Функциональные опции (WithAuthToken, WithBaseURL, WithHTTPClient)
├── transport.go         # authTransport — инъекция авторизации на уровне RoundTripper
├── types.go             # Типы данных и ошибок (UserByDiscord, APIError, ErrNotFound)
├── users.go             # UsersService — эндпоинты /api/users/*
└── examples/
    └── basic/
        └── main.go      # Пример использования (go run ./examples/basic/)
```

---

## Расширение библиотеки

### Добавить новый эндпоинт в существующий ресурс

Достаточно добавить метод в соответствующий сервис. Например, `GetByMinecraftUUID` в `users.go`:

```go
func (s *UsersService) GetByMinecraftUUID(ctx context.Context, uuid string) (*UserByMinecraft, error) {
    url := fmt.Sprintf("%s/users/minecraft/%s", s.client.baseURL, uuid)
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        return nil, fmt.Errorf("lemicraft: building request: %w", err)
    }
    req.Header.Set("Accept", "application/json")
    var user UserByMinecraft
    if err := s.client.do(req, &user); err != nil {
        return nil, err
    }
    return &user, nil
}
```

### Добавить новый ресурс (например, серверы)

**1.** Создать файл `servers.go` в корне модуля:

```go
package lemicraft

import (
    "context"
    "fmt"
    "net/http"
)

type ServersService struct {
    client *Client
}

func (s *ServersService) List(ctx context.Context) ([]Server, error) {
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.client.baseURL+"/servers", nil)
    if err != nil {
        return nil, fmt.Errorf("lemicraft: building request: %w", err)
    }
    req.Header.Set("Accept", "application/json")
    var servers []Server
    if err := s.client.do(req, &servers); err != nil {
        return nil, err
    }
    return servers, nil
}
```

**2.** Добавить поле в `Client` и инициализировать в `New()` в `client.go`:

```go
type Client struct {
    // ...
    Users   *UsersService
    Servers *ServersService  // новый сервис
}

func New(opts ...Option) *Client {
    // ...
    c.Users   = &UsersService{client: c}
    c.Servers = &ServersService{client: c}  // инициализация
    return c
}
```

Авторизация, таймауты и базовый URL применяются автоматически — менять ничего не нужно.

