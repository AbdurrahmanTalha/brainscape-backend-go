package api

import (
	"fmt"

	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/middlewares"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/api/routers"
	"github.com/AbdurrahmanTalha/brainscape-backend-go/config"
	"github.com/gin-gonic/gin"
)

func SetupServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(middlewares.Cors(cfg))

	RegisterRoutes(r, cfg)

	err := r.Run(":" + cfg.Server.Port)
	if err != nil {
		// panic("[ERROR] Failed to start server")
		fmt.Println(err)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		user := v1.Group("/user")
		
		routers.User(user, cfg)
	}
}
