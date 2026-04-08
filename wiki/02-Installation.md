# Installation

```
go work use /home/mrirbis/GolandProjects/LemicraftGo
go work use .
go work init
```bash

## go work (опционально)

```
replace github.com/Mr1rbis/lemicraft-go => /home/mrirbis/GolandProjects/LemicraftGo
```go

## Локальная разработка через replace

```
go build ./...
go mod tidy
```bash

## Проверка установки

```
require github.com/Mr1rbis/lemicraft-go latest

go 1.21

module your-app
```go

## go.mod

```
go get github.com/Mr1rbis/lemicraft-go
```bash

## Установка модуля

