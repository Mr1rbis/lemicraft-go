# Authentication

## Обязательно ли передавать токен

Да. Конструктор клиента требует токен:

```go
client := lemicraft.New("YOUR_API_TOKEN")
```

Если передать пустую строку, будет `panic`.

## Где получить токен

- `https://lemicraft.ru/settings`

## Как передается токен

Библиотека автоматически ставит заголовок:

```http
Authorization: Bearer <token>
```

## Публичные эндпоинты

`/launcher/*` могут работать без токена на стороне API. В библиотеке для них используется отдельный внутренний поток запроса (`doUnauth`).

## Настройка клиента

```go
custom := &http.Client{Timeout: 15 * time.Second}
client := lemicraft.New(
	"YOUR_API_TOKEN",
	lemicraft.WithBaseURL("https://lemicraft.ru/api"),
	lemicraft.WithHTTPClient(custom),
)
```

