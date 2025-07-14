
---

## Building a Scalable, Secure Go Web App


![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-2.x-29ABE2?style=flat&logo=fiber)
![GORM](https://img.shields.io/badge/GORM-ORM-FCA121?style=flat&logo=sqlite)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

A high-performance product catalog API built with **Fiber** (Go web framework), **GORM** (ORM), and layered with middleware for **security, observability, and scalability**.

Designed to be scalable, developer-friendly, and ready for production from day one.


- ✅ Fully RESTful Product/Brand/Category schema
- ✅ Middleware stack: logger, CORS, helmet, recover, rate-limiter
- ✅ Environment-based config (`.env` or system env)
- ✅ GORM auto-migrations (SQLite by default, easily swappable)
- ✅ Fiber HTTP server with sane defaults
- ✅ Ready for Swagger integration & validation (`example` & `validate` tags)

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
go run .main.go
```
### 🔄 JSON Metrics (Optional)

Visit 👉 [http://localhost:8080/health](http://localhost:8080/health)

Returns structured runtime metrics in JSON for health checks, CI/CD, or monitoring agents:


| Metric              | Description                                      |
|---------------------|--------------------------------------------------|
| Uptime              | How long the server has been running             |
| Memory Usage        | Heap, stack, and total memory allocated          |
| Goroutines          | Number of active Go routines                     |
| Threads             | Number of OS threads used                        |
| CPU Cores           | Total available processor cores                  |
| Request Count       | Total number of handled HTTP requests            |
| Response Times      | Avg. latency across all requests                 |
| Status Code Summary | Breakdown of 1xx, 2xx, 3xx, 4xx, 5xx responses   |
| Last GC             | Time of the last garbage collection              |

## 📦 API Entities

### 📦 Product

Represents an item in your catalog with linked `Brand` and `Category`.

```json
{
  "id": 1,
  "name": "iPhone 14",
  "description": "Latest Apple smartphone",
  "price": 999.99,
  "cover_image": "https://example.com/images/iphone14.jpg",
  "category_id": 2,
  "brand_id": 1,
  "category": {
    "id": 2,
    "title": "Smartphones",
    "cover_image": "https://example.com/images/smartphones.jpg"
  },
  "brand": {
    "id": 1,
    "name": "Apple",
    "cover_image": "https://example.com/images/apple.png"
  },
  "created_at": "2025-07-09T15:04:05Z",
  "updated_at": "2025-07-09T15:04:05Z"
}
```

### 🧢 Brand

A brand/manufacturer associated with one or more products.

```json
{
  "id": 1,
  "name": "Apple",
  "cover_image": "https://example.com/images/apple.png",
  "created_at": "2025-07-09T15:04:05Z",
  "updated_at": "2025-07-09T15:04:05Z"
}
```

### Category

Used to group related products (e.g. smartphones, tablets, accessories).

```json
{
  "id": 2,
  "title": "Smartphones",
  "cover_image": "https://example.com/images/smartphones.jpg",
  "created_at": "2025-07-09T15:04:05Z",
  "updated_at": "2025-07-09T15:04:05Z"
}
```
## 📬 API Routes

Base URL: `/api/v1`

> All endpoints return standardized JSON using the `APIResponse` format:
>
> ```json
> {
>   "status": "success",
>   "status_code": 200,
>   "data": {},
>   "message": "Product created successfully"
> }
> ```

---

### Products

| Method | Route                | Description              |
|--------|----------------------|--------------------------|
| GET    | `/products`          | Get all products         |
| GET    | `/products/:id`      | Get a product by ID      |
| POST   | `/products`          | Create a new product     |
| PUT    | `/products/:id`      | Update an existing product |
| DELETE | `/products/:id`      | Delete a product         |

---

### Categories

| Method | Route                 | Description               |
|--------|-----------------------|---------------------------|
| GET    | `/categories`         | Get all categories        |
| GET    | `/categories/:id`     | Get a category by ID      |
| POST   | `/categories`         | Create a new category     |
| PUT    | `/categories/:id`     | Update a category         |
| DELETE | `/categories/:id`     | Delete a category         |

---

### Brands

| Method | Route             | Description              |
|--------|-------------------|--------------------------|
| GET    | `/brands`         | Get all brands           |
| GET    | `/brands/:id`     | Get a brand by ID        |
| POST   | `/brands`         | Create a new brand       |
| PUT    | `/brands/:id`     | Update a brand           |
| DELETE | `/brands/:id`     | Delete a brand           |

---

---

## 🛡️ Middleware Stack (Always On Guard)

| Middleware | Purpose                                                                 |
|------------|-------------------------------------------------------------------------|
| `logger`   | Logs every request (`method`, `path`, `status`, `latency`, `IP`)        |
|            | Supports log to file via `LOG_TO_FILE=true`                             |
| `recover`  | Catches panics, logs stack traces, and returns standardized error JSON  |
|            | Stack traces enabled for debugging; customizable error structure        |
| `cors`     | Dynamically allows frontend origins via `FRONTEND_ORIGINS`              |
|            | Supports multiple origins and credential mode (`CORS_ALLOW_CREDENTIALS`)|
| `helmet`   | Adds secure HTTP headers (CSP, XSS protection, COOP/CORP, etc.)         |
|            | Toggled by `ENABLE_HELMET`                                              |
| `limiter`  | IP-based rate limiting using `RATE_LIMIT_MAX` and `RATE_LIMIT_WINDOW`   |
|            | Toggled by `ENABLE_RATE_LIMITER`                                        |

All middleware is configured via environment variables in `.env`.

---

## Env Configuration

| Key                    | Description                                     | Example                                                  |
|------------------------|-------------------------------------------------|----------------------------------------------------------|
| APP_PORT               | Port where the Fiber server listens             | 8080                                                     |
| ENVIRONMENT            | App environment (`development`, `production`)  | development                                              |
| FRONTEND_ORIGINS       | Allowed CORS origins (comma-separated)         | https://app.myfrontend.com,https://admin.myfrontend.com |
| CORS_ALLOW_CREDENTIALS | Allow CORS with credentials                    | true                                                     |
| ENABLE_HELMET          | Enable HTTP security headers via Helmet        | true                                                     |
| ENABLE_RATE_LIMITER    | Enable IP-based rate limiting                  | true                                                     |
| RATE_LIMIT_MAX         | Max requests per rate window                   | 100                                                      |
| RATE_LIMIT_WINDOW      | Duration of rate limiting window               | 1m                                                       |
| LOG_TO_FILE            | Enable logging to logs/server.log             | false                                                    |
| DB_DRIVER              | Database driver (`sqlite`, `postgres`, `mysql`) | sqlite                                                  |
| DB_DSN                 | Connection string for selected DB              | ./catalog.db (or DSN for PostgreSQL/MySQL)              |
---

## Tests & Swagger (Coming Soon)

* Add [Swaggo](https://github.com/swaggo/swag) for auto-generated API docs
* Add unit/integration tests using `httptest`, `testify`, etc.

---

## Dev Notes

- ✅ The database driver and connection are now configured via environment variables:

    ```env
    DB_DRIVER=sqlite         # Options: sqlite, postgres, mysql
    DB_DSN=./catalog.db      # Or use full DSN string for Postgres/MySQL
    ```

  Example for PostgreSQL:

    ```env
    DB_DRIVER=postgres
    DB_DSN=host=localhost user=postgres password=secret dbname=catalog port=5432 sslmode=require TimeZone=UTC
    ```

  Example for MySQL:

    ```env
    DB_DRIVER=mysql
    DB_DSN=root:root@tcp(127.0.0.1:3306)/catalog?charset=utf8mb4&parseTime=True&loc=Local&tls=true
    ```

- ✅ To change databases in code, the selection is handled automatically:

    ```go
    switch cfg.DBDriver {
    case "postgres":
        gorm.Open(postgres.Open(cfg.DBDSN), ...)
    case "mysql":
        gorm.Open(mysql.Open(cfg.DBDSN), ...)
    default:
        gorm.Open(sqlite.Open(cfg.DBDSN), ...)
    }
    ```

- ✅ To run in production, make sure the following are set in `.env`:

    ```env
    ENVIRONMENT=production
    ENABLE_HELMET=true
    ENABLE_RATE_LIMITER=true
    LOG_TO_FILE=true
    ```

- ✅ Connection pooling can be tuned inside `config.Connect()`:

    ```go
    sqlDB.SetMaxOpenConns(50)
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetConnMaxLifetime(time.Hour)
    ```

- ✅ Migrations auto-run on startup via:

    ```go
    DB.AutoMigrate(&Brand{}, &Category{}, &Product{})
    ```

  Models should include appropriate `gorm:"index"` and `gorm:"unique"` tags for indexing and constraints.

---

💡 Pro tip: Commit a `config.env.example` file for reference and exclude your actual `.env` in `.gitignore`.


## License

MIT © [Kevin Limon](https://github.com/itskenzyliom)

---

## ❤️ Contributing

PRs welcome! Please open an issue first for discussion.

> Built with 💙 and panic recovery
