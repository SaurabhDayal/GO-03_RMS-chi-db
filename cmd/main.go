package main

import (
	"03_RMS/database"
	"03_RMS/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := server.SetupRoutes()

	if err := database.ConnectAndMigrate(
		"localhost",
		"5434",
		"rms-go",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		log.Fatal("Failed to initialize and migrate database with error: %+v", err)
	}
	fmt.Println("migration successful!!")

	log.Fatal(http.ListenAndServe(":8080", r))
}
