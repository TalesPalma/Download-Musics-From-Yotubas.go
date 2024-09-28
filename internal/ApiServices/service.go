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
	r.GET("/download/:musicName", getDownloadMusic) //A gente manda o title da musica para o get e pronto já vai puchar o donwload
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
	go youtubev2services.DownloadPlaylist(linkPlaylist.Link, youtubev2services.GetClient(), &ListMusics)
	c.JSON(http.StatusOK, gin.H{"Message": "Download inicado para " + linkPlaylist.Link})
}

func getDownloadMusic(c *gin.Context) {
	link := c.Param("musicName")
	music := link + ".mp3"
	filePath := "musics/" + music

	c.Header("Content-Disposition", "attachment; filename="+music)
	c.Header("Content-Type", "audio/mpeg")
	c.File(filePath)

}
