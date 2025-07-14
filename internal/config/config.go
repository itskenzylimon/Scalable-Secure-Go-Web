package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// App holds all runtime configuration.
type App struct {
	Port        int
	Environment string

	FrontendOrigins []string
	CORSAllowCreds  bool
	RateLimitMax    int
	RateLimitWindow time.Duration
	EnableHelmet    bool
	EnableLimiter   bool

	DBDriver string
	DBDSN    string

	LogToFile bool
}

// Load reads .env / environment and returns a populated App struct.
func Load() (*App, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  .env file not found or unreadable. Falling back to OS environment...", err)
	} else {
		log.Printf("✅ Loaded config from: %s\n", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("CORS_ALLOW_CREDENTIALS", true)
	viper.SetDefault("RATE_LIMIT_MAX", 100)
	viper.SetDefault("RATE_LIMIT_WINDOW", "1m")
	viper.SetDefault("ENABLE_HELMET", true)
	viper.SetDefault("ENABLE_RATE_LIMITER", true)
	viper.SetDefault("DB_DRIVER", "sqlite")
	viper.SetDefault("DB_DSN", "catalog.db")
	viper.SetDefault("LOG_TO_FILE", true)

	// Parse duration safely
	windowStr := viper.GetString("RATE_LIMIT_WINDOW")
	window, err := time.ParseDuration(windowStr)
	if err != nil {
		log.Printf("❌ Failed to parse RATE_LIMIT_WINDOW '%s': %v\n", windowStr, err)
		return nil, err
	}

	// Debug log: Print loaded values
	log.Println("    Loaded Configuration:")
	log.Printf("   APP_PORT: %d\n", viper.GetInt("APP_PORT"))
	log.Printf("   ENVIRONMENT: %s\n", viper.GetString("ENVIRONMENT"))
	log.Printf("   DB_DRIVER: %s\n", viper.GetString("DB_DRIVER"))
	log.Printf("   DB_DSN: %s\n", viper.GetString("DB_DSN"))
	log.Printf("   LOG_TO_FILE: %v\n", viper.GetBool("LOG_TO_FILE"))
	log.Printf("   FRONTEND_ORIGINS: %v\n", viper.GetStringSlice("FRONTEND_ORIGINS"))
	log.Printf("   RATE_LIMIT_MAX: %d\n", viper.GetInt("RATE_LIMIT_MAX"))
	log.Printf("   RATE_LIMIT_WINDOW: %s\n", window)
	log.Printf("   ENABLE_HELMET: %v\n", viper.GetBool("ENABLE_HELMET"))
	log.Printf("   ENABLE_RATE_LIMITER: %v\n", viper.GetBool("ENABLE_RATE_LIMITER"))

	// Return the populated config
	return &App{
		Port:            viper.GetInt("APP_PORT"),
		Environment:     viper.GetString("ENVIRONMENT"),
		FrontendOrigins: viper.GetStringSlice("FRONTEND_ORIGINS"),
		CORSAllowCreds:  viper.GetBool("CORS_ALLOW_CREDENTIALS"),
		RateLimitMax:    viper.GetInt("RATE_LIMIT_MAX"),
		RateLimitWindow: window,
		EnableHelmet:    viper.GetBool("ENABLE_HELMET"),
		EnableLimiter:   viper.GetBool("ENABLE_RATE_LIMITER"),
		DBDriver:        viper.GetString("DB_DRIVER"),
		DBDSN:           viper.GetString("DB_DSN"),
		LogToFile:       viper.GetBool("LOG_TO_FILE"),
	}, nil
}
