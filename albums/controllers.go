package albums

import (
	"app/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseParamId(id string) int64 {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	return idInt
}

func getAlbums(c *gin.Context) {
	repo := AbumRepository(config.GetDB())
	albums, err := repo.All()
	if err != nil {
		c.IndentedJSON(500, err)
	}
	c.IndentedJSON(200, albums)
}

func getAlbum(c *gin.Context) {
	albumId := parseParamId(c.Param("id"))
	repo := AbumRepository(config.GetDB())
	album, err := repo.GetByID(albumId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(200, album)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(400, err)
		return
	}
	repo := AbumRepository(config.GetDB())
	album, err := repo.Create(newAlbum)
	if err != nil {
		c.IndentedJSON(400, err)
	}
	c.IndentedJSON(201, album)
}

func putAlbum(c *gin.Context) {
	albumId := parseParamId(c.Param("id"))
	var newDataAlbum Album
	if err := c.BindJSON(&newDataAlbum); err != nil {
		c.IndentedJSON(400, err)
		return
	}

	repo := AbumRepository(config.GetDB())
	album, err := repo.Update(albumId, newDataAlbum)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	album.ID = albumId
	c.IndentedJSON(200, album)
}

func deleteAlbum(c *gin.Context) {
	albumId := parseParamId(c.Param("id"))
	repo := AbumRepository(config.GetDB())
	err := repo.Delete(albumId)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(204, gin.H{"message": "Album not founded"})
}
