package routes

import (
	"database/sql"
	"net/http"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/handlers"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/repository"
	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(server *gin.Engine, db *sql.DB) {
	eventsRepo := repository.NewEventRepo(db)
	eventsHandler := handlers.NewEventHandler(eventsRepo)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.GET("/events", eventsHandler.GetAll)
	server.GET("/events/:id", eventsHandler.GetByID)
	server.POST("/events", eventsHandler.CreateEvent)
	server.PUT("/events/:id", eventsHandler.Update)
	server.DELETE("/events/:id", eventsHandler.Delete)
}
