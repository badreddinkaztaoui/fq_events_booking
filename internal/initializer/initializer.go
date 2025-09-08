package initializer

import (
	"log"
	"os"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/database"
	"github.com/joho/godotenv"
)

func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func InitDB() {
	dcs := os.Getenv("DATABASE_URL")
	if dcs == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	database.ConnectDB(dcs)
}
