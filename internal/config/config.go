package config

import (
	"time"

	"github.com/spf13/viper"
)

// App holds all runtime configuration.
type App struct {
	Port            int
	FrontendOrigins []string
	CORSAllowCreds  bool
	RateLimitMax    int
	RateLimitWindow time.Duration
	EnableHelmet    bool
	EnableLimiter   bool
}

// Load reads .env / environment and returns a populated App struct.
func Load() (*App, error) {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig() // Ignore error if .env doesn't exist – we fall back to OS env

	viper.AutomaticEnv()

	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("CORS_ALLOW_CREDENTIALS", true)
	viper.SetDefault("RATE_LIMIT_MAX", 100)
	viper.SetDefault("RATE_LIMIT_WINDOW", "1m")
	viper.SetDefault("ENABLE_HELMET", true)
	viper.SetDefault("ENABLE_RATE_LIMITER", true)

	window, err := time.ParseDuration(viper.GetString("RATE_LIMIT_WINDOW"))
	if err != nil {
		return nil, err
	}

	return &App{
		Port:            viper.GetInt("APP_PORT"),
		FrontendOrigins: viper.GetStringSlice("FRONTEND_ORIGINS"),
		CORSAllowCreds:  viper.GetBool("CORS_ALLOW_CREDENTIALS"),
		RateLimitMax:    viper.GetInt("RATE_LIMIT_MAX"),
		RateLimitWindow: window,
		EnableHelmet:    viper.GetBool("ENABLE_HELMET"),
		EnableLimiter:   viper.GetBool("ENABLE_RATE_LIMITER"),
	}, nil
}
