# SnipLink API - REST service for shortening links in Go  

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/GoX7/SnipLink-api/go.yml)](https://github.com/GoX7/SnipLink-api/actions)

**🔥 Key Features:**  
- ✂️ Creating short links  
- 🔀 Redirecting by aliases  
- 🧩 Easy integration  
- 📝 Well-commented code  

**🛠 Technologies:**  
`github.com/go-chi/chi/v5` `github.com/go-playground/validator`  
`github.com/ilyakaznacheev/cleanenv` `log/slog` `database/sql`  

**ℹ️ Version:** 1.0.0  
**👤 Author:** GoX7  
**📜 License:** MIT (LICENSE file)  

---

## API Endpoints  
**Check API Status**  
```http 
GET / 
```
**Redirect to short link**  
```http
GET /l/{alias}
```

**Create a short link**  
```http
Content-Type: application/json

{
  "link": "https://example.com"
}
```

## Configuration
```yaml
path: 
  server_log: "logs/..."
  sqlite_log: "logs/..."
  mw_log:     "logs/..."
server:
  addr:       "host:port"
  wto:        10s # write time out
  rto:        10s # read time out
logger:
  level:      "" # debug, info, warn, error
```

## Structure
```
.
├── config/
│   └── config.yaml         # Configuration settings
├── internal/
│   ├── config              # Configuration
│   ├── controllers         # HTTP handlers
│   ├── logger              # Logger
│   └── sqlite              # Database operations
├── pkg/
│   ├── response            # JSON response helper
│   └── mw_logger           # Request logging
├── go.mod
├── go.sum
└── main.go                 # Entry point
```
