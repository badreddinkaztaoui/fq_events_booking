package handlers

import (
	"log"
	"net/http"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/auth"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/models"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/repository"
	"github.com/badreddinkaztaoui/fq_events_booking/internal/validation"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	repo *repository.UsersRepo
}

type LoginCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func NewUserHandler(repo *repository.UsersRepo) *UsersHandler {
	return &UsersHandler{repo: repo}
}

func (h *UsersHandler) SignUp(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse user data",
		})
		return
	}

	if err := validation.Validate.Struct(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validation.FormatValidationErrors(err),
		})
		return
	}

	usr, _ := h.repo.GetByEmail(ctx.Request.Context(), user.Email)
	if usr != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "User with this email already exists",
		})
		return
	}

	plainTextPassword := user.Password
	if err := user.HashPassword(plainTextPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process password",
		})
		return
	}

	err = h.repo.Create(ctx.Request.Context(), &user)
	if err != nil {
		log.Printf("Failed to create user: %v\n", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create user",
		})
		return
	}

	user.Password = ""

	token, err := auth.NewAuthToken(uint(user.ID))
	if err != nil {
		log.Printf("Failed to create auth token: %v\n", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create auth token",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
		"user":    user,
	})
}

func (h *UsersHandler) SignIn(ctx *gin.Context) {
	var credentials LoginCredentials

	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	user, err := h.repo.GetByEmail(ctx.Request.Context(), credentials.Email)
	if err != nil || user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	err = user.CheckPassword(credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}
	user.Password = ""

	token, err := auth.NewAuthToken(uint(user.ID))
	if err != nil {
		log.Printf("Failed to generate token for user %d: %v\n", user.ID, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})

}
