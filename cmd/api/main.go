package main

import (
	"time"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/initializer"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/middleware"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/routes"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/validation"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadENV()
	initializer.InitDB()
	validation.Init()
}

func main() {

	server := gin.Default()
	server.Use(middleware.TimeoutMiddleware(15 * time.Second))
	routes.RegisterEventRoutes(server)
	server.Run(":8080")
}
