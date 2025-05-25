package main

import (
	"github.com/khareyash05/uptime-backend-api/cmd"
)

func main() {
	cmd.InitDB()
	cmd.RunMigrations()
}
