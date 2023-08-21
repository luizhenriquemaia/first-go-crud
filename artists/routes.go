package artists

import "github.com/gin-gonic/gin"

func ArtistRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/artist")
	{
		router.GET("", getArtists)
		router.POST("/", postArtist)
	}
}
