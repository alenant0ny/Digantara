package main

import (
	api "digantara/internal/api"
	"digantara/internal/db"
	"digantara/internal/scheduler"
	"fmt"
	"log"
)

func main() {

	fmt.Println("hello world")
	_, dbErr := db.Connect()
	if dbErr != nil {
		log.Fatalf("Failed to connect to DB: %v", dbErr)
	}

	scheduler.StartScheduler()

	router := api.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	log.Println("Server running on :8080")
}
