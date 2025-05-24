package db

import (
	"fmt"
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

var (
	DB *gorm.DB
)

// Init initializes the database connection
func Init() error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("failed to initialize zap logger: %v", err)
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
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	zapLogger.Info("Successfully connected to database")

	// Run migrations
	if err := db.AutoMigrate(&models.User{}, &models.Website{}, &models.WebsiteTick{}, &models.Validator{}); err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	zapLogger.Info("Successfully ran database migrations")

	// Set the global DB variable
	DB = db
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
