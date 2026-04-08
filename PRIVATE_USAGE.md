# Использование приватной LemiCraft Go библиотеки

## Быстрый старт

### Шаг 1: Получить API токен

1. Зайдите на https://lemicraft.ru/settings
2. Сгенерируйте API ключ
3. Скопируйте его

### Шаг 2: Использовать в проекте

#### Вариант A: Локальная разработка

```bash
# В вашем проекте
go mod init my-lemicraft-app

# Добавьте в go.mod:
replace github.com/Mr1rbis/lemicraft-go => /path/to/local/lemicraft-go

# Затем импортируйте
go get github.com/Mr1rbis/lemicraft-go
```

#### Вариант B: GitHub (с Personal Access Token)

```bash
# 1. Создайте PAT на https://github.com/settings/tokens/new
#    Нужны scopes: "repo" и "read:packages"

# 2. Установите переменные окружения (в ~/.bashrc или ~/.zshrc)
export GOPRIVATE="github.com/Mr1rbis/*"
export GITHUB_TOKEN="ghp_xxxxx..."

# 3. Настройте Git
git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

# 4. Теперь можно использовать
go get github.com/Mr1rbis/lemicraft-go@latest
```

#### Вариант C: SSH ключи (рекомендуется)

```bash
# 1. Проверьте SSH ключ
cat ~/.ssh/id_ed25519.pub

# 2. Добавьте его на https://github.com/settings/keys

# 3. Установите GOPRIVATE
export GOPRIVATE="github.com/Mr1rbis/*"

# 4. Используйте SSH вместо HTTPS
git config --global url."git@github.com:".insteadOf "https://github.com/"

# 5. Теперь работает:
go get github.com/Mr1rbis/lemicraft-go@latest
```

### Шаг 3: Использовать в коде

```go
package main

import (
    "context"
    "log"
    "time"
    
    lemicraft "github.com/Mr1rbis/lemicraft-go"
)

func main() {
    // Создайте клиент с токеном
    client := lemicraft.New("ваш-api-токен")
    
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Используйте API
    players, err := client.Players.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Players: %d\n", players.Total)
}
```

## Решение проблем

### Ошибка: "module github.com/Mr1rbis/lemicraft-go@latest: invalid version"

**Решение:**
```bash
# Убедитесь что GOPRIVATE установлен
export GOPRIVATE="github.com/Mr1rbis/*"

# Убедитесь что Git настроен с вашим токеном/SSH
git config --global url."git@github.com:".insteadOf "https://github.com/"

# Очистите кэш Go
go clean -modcache

# Попробуйте снова
go get github.com/Mr1rbis/lemicraft-go@latest
```

### Ошибка: "terminal prompts disabled"

**Решение:** GitHub требует аутентификацию. Используйте один из вариантов выше (SSH или Personal Access Token).

### Ошибка: "could not read Username for 'https://github.com'"

**Решение:** Настройте Git:
```bash
# SSH (рекомендуется)
git config --global url."git@github.com:".insteadOf "https://github.com/"

# Или с токеном
git config --global url."https://TOKEN@github.com/".insteadOf "https://github.com/"
```

## Docker / CI/CD

### GitHub Actions

```yaml
name: Test

on: [push]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Configure Go private packages
        run: |
          git config --global url."https://${{ secrets.GITHUB_TOKEN }}@github.com/".insteadOf "https://github.com/"
        env:
          GOPRIVATE: github.com/Mr1rbis/*
      
      - name: Download dependencies
        run: go mod download
      
      - name: Test
        run: go test -v ./...
```

### Docker

```dockerfile
FROM golang:1.21

# Копируйте SSH ключ (опционально)
COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

RUN git config --global url."git@github.com:".insteadOf "https://github.com/" && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN go build -o app
```

## Локальное тестирование

```bash
# Клонируйте обе директории рядом
/path/to/projects/
├── lemicraft-go/          # Клонированная библиотека
└── my-app/                # Ваше приложение

# В my-app/go.mod добавьте:
replace github.com/Mr1rbis/lemicraft-go => ../lemicraft-go

# Теперь при разработке библиотеки изменения видны сразу в app
```

## Создание новой версии библиотеки

Если вы разработчик lemicraft-go:

```bash
cd lemicraft-go

# Создайте тэг
git tag -a v0.2.0 -m "Release version 0.2.0"

# Запушьте тэг
git push origin v0.2.0

# Go автоматически создаст версию
# Другие проекты смогут использовать:
go get github.com/Mr1rbis/lemicraft-go@v0.2.0
```

## FAQ

**Q: Могу ли я использовать эту библиотеку в своем приватном проекте?**  
A: Да! Используйте любой из вариантов выше (локальный, SSH, или с токеном).

**Q: Можно ли поделиться библиотекой с другим разработчиком?**  
A: Да, но они должны иметь доступ к приватному репозиторию на GitHub.

**Q: Работает ли с Go < 1.21?**  
A: Нет, требуется Go 1.21+.

**Q: Как обновить до новой версии?**  
A: `go get -u github.com/Mr1rbis/lemicraft-go`

**Q: Могу ли я использовать разные версии?**  
A: Да: `go get github.com/Mr1rbis/lemicraft-go@v0.1.0`

---

Для полной документации смотрите [README.md](README.md) и [DEVELOPMENT.md](DEVELOPMENT.md).

