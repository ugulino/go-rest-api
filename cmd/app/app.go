package app

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-rest-api/cmd/data"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}
