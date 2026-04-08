# Pagination

В библиотеке есть два cursor-потока пагинации.

## Новости (`/api/news`)

- Первая страница: `beforeID=nil`
- Следующая: передать `news.LastID` в `GetNews`

```go
news, err := client.News.GetNews(ctx, nil)
for err == nil && news.HasMore && news.LastID != nil {
	news, err = client.News.GetNews(ctx, news.LastID)
}
```

## Посты (`/api/posts`)

- Первая страница: `opts=nil`
- Следующая: `Before: posts.LastID`

```go
posts, err := client.News.GetPosts(ctx, nil)
for err == nil && posts.HasMore && posts.LastID != nil {
	posts, err = client.News.GetPosts(ctx, &lemicraft.PostsOptions{Before: posts.LastID})
}
```

## Советы

- Держите `context.WithTimeout` на каждый запрос.
- Сохраняйте последний cursor в хранилище, если нужен resume.

