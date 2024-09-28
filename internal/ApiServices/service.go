package api

import (
	"net/http"

	youtubev2services "github.com/TalesPalma/internal/MotorDownload/youtubev2Services"
	"github.com/TalesPalma/internal/models"
	"github.com/gin-gonic/gin"
)

var ListMusics []models.Music //Vai add musicas a serem baixadas

func Handler(r *gin.Engine) {
	r.GET("/musics", GetMusicsInDonwload)
	r.POST("/download", PostDownloadPlaylist)
}

func GetMusicsInDonwload(c *gin.Context) {
	c.JSON(http.StatusOK, ListMusics)
}

func PostDownloadPlaylist(c *gin.Context) {
	var linkPlaylist models.PlaylistLink
	if err := c.ShouldBindJSON(&linkPlaylist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	youtubev2services.DownloadPlaylist(linkPlaylist.Link, youtubev2services.GetClient(), &ListMusics)
	c.JSON(http.StatusOK, gin.H{"Message": "Download inicado para " + linkPlaylist.Link})
}
