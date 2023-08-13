package albums

import (
	"github.com/gin-gonic/gin"
)

func AlbumRoutes(superRouter *gin.RouterGroup) {
	router := superRouter.Group("/album")
	{

		router.GET("", getAlbums)
		router.GET("/:id", getAlbum)
		router.POST("/", postAlbum)
		router.PUT("/:id/", putAlbum)
		router.DELETE("/:id/", deleteAlbum)
	}
}
