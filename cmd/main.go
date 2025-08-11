package main

import (
	api "digantara/internal/api"
	"digantara/internal/scheduler"
	"log"
)

func main() {

	scheduler.StartScheduler()
	scheduler.StartDbJobs()

	router := api.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on :8080")
}
