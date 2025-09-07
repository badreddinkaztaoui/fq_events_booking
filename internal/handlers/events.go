package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/models"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/repository"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	Repo *repository.EventRepo
}

func NewEventHandler(repo *repository.EventRepo) *EventHandler {
	return &EventHandler{Repo: repo}
}

func (h *EventHandler) GetAll(ctx *gin.Context) {
	events, err := h.Repo.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch events.",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (h *EventHandler) CreateEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload.",
		})
		return
	}
	if err := h.Repo.Create(ctx.Request.Context(), &event); err != nil {
		log.Printf("Failed to create event: %v\n", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create event.",
		})
		return
	}
	ctx.JSON(http.StatusCreated, event)
}

func (h *EventHandler) GetByID(ctx *gin.Context) {
	idInt, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event ID.",
		})
		return
	}

	event, err := h.Repo.GetByID(ctx.Request.Context(), idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not fetch event.",
		})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) Update(ctx *gin.Context) {
	idInt, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	var eventToUpdate models.Event
	if err := ctx.ShouldBindJSON(&eventToUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload."})
		return
	}

	if err := h.Repo.Update(ctx.Request.Context(), idInt, &eventToUpdate); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found."})
			return
		}

		log.Printf("Error updating event: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event."})
		return
	}

	ctx.JSON(http.StatusOK, eventToUpdate)
}

func (h *EventHandler) Delete(ctx *gin.Context) {
	idInt, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	if err := h.Repo.Delete(ctx.Request.Context(), idInt); err != nil {
		log.Printf("Error deleting event: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})
}
