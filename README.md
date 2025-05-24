# Цитатник — мини‑REST сервис на Go 📚

Проект демонстрирует базовые принципы Clean Architecture: чистый домен, слой use‑case, HTTP‑адаптер и in‑memory хранилище.

---

## Требования

* **Go 1.24**

---

## Структура репозитория

```
quote-service/
├── cmd/server/          # точка входа (main.go)
├── db/quotes.json       # начальный набор цитат 
├── internal/
│   ├── domain/          # сущности
│   ├── usecase/         # бизнес‑логика (QuoteService) + интерфейсы
│   └── adapters/
│       ├── http/        # обработчики, DTO
│       └── storage/     # in‑memory хранилище (+ JSON‑файл )
└── go.mod
```

---

## Быстрый запуск

```bash
git clone https://github.com/BigMakClub/quote-service.git
cd quote-service
make build
make run  
```

### Хранение в файле

Записи цитат хранятся в JSON‑файл:

```go
repo, _ := storage.NewStorage("./db/quotes.json")
```

Файл находится в db/quotes.json .

---

## Примеры запросов

```bash
# 1. Добавить
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius","quote":"Life is simple, but we insist on making it complicated."}'

# 2. Получить все
curl http://localhost:8080/quotes

# 3. Фильтр по автору
curl "http://localhost:8080/quotes?author=Confucius"

# 4. Случайная
curl http://localhost:8080/quotes/random

# 5. Удалить
curl -X DELETE http://localhost:8080/quotes/1
```

---
## Тесты
```bash
make test  
```