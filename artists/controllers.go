package artists

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

func getArtists(c *gin.Context) {
	repo := ArtistRepository(config.GetDB())
	artists, err := repo.All()
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"data": artists})
}

func postArtist(c *gin.Context) {
	var newArtist Artist
	if err := c.BindJSON(&newArtist); err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	repo := ArtistRepository(config.GetDB())
	artist, err := repo.Create(newArtist)
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": artist})
}
