package routers

import (
	controllers "github.com/AbdurrahmanTalha/brainscape-backend-go/api/controllers"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	controller := controllers.NewUserController(cfg)

	router.GET("/", controller.GetAllUsers)

	router.GET("/:id", controller.GetSpecificUser)

	router.PATCH("/:id", controller.UpdateSpecificUser)

	router.POST("/register", controller.Register)

	router.POST("/login", controller.Login)
}
