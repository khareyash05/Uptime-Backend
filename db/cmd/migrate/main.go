package main

import (
	"log"

	"github.com/khareyash05/uptime-backend-db"
)

func main() {
	// Initialize database and run migrations
	if err := db.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	log.Println("Successfully ran database migrations")
}
