
---

## Building a Scalable, Secure Go Web App


![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-2.x-29ABE2?style=flat&logo=fiber)
![GORM](https://img.shields.io/badge/GORM-ORM-FCA121?style=flat&logo=sqlite)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

A high-performance product catalog API built with **Fiber** (Go web framework), **GORM** (ORM), and layered with middleware for **security, observability, and scalability**.

Designed to be scalable, developer-friendly, and ready for production from day one.


- ✅ Fully RESTful Product/Brand/Category schema
- 🔐 Middleware stack: logger, CORS, helmet, recover, rate-limiter
- 🌍 Environment-based config (`.env` or system env)
- 🧩 GORM auto-migrations (SQLite by default, easily swappable)
- 🚀 Fiber HTTP server with sane defaults
- 🧪 Ready for Swagger integration & validation (`example` & `validate` tags)

---

## 📁 Project Structure

```

.
├── cmd/
│   └── api/            # Entry point (main.go)
├── internal/
│   ├── config/         # Loads env vars and runtime settings
│   └── models/         # Product, Brand, Category structs
├── .env.example        # Default environment variables
├── go.mod / go.sum     # Module deps
└── README.md           # You are here

````

---

## Tech Stack

| Layer         | Tool                     |
|---------------|--------------------------|
| Web Framework | [Fiber](https://gofiber.io/) |
| ORM           | [GORM](https://gorm.io/) |
| Config        | [Viper](https://github.com/spf13/viper) |
| Env Support   | [godotenv](https://github.com/joho/godotenv) |
| Database      | SQLite (local), Postgres/MySQL (prod) |
| Validation    | Tags (`validate`, `binding`) |

---

## Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/itskenzylimon/Scalable-Secure-Go-Web.git
cd fiber-catalog-api
````

### 2. Setup environment

```bash
cp .env.example .env
```

Edit `.env` to match your needs (e.g. port, frontend origins, security toggles).

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the app

```bash
go run ./cmd/api
```

Visit 👉 [http://localhost:8080/health](http://localhost:8080/health)

---

## 📦 API Entities

### 📦 Product

```json
{
  "id": 1,
  "name": "iPhone 14",
  "description": "Latest Apple smartphone",
  "price": 999.99,
  "cover_image": "https://example.com/images/iphone14.jpg",
  "category_id": 2,
  "brand_id": 1,
  "created_at": "2025-07-09T15:04:05Z"
}
```

### 🧢 Brand

```json
{
  "id": 1,
  "name": "Apple",
  "cover_image": "https://example.com/images/apple.png"
}
```

### Category

```json
{
  "id": 1,
  "title": "Smartphones",
  "cover_image": "https://example.com/images/smartphones.jpg"
}
```

---

## Middleware Stack (Always On Guard)

| Middleware | Purpose                                           |
| ---------- | ------------------------------------------------- |
| `logger`   | Logs every request (method, path, duration, etc.) |
| `recover`  | Prevents panics from crashing the app             |
| `cors`     | Dynamically sets allowed frontend origins         |
| `helmet`   | Adds secure HTTP headers                          |
| `limiter`  | Basic rate limiting (per IP per minute)           |

All toggled via `.env` 🔧

---

##️ Env Configuration

| Key                   | Description                            | Example                             |
| --------------------- | -------------------------------------- | ----------------------------------- |
| `APP_PORT`            | Server port                            | `8080`                              |
| `FRONTEND_ORIGINS`    | Allowed CORS origins (comma-separated) | `https://app.com,https://admin.app` |
| `ENABLE_HELMET`       | Enable secure headers                  | `true`                              |
| `ENABLE_RATE_LIMITER` | Enable IP-based rate limiting          | `true`                              |
| `RATE_LIMIT_MAX`      | Max requests per window                | `100`                               |
| `RATE_LIMIT_WINDOW`   | Window duration (e.g. `1m`)            | `1m`                                |

---

## Tests & Swagger (Coming Soon)

* Add [Swaggo](https://github.com/swaggo/swag) for auto-generated API docs
* Add unit/integration tests using `httptest`, `testify`, etc.

---

## Dev Notes

* Swap `sqlite` with Postgres or MySQL by replacing:

  ```go
  gorm.Open(sqlite.Open("catalog.db"), ...)
  ```

  with:

  ```go
  gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), ...)
  ```

* To run in production, set:

    * `ENABLE_HELMET=true`
    * `ENABLE_RATE_LIMITER=true`
    * Use a real DB connection pool

---

## License

MIT © [Your Name](https://github.com/yourusername)

---

## ❤️ Contributing

PRs welcome! Please open an issue first for discussion.

> Built with 💙 and panic recovery
