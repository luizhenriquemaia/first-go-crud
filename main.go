package main

import (
	"app/albums"
	"app/artists"
	"app/config"
	"app/tracks"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.InitDB()
	app := gin.New()

	router := app.Group("/api/")
	albums.AlbumRoutes(router)
	artists.ArtistRoutes(router)
	tracks.TrackRoutes(router)

	app.Run("localhost:42069")
}
