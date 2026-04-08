# API: Court

Сервис: `client.Court`

## Методы

### `List(ctx, opts)`

- Endpoint: `GET /api/court`
- Query: `status` (`open|closed`)
- Возвращает: `*CourtCasesResponse`

```go
cases, err := client.Court.List(ctx, nil)
```

### `Get(ctx, id)`

- Endpoint: `GET /api/court/{id}`
- Возвращает: `*CourtCase`

```go
courtCase, err := client.Court.Get(ctx, 1)
```

### `GetMessages(ctx, caseID)`

- Endpoint: `GET /api/court/messages/{id}`
- Возвращает: `*CourtMessagesResponse`

```go
msgs, err := client.Court.GetMessages(ctx, 1)
```

