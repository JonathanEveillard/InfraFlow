package main

import (
	"github.com/JonathanEveillard/InfraFlow/api"
)

func main() {
	// cfg := config.Load() // Read Settings
	// fmt.Println(cfg)
	// db := store.Connect(cfg) // PostgreSQL
	//px := proxmox.New(cfg) // Proxmox API Client

	// jobs := jobs.Start(db) // Job Queue + Scheduler
	// monitor.Start(db,px) // Metrics

	api.Server()

}
