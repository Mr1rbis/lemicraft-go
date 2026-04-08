# API: News

Сервис: `client.News`

## Методы

### `GetNews(ctx, beforeID)`

- Endpoint: `GET /api/news`
- Query: `before` (cursor)
- Возвращает: `*NewsResponse`

```go
news, err := client.News.GetNews(ctx, nil)
if err == nil && news.HasMore && news.LastID != nil {
	next, _ := client.News.GetNews(ctx, news.LastID)
	_ = next
}
```

### `GetGallery(ctx)`

- Endpoint: `GET /api/gallery`
- Возвращает: `*GalleryResponse`

```go
gallery, err := client.News.GetGallery(ctx)
```

### `GetPosts(ctx, opts)`

- Endpoint: `GET /api/posts`
- Query: `before` (int cursor)
- Возвращает: `*PostsResponse`

```go
posts, err := client.News.GetPosts(ctx, nil)
```

### `GetPost(ctx, id)`

- Endpoint: `GET /api/posts/{id}`
- Возвращает: `*Post`

```go
post, err := client.News.GetPost(ctx, 42)
```

