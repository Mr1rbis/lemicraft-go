# Error Handling

Библиотека возвращает ошибки двух основных видов:

- `*lemicraft.APIError` — не-2xx ответ API
- `lemicraft.ErrNotFound` — специальный случай для 404

## Базовый шаблон

```go
var apiErr *lemicraft.APIError

if err != nil {
	if errors.Is(err, lemicraft.ErrNotFound) {
		// 404
		return
	}
	if errors.As(err, &apiErr) {
		switch apiErr.StatusCode {
		case 400:
		case 401:
		case 403:
		case 429:
		case 500, 502, 503, 504:
		}
		return
	}
	// сетевые/транспортные ошибки
}
```

## Что хранит APIError

- `StatusCode int`
- `Message string`

```go
if errors.As(err, &apiErr) {
	log.Printf("status=%d body=%s", apiErr.StatusCode, apiErr.Message)
}
```

