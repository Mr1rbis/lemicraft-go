# Best Practices

## 1) Переиспользуйте один клиент

```go
client := lemicraft.New(token)
```

Не создавайте новый `Client` на каждый запрос.

## 2) Всегда используйте context timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

## 3) Проверяйте ошибки типами, а не строками

```go
if errors.Is(err, lemicraft.ErrNotFound) { /* ... */ }
```

```go
var apiErr *lemicraft.APIError
if errors.As(err, &apiErr) { /* ... */ }
```

## 4) Валидируйте входные параметры

- `nick` не пустой
- `id > 0`
- `limit` для launcher news в разумном диапазоне

## 5) Храните токен безопасно

- env vars
- secrets manager
- не логируйте токен

