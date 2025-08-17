package routes

import (
	"auth/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.RegisterUser)
		auth.POST("/login", h.LoginUser)

		auth.GET("/info", h.GetUserInfo)

		auth.PUT("/update", h.UpdateUser)

		auth.DELETE("/delete", h.DeleteUser)
	}

	return router
}
