# Разработка и использование библиотеки локально

## Локальное тестирование

### Вариант 1: Прямой импорт с локального диска

Если вы разрабатываете приложение на том же компьютере:

```go
import (
    lemicraft "github.com/Mr1rbis/lemicraft-go"
)
```

И в `go.mod` добавьте:

```
require github.com/Mr1rbis/lemicraft-go v0.1.0

replace github.com/Mr1rbis/lemicraft-go => /home/mrirbis/GolandProjects/LemicraftGo
```

Тогда Go будет использовать локальную версию.

### Вариант 2: go work (Go 1.18+)

```bash
# В директории вашего проекта
go work init
go work use /home/mrirbis/GolandProjects/LemicraftGo
go work use .
```

## Использование приватной библиотеки

### Проблема: GitHub требует аутентификации для приватных репозиториев

Если ваш репозиторий на GitHub частный, Go не может скачать его без аутентификации.

**Решение 1: GitHub Personal Access Token (Рекомендуется)**

1. Создайте Personal Access Token на https://github.com/settings/tokens
   - Выберите scopes: `repo`, `read:packages`

2. Настройте Git:

```bash
git config --global url."https://<YOUR_TOKEN>@github.com/".insteadOf "https://github.com/"
```

3. Настройте Go:

```bash
# В ~/.bashrc или ~/.zshrc
export GOPRIVATE=github.com/Mr1rbis/*
```

4. Теперь `go get` будет работать:

```bash
go get github.com/Mr1rbis/lemicraft-go@v0.1.0
```

**Решение 2: SSH ключи**

1. Сгенерируйте SSH ключ (если его нет):

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

2. Добавьте публичный ключ на GitHub в SSH keys

3. Настройте Git использовать SSH:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

4. Настройте Go:

```bash
export GOPRIVATE=github.com/Mr1rbis/*
```

5. Используйте SSH при импорте:

```bash
go get git@github.com:Mr1rbis/lemicraft-go@v0.1.0
```

**Решение 3: Локальная директория (для разработки)**

При разработке просто используйте `replace` в `go.mod`:

```
replace github.com/Mr1rbis/lemicraft-go => /path/to/local/repo
```

## Создание тэгов версий

Для публикации нового релиза:

```bash
# Убедитесь что все коммиты запушены
git push origin master

# Создайте тэг
git tag -a v0.1.0 -m "Initial release"

# Запушьте тэг
git push origin v0.1.0
```

Go будет автоматически видеть тэги как версии.

## Использование в CI/CD

Для GitHub Actions:

```yaml
name: Build

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Configure Git for private packages
        run: |
          git config --global url."https://${{ secrets.GH_TOKEN }}@github.com/".insteadOf "https://github.com/"
      
      - name: Build
        run: go build -v ./...
        env:
          GOPRIVATE: github.com/Mr1rbis/*
```

## Структура модуля

Текущая структура Go модуля:

```
github.com/Mr1rbis/lemicraft-go (v0.1.0)
├── client.go         - Основной клиент
├── types.go          - Типы данных
├── options.go        - Опции конфигурации
├── players.go        - API для игроков
├── launcher.go       - API для лаунчера
├── news.go           - API новостей/контента
├── petitions.go      - API петиций
├── court.go          - API судебных дел
├── transport.go      - HTTP транспорт (опционально)
├── examples/
│   └── basic/
│       └── main.go   - Пример использования
├── README.md         - Документация
└── go.mod            - Определение модуля
```

## Тестирование

Запустите примеры с вашим API токеном:

```bash
go run ./examples/basic/main.go -token="your-api-token"
```

## Структура версий

- `v0.1.0` - Основная функциональность (5 сервисов)
- Планируемые:
  - `v0.2.0` - Дополнительные методы (поиск, фильтрация)
  - `v1.0.0` - Стабильный API

---

**Примечание:** Эта библиотека может быть используется в приватном режиме. Просто убедитесь что репозиторий на GitHub помечен как private и используйте аутентификацию как описано выше.

