package tracks

import "github.com/gin-gonic/gin"

func TrackRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/track")
	{
		router.GET("", getTracks)
		router.POST("/", postTrack)
	}
}
