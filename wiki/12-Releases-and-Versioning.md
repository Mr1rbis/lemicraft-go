# Releases and Versioning

Проект использует SemVer теги (`vX.Y.Z`).

## Релизный чеклист

1. Обновить changelog/README при необходимости.
2. Убедиться, что `go build ./...` проходит.
3. Сделать commit.
4. Поставить тег.
5. Запушить branch и тег.

## Команды

```bash
git add -A
git commit -m "release: v0.2.0"
```

```bash
git tag -a v0.2.0 -m "Release v0.2.0"
```

```bash
git push origin master
git push origin v0.2.0
```

## Использование фиксированной версии

```bash
go get github.com/Mr1rbis/lemicraft-go@v0.2.0
```

