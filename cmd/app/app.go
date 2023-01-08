package app

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"go-rest-api/cmd/data"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func newAlbum(c *gin.Context) {
	var newAlbum data.Album

	//Call BindJSON to bind the received JSON to new album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add new album
	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByArtist(c *gin.Context) {
	artist := c.Param("artist")

	for _, a := range data.Albums {
		if strings.Contains(a.Artist, artist) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
