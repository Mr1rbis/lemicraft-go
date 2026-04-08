# Getting Started

## Что это за библиотека

`lemicraft-go` — typed Go SDK для LemiCraft API с сервисами:

- `Client.Players`
- `Client.Launcher`
- `Client.News`
- `Client.Petitions`
- `Client.Court`

## Требования

- Go `1.21+`
- API token с `https://lemicraft.ru/settings`

## Минимальный пример

```go
client := lemicraft.New("YOUR_API_TOKEN")
ctx := context.Background()

p, err := client.Players.GetByNick(ctx, "MrIrbis")
if err != nil {
	// обработка
}
fmt.Println(p.Name)
```

## Ключевые идеи

- Клиент создается один раз и переиспользуется.
- Почти все эндпоинты требуют токен.
- Ошибки API возвращаются как `*lemicraft.APIError`.
- Для 404 работает `errors.Is(err, lemicraft.ErrNotFound)`.

