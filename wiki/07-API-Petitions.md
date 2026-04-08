# API: Petitions

Сервис: `client.Petitions`

## Метод `List(ctx, opts)`

- Endpoint: `GET /api/petitions`
- Query: `status` (`active|closed|blocked`)
- Возвращает: `*PetitionsResponse`

```go
res, err := client.Petitions.List(ctx, nil)

status := "closed"
closedRes, err := client.Petitions.List(ctx, &lemicraft.PetitionsListOptions{Status: &status})
_ = closedRes
```

