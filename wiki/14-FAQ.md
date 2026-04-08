# FAQ

## Можно ли использовать библиотеку с приватным GitHub-репозиторием?

Да. Настройте `GOPRIVATE` и Git auth. См. [Private Repository](11-Private-Repository).

## Можно ли без токена?

Конструктор `New` требует токен. Для лаунчер-эндпоинтов API допускает public доступ, но библиотека клиент создается с токеном.

## Как обрабатывать 404?

```go
if errors.Is(err, lemicraft.ErrNotFound) {
	// not found
}
```

## Как лучше пагинировать посты?

Используйте `PostsOptions.Before` и цикл по `HasMore` + `LastID`.

## Можно ли кастомизировать HTTP client?

Да, через `WithHTTPClient`.

## Можно ли изменить base URL?

Да, через `WithBaseURL`.

