package app

import "github.com/gin-gonic/gin"

func Router() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", newAlbum)
	router.GET("/albums/:artist", getAlbumsByArtist)
	router.Run("localhost:8080")
}
