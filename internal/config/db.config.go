package config

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"Scalable-Secure-Go-Web/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *App) {
	// Configure custom logger for GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Log slow queries
			LogLevel:                  logger.Warn,            // Use Info or Error as needed
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// GORM config with improved safety/logging
	gormConfig := &gorm.Config{
		Logger: newLogger,
		// NamingStrategy: schema.NamingStrategy{}, // Optional: customize naming
		DisableForeignKeyConstraintWhenMigrating: false, // Let GORM manage FKs
		// Add more config if needed
	}

	var err error
	switch cfg.DBDriver {
	case "postgres":
		DB, err = gorm.Open(postgres.Open(cfg.DBDSN), gormConfig)

	case "mysql":
		DB, err = gorm.Open(mysql.Open(cfg.DBDSN), gormConfig)

	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(cfg.DBDSN), gormConfig)

	default:
		log.Fatalf("❌ Unsupported DB driver: %s", cfg.DBDriver)
	}

	if err != nil {
		log.Fatalf("❌ Failed to connect to database (%s): %v", cfg.DBDriver, err)
	}

	// Run migrations (in correct order)
	if err := DB.AutoMigrate(
		&models.Brand{},
		&models.Category{},
		&models.Product{},
	); err != nil {
		log.Fatalf("❌ Failed to auto-migrate database: %v", err)
	}

	log.Println("✅ Database connection & migration successful")
}
