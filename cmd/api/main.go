package main

import (
	"log"
	"os"
	"time"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/database"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/middleware"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dcs := os.Getenv("DATABASE_URL")
	if dcs == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := database.ConnectDB(dcs)
	if err != nil {
		log.Fatalf("can't connect to database: %v", err)
	}
	defer db.Close()

	server := gin.Default()
	server.Use(middleware.TimeoutMiddleware(15 * time.Second))
	routes.RegisterEventRoutes(server, db)
	server.Run(":8080")
}
