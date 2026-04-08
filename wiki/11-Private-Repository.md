# Private Repository

Если репозиторий приватный, настройте доступ Go к private modules.

## Вариант 1: SSH (рекомендуется)

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

```bash
export GOPRIVATE=github.com/Mr1rbis/*
```

```bash
go get github.com/Mr1rbis/lemicraft-go@latest
```

## Вариант 2: Personal Access Token

```bash
git config --global url."https://<TOKEN>@github.com/".insteadOf "https://github.com/"
```

```bash
export GOPRIVATE=github.com/Mr1rbis/*
```

```bash
go get github.com/Mr1rbis/lemicraft-go@latest
```

## Локально без GitHub

Используйте `replace` в `go.mod`:

```go
replace github.com/Mr1rbis/lemicraft-go => /home/mrirbis/GolandProjects/LemicraftGo
```

