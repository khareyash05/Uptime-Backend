package cmd

import (
	"log"

	db "github.com/khareyash05/uptime-backend-db"
)

func InitDB() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
}

func RunMigrations() {
	if err := db.RunMigrations(); err != nil {
		log.Fatal(err)
	}
}
