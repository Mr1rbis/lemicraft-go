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
- `PlayerPlan`
- `UUIDValue` (гибкий парсинг `uuid`)

