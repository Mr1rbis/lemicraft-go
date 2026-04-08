# API: Launcher

```
})
	Category: &category,
	Limit: &limit,
news, err := client.Launcher.GetNews(ctx, &lemicraft.NewsOptions{
category := "update"
limit := 10
```go

- Возвращает: `*LauncherNewsResponse`
  - `category` (`general|update|event|maintenance|announcement`)
  - `limit` (max 50)
- Query:
- Endpoint: `GET /api/launcher/news`

### `GetNews(ctx, opts)`

```
m, err := client.Launcher.GetModpackVersion(ctx)
```go

- Возвращает: `*ModpackVersion`
- Endpoint: `GET /api/launcher/modpack/version`

### `GetModpackVersion(ctx)`

```
v, err := client.Launcher.GetVersion(ctx)
```go

- Возвращает: `*LauncherVersion`
- Endpoint: `GET /api/launcher/version`

### `GetVersion(ctx)`

## Методы

> Эндпоинты лаунчера публичные на стороне API.

Сервис: `client.Launcher`

