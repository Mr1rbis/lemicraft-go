# API: Players

Сервис: `client.Players`

## Методы

### `List(ctx, opts)`

- Endpoint: `GET /api/players`
- Query: `search` (optional)
- Возвращает: `*PlayersListResponse`

```go
search := "Steve"
res, err := client.Players.List(ctx, &lemicraft.ListOptions{Search: &search})
```

### `GetByNick(ctx, nick)`

- Endpoint: `GET /api/players/{nick}`
- Возвращает: `*PlayerFull`

```go
player, err := client.Players.GetByNick(ctx, "MrIrbis")
```

### `GetByDiscord(ctx, discordid)`

- Endpoint: `GET /api/users/discord/{discordid}`
- Возвращает: `*PlayerByDiscordId`
- Описание: Получает информацию об игроке по его Discord ID

```go
player, err := client.Players.GetByDiscord(ctx, "535868441433735188")
if err != nil {
	if errors.Is(err, lemicraft.ErrNotFound) {
		fmt.Println("Игрок не найден")
	}
	return
}

fmt.Printf("Discord: %s\n", player.DiscordUsername)
fmt.Printf("Minecraft: %s\n", player.MinecraftNick)
fmt.Printf("UUID: %s\n", player.MinecraftUuid)
fmt.Printf("В вайтлисте: %v\n", player.Whitelisted)
fmt.Printf("Источник ника: %s\n", player.NickSource)
```

**Структура ответа:**

```go
type PlayerByDiscordId struct {
    DiscordId       string // Discord ID пользователя
    DiscordUsername string // Имя в Discord
    MinecraftNick   string // Ник в Minecraft
    MinecraftUuid   string // UUID в Minecraft
    Whitelisted     bool   // Находится ли в вайтлисте
    NickSource      string // Источник получения ника (например, "multilogin")
}
```

### `GetPlan(ctx, nick)`

- Endpoint: `GET /api/plan/{nick}`
- Возвращает: `*PlayerPlan`
- `registered`/`lastSeen` поддерживают unix-ms и строковые даты через `APITime`.

```go
plan, err := client.Players.GetPlan(ctx, "MrIrbis")
if err == nil && plan.LastSeen != nil {
	fmt.Println(plan.LastSeen.Time)
}
```

### `GetAvatar(ctx, nick)`

- Endpoint: `GET /api/avatar/{nick}`
- Возвращает: `[]byte` (PNG)

```go
avatar, err := client.Players.GetAvatar(ctx, "MrIrbis")
```

### `GetSkin(ctx, nick)`

- Endpoint: `GET /api/skin/{nick}`
- Возвращает: `[]byte` (PNG)

```go
skin, err := client.Players.GetSkin(ctx, "MrIrbis")
```

## Важные типы

- `PlayerShort`
- `PlayerFull`
- `PlayerByDiscordId`
- `PlayerPlan`
- `UUIDValue` (гибкий парсинг `uuid`)

