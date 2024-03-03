package main

import (
	"github.com/AbdurrahmanTalha/brainscape-backend-go/api"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/data/db"
)

func main() {
	cfg := config.SetupConfig()
	err := db.SetupDB(cfg)

	if err != nil {
		panic("[ERROR] Failed to connect to database")
	}

	api.SetupServer(cfg)
}
