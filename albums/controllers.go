package albums

import (
	"app/config"
	"app/utils"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	repo := AbumRepository(config.GetDB())
	albums, err := repo.All()
	if err != nil {
		c.IndentedJSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"data": albums})
}

func getAlbum(c *gin.Context) {
	albumId := utils.ParseParamId(c.Param("id"))
	repo := AbumRepository(config.GetDB())
	album, err := repo.GetByID(albumId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"data": album})
}

func postAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	repo := AbumRepository(config.GetDB())
	album, err := repo.Create(newAlbum)
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": album})
}

func putAlbum(c *gin.Context) {
	albumId := utils.ParseParamId(c.Param("id"))
	var newDataAlbum Album
	if err := c.BindJSON(&newDataAlbum); err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}

	repo := AbumRepository(config.GetDB())
	album, err := repo.Update(albumId, newDataAlbum)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	album.ID = albumId
	c.IndentedJSON(200, gin.H{"data": album})
}

func deleteAlbum(c *gin.Context) {
	albumId := utils.ParseParamId(c.Param("id"))
	repo := AbumRepository(config.GetDB())
	err := repo.Delete(albumId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(204, gin.H{"message": "Album not founded"})
}
