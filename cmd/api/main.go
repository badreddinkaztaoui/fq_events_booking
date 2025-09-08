package main

import (
	"log"
	"os"
	"time"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/database"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/middleware"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/routes"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dcs := os.Getenv("DATABASE_URL")
	if dcs == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}

	db, err := database.ConnectDB(dcs)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	validation.Init()

	server := gin.Default()
	server.Use(middleware.TimeoutMiddleware(15 * time.Second))

	routes.RegisterEventRoutes(server, db)

	server.Run(":8080")
}
