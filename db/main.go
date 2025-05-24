package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/khareyash05/uptime-backend-db/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to initialize zap logger:", err)
	}
	defer zapLogger.Sync()

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		zapLogger.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		zapLogger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		zapLogger.Fatal("Failed to get database instance", zap.Error(err))
	}

	if err := sqlDB.Ping(); err != nil {
		zapLogger.Fatal("Failed to ping database", zap.Error(err))
	}

	zapLogger.Info("Successfully connected to database")

	// Run migrations
	if err := db.AutoMigrate(&models.User{}, &models.Website{}, &models.WebsiteTick{}, &models.Validator{}); err != nil {
		zapLogger.Fatal("Failed to run migrations", zap.Error(err))
	}

	zapLogger.Info("Successfully ran database migrations")
}
