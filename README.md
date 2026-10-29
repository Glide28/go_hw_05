# Домашнее задание №5 по Go

**Группа:** BHEMBD-25  
**Студент:** Арсений Паниклов

Тема: «Разработка веб-приложений».

В репозитории находятся два Go-модуля: бизнес-логика `ledger` и REST API
`gateway`. Данные хранятся в памяти.

## Запуск

```powershell
go run ./gateway/cmd/gateway
```

Сервер будет доступен по адресу `http://localhost:8080`.

## Эндпоинты

| Метод | Путь | Назначение |
|---|---|---|
| `POST` | `/api/transactions` | создать транзакцию |
| `GET` | `/api/transactions` | получить транзакции |
| `POST` | `/api/budgets` | создать или обновить бюджет |
| `GET` | `/api/budgets` | получить бюджеты |

Все ответы имеют тип `application/json; charset=utf-8`. Ошибки возвращаются
в формате `{"error":"описание"}`.

## Примеры cURL

Создание бюджета:

```powershell
curl.exe -X POST http://localhost:8080/api/budgets `
  -H "Content-Type: application/json" `
  -d '{"category":"еда","limit":5000}'
```

Создание транзакции:

```powershell
curl.exe -X POST http://localhost:8080/api/transactions `
  -H "Content-Type: application/json" `
  -d '{"amount":450,"category":"еда","description":"ланч","date":"2026-10-10"}'
```

Получение данных:

```powershell
curl.exe http://localhost:8080/api/transactions
curl.exe http://localhost:8080/api/budgets
```

## Проверка

```powershell
go test ./ledger/... ./gateway/...
go vet ./ledger/... ./gateway/...
```

Тесты проверяют успешный сценарий, валидацию, пустые массивы, заголовок
ответа и возврат `409 Conflict` при превышении бюджета.
