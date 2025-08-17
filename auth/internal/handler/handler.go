package handler

import (
	"auth/internal/config"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) RegisterUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user is successfully registered",
	})
}

func (h *Handler) LoginUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":       "user successfully authorized",
		"access_token":  "example_access_token",
		"refresh_token": "example_refresh_token",
	})
}

func (h *Handler) GetUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user information successfully received",
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user data updated successfully",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user deleted successfully",
	})
}
