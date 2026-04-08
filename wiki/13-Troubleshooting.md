# Troubleshooting

## `invalid version` / `could not read Username for 'https://github.com'`

Причина: приватный репозиторий без настроенной аутентификации.

Решение: страница [Private Repository](11-Private-Repository).

## `src refspec main does not match any`

Причина: локальная ветка называется не `main` (например, `master`) или нет коммитов.

Проверка:

```bash
git branch
git log --oneline -n 1
```

## `json: cannot unmarshal ...`

Причина: формат данных API может отличаться от ожидаемого.

Что уже учтено в библиотеке:

- `uuid` парсится гибко через `UUIDValue`
- `registered`/`lastSeen` в Plan парсятся через `APITime`

## Что собирать для диагностики issue

- Версия библиотеки
- Фрагмент кода вызова
- Полный текст ошибки
- Endpoint и пример response (без токена)

