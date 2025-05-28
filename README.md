# SnipLink API - REST сервис для сокращения ссылок на Go  

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/GoX7/SnipLink-api/go.yml)](https://github.com/GoX7/SnipLink-api/actions)

**🔥 Основные аспекты:**  
- ✂️ Создание коротких ссылок  
- 🔀 Редирект по алиасам  
- 🧩 Простая интеграция  
- 📝 Комментированный код  

**🛠 Технологии:**  
`github.com/go-chi/chi/v5` `github.com/go-playground/validator`  
`github.com/ilyakaznacheev/cleanenv` `log/slog` `database/sql`  

**ℹ️ Версия:** 1.0.0  
**👤 Автор:** GoX7  
**📜 Лицензия:** MIT (файл LICENSE)  

---

## 🌐 API Endpoints  
**Проверка работы API**  
```http 
GET / 
```
Перейти по короткой ссылке
``` http
GET /l/{alias}
```

Создать короткую ссылку
``` http
Content-Type: application/json

{
  "link": "https://example.com"
}
```

## Конфигурация
``` yaml
path: 
  server_log: "logs/..."
  sqlite_log: "logs/..."
  mw_log:     "logs/..."
server:
  addr:       "host:port"
  wto:        10s #write time out
  rto:        10s #read time out
logger:
  level:      "" #debug, info, warn, error
```

## Структура
```
.
├── config/
│   └── config.yaml         # Настройка конфигурации
├── internal/
│   ├── config              # Конфигурация
│   ├── controlers          # HTTP обработчики
│   ├── logger              # Логгер
│   └── sqlite              # Работа с БД
├── pkg/
│   ├── response            # Помощник в json ответах
│   └── mw_logger           # Логирование запросов
├── go.mod
├── go.sum
└── main.go  # Entry point
```
