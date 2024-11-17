package routes

import (
	"RyuLdnWebsite/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/", controllers.GetLDNData)
		api.GET("/public_games", controllers.GetPublicGames)
		api.GET("/status", controllers.GetLdnStatus)
	}
}
