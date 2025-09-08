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
	userRepo := repository.NewUsersRepo(db)

	eventsHandler := handlers.NewEventHandler(eventsRepo)
	userHandler := handlers.NewUserHandler(userRepo)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.POST("/signup", userHandler.SignUp)
	server.POST("/signin", userHandler.SignIn)

	server.GET("/events", eventsHandler.GetAll)
	server.GET("/events/:id", eventsHandler.GetByID)
	server.POST("/events", eventsHandler.CreateEvent)
	server.PUT("/events/:id", eventsHandler.Update)
	server.DELETE("/events/:id", eventsHandler.Delete)
}
