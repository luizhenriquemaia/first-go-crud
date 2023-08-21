package main

import (
	"app/albums"
	"app/artists"
	"app/config"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	config.InitDB()
	app := gin.New()

	router := app.Group("/api/")
	albums.AlbumRoutes(router)
	artists.ArtistRoutes(router)

	app.Run("localhost:42069")
}
