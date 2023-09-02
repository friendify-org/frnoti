package routes

import (
	"main/config"
	"main/middlewares"
	"main/models"
	"main/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	route gin.RouterGroup
}

func NewUserController(route *gin.RouterGroup) {

	config.Db.AutoMigrate(&models.User{})

	userService := services.UserService{}

	route = route.Group("/users")
	{
		route.POST("/register", userService.Register)

		route.POST("/login", userService.Login)

		route.GET("/profile", middlewares.Authorization(), userService.GetProfile)

		route.PUT("/active", userService.ActiveUser)

		route.PUT("/update_password", middlewares.Authorization(), userService.UpdatePassword)
	}
}
