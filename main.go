package main

import (
	_ "Scalable-Secure-Go-Web/docs"
	"Scalable-Secure-Go-Web/internal/config"
	"Scalable-Secure-Go-Web/internal/handlers"
	"Scalable-Secure-Go-Web/internal/models"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var startTime = time.Now()

// @title Product Catalog API
// @version 1.0
// @description A simple GoFiber + GORM + Swagger API for managing products, categories, and brands.
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Setup DB (SQLite for demo; swap for Postgres/MySQL in prod)
	config.Connect(cfg)

	// Initialize Fiber
	app := fiber.New()

	//⃣ Global middleware
	var output io.Writer = os.Stdout

	// Optional file logging
	if cfg.LogToFile {
		file := handlers.SetupLogFile()
		output = io.MultiWriter(os.Stdout, file)
	}

	// Unified logger config
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] | ${status} | ${latency} | ${method} ${path} | ${ip}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Local",
		Output:     output,
	}))

	// Register Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Panic - proof the thing.
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, err interface{}) {
			log.Printf("🔥 Panic: %v", err)

			msg := "Internal server error"
			if cfg.Environment == "development" {
				msg = fmt.Sprintf("Panic: %v", err)
			}

			_ = c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
				Status:     "error",
				StatusCode: fiber.StatusInternalServerError,
				Data:       nil,
				Message:    msg,
			})
		},
	}))

	if cfg.Environment == "production" {

		// CORS – only if origins are configured
		if len(cfg.FrontendOrigins) > 0 {
			app.Use(cors.New(cors.Config{
				AllowOrigins:     join(cfg.FrontendOrigins, ","),
				AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
				AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
				AllowCredentials: cfg.CORSAllowCreds,
			}))
		}

		// ⃣ Helmet – optional toggle
		if cfg.EnableHelmet {
			app.Use(helmet.New(helmet.Config{
				XSSProtection:             "1; mode=block", // Protect against reflected XSS attacks
				ContentTypeNosniff:        "nosniff",       // Prevent MIME-type sniffing
				XFrameOptions:             "DENY",          // Prevent clickjacking (e.g., iframe embedding)
				ReferrerPolicy:            "no-referrer",   // Control how much referrer info is sent
				CrossOriginEmbedderPolicy: "unsafe-none",
				CrossOriginOpenerPolicy:   "same-origin-allow-popups",
				CrossOriginResourcePolicy: "cross-origin",
			}))
		}

		// 4 Rate limiter – optional toggle
		if cfg.EnableLimiter {
			app.Use(limiter.New(limiter.Config{
				Max:          cfg.RateLimitMax,
				Expiration:   cfg.RateLimitWindow,
				LimitReached: tooManyRequestsJSON,
			}))
		}

	} else {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*", // or restrict with a comma-separated list
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	// Fake route to prove it works
	app.Get("/health", healthJSON)

	// API version group
	api := app.Group("/api/v1")

	// Product routes group
	productApi := api.Group("/products")
	productApi.Get("/", handlers.GetAllProducts)
	productApi.Get("/:id", handlers.GetProductByID)
	productApi.Post("/", handlers.CreateProduct)
	productApi.Put("/:id", handlers.UpdateProduct)
	productApi.Delete("/:id", handlers.DeleteProduct)

	// Category routes group
	categoryApi := api.Group("/categories")
	categoryApi.Get("/", handlers.GetAllCategories)
	categoryApi.Get("/:id", handlers.GetCategoryByID)
	categoryApi.Post("/", handlers.CreateCategory)
	categoryApi.Put("/:id", handlers.UpdateCategory)
	categoryApi.Delete("/:id", handlers.DeleteCategory)

	// Brand routes group
	brandApi := api.Group("/brands")
	brandApi.Get("/", handlers.GetAllBrands)
	brandApi.Get("/:id", handlers.GetBrandByID)
	brandApi.Post("/", handlers.CreateBrand)
	brandApi.Put("/:id", handlers.UpdateBrand)
	brandApi.Delete("/:id", handlers.DeleteBrand)

	//⃣ Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("⇨ Listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func healthJSON(c *fiber.Ctx) error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return c.JSON(fiber.Map{
		"status":       "ok",
		"uptime":       time.Since(startTime).String(),
		"go_version":   runtime.Version(),
		"os":           runtime.GOOS,
		"arch":         runtime.GOARCH,
		"cpu_cores":    runtime.NumCPU(),
		"goroutines":   runtime.NumGoroutine(),
		"memory_alloc": memStats.Alloc,
		"memory_total": memStats.TotalAlloc,
		"memory_sys":   memStats.Sys,
		"heap_objects": memStats.HeapObjects,
		"gc_count":     memStats.NumGC,
		"last_gc":      memStats.LastGC,
	})
}

// tooManyRequestsJSON returns 429 with a JSON message.
func tooManyRequestsJSON(c *fiber.Ctx) error {
	return c.Status(fiber.StatusTooManyRequests).JSON(models.APIResponse{
		Status:     "error",
		StatusCode: fiber.StatusTooManyRequests,
		Data:       nil,
		Message:    "Too many requests. Calm down, champ.",
	})
}

// join is a tiny helper (because strings.Join needs []string anyway)
func join(ss []string, sep string) string {
	if len(ss) == 0 {
		return ""
	}
	out := ss[0]
	for _, s := range ss[1:] {
		out += sep + s
	}
	return out
}
