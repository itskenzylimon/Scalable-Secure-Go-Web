# 1. Copy .env.example to .env and tweak for your environment
cp .env.example .env

# 2. Get dependencies
go mod init untitled
go get github.com/gofiber/fiber/v2 github.com/gofiber/fiber/v2/middleware/{logger,recover,cors,helmet,limiter} \
github.com/joho/godotenv github.com/spf13/viper gorm.io/gorm gorm.io/driver/sqlite

# 3. Run
go run ./cmd/api
