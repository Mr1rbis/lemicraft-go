# Examples

```
}
	posts, err = client.News.GetPosts(ctx, &lemicraft.PostsOptions{Before: posts.LastID})
for err == nil && posts.HasMore && posts.LastID != nil {
posts, err := client.News.GetPosts(ctx, nil)
```go

## Пример paginated posts

```
return os.WriteFile("avatar.png", avatar, 0o644)
}
	return err
if err != nil {
avatar, err := client.Players.GetAvatar(ctx, "MrIrbis")
```go

## Пример сохранения аватара

```
go run ./examples/complete/main.go -token="YOUR_API_TOKEN"
```bash

## Запуск полного примера

```
go run ./examples/basic/main.go -token="YOUR_API_TOKEN"
```bash

## Запуск базового примера

- `examples/EXAMPLES.md`
- `examples/complete/main.go`
- `examples/basic/main.go`

## Официальные примеры в репозитории

