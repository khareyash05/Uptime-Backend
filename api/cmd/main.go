package cmd

import (
	"log"

	"github.com/khareyash05/uptime-backend-db"
)

func InitDB() {
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
}
