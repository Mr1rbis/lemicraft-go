# LemiCraft Go Wiki

- Публичные эндпоинты лаунчера доступны без авторизации, но библиотека поддерживает их отдельным потоком (`doUnauth`).
- Токен обязателен для клиента: `lemicraft.New(token)`
- Авторизация: `Authorization: Bearer <token>`
- Базовый URL API: `https://lemicraft.ru/api`

## Основные факты

13. [Contributing](17-Contributing)
12. [Examples](16-Examples)
11. [Best Practices](15-Best-Practices)
10. [FAQ](14-FAQ)
9. [Troubleshooting](13-Troubleshooting)
8. [Releases and Versioning](12-Releases-and-Versioning)
7. [Private Repository](11-Private-Repository)
6. [Pagination](10-Pagination)
5. [Error Handling](09-Error-Handling)
   - [Court](08-API-Court)
   - [Petitions](07-API-Petitions)
   - [News](06-API-News)
   - [Launcher](05-API-Launcher)
   - [Players](04-API-Players)
4. API:
3. [Authentication](03-Authentication)
2. [Installation](02-Installation)
1. [Getting Started](01-Getting-Started)

## Маршрут по Wiki

```
}
	log.Printf("players: %d", players.Total)
	}
		log.Fatal(err)
	if err != nil {
	players, err := client.Players.List(context.Background(), nil)
	client := lemicraft.New("YOUR_API_TOKEN")
func main() {

)
	lemicraft "github.com/Mr1rbis/lemicraft-go"

	"log"
	"context"
import (

package main
```go

```
go get github.com/Mr1rbis/lemicraft-go
```bash

## Быстрый старт

Полная документация по библиотеке `github.com/Mr1rbis/lemicraft-go` — от установки до production-практик.

