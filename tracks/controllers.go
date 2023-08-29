package tracks

import (
	"app/config"

	"github.com/gin-gonic/gin"
)

func getTracks(c *gin.Context) {
	repo := TrackRepository(config.GetDB())
	tracks, err := repo.All()
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"data": tracks})
}

func postTrack(c *gin.Context) {
	var newTrack Track
	if err := c.Bind(&newTrack); err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	repo := TrackRepository(config.GetDB())
	track, err := repo.Create(newTrack)
	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(201, gin.H{"data": track})
}
