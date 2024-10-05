package handler

import (
	"log"
	"net/http"

	managerfiles "github.com/TalesPalma/internal/MotorDownload/managerFiles"
	youtubev2services "github.com/TalesPalma/internal/MotorDownload/youtubev2Services"
	"github.com/TalesPalma/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	ListMusics []models.Music //Vai add musicas a serem baixadas
	upgrader   = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	Broadcast = make(chan []models.Music) //Vai add musicas a serem baixadas
)

func Handler(r *gin.Engine) {
	r.GET("/musics", GetMusicsInDonwload)
	r.POST("/download", PostDownloadPlaylist)
	r.GET("/download/:musicName", getDownloadMusic) //A gente manda o title da musica para o get e pronto já vai puchar o donwload
	r.GET("/ws", wsHandle)
}

func wsHandle(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()
	clients[conn] = true
	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", string(message))

		if string(message) == "clear" {
			ListMusics = []models.Music{}
			Broadcast = make(chan []models.Music)
			managerfiles.CleanVideoMp3Folder()
			log.Println("Limpou a pasta")
		}

		select {
		case newMusics := <-Broadcast: // Escuta o canal de broadcasta ou seja quando algo for atualizado na lista
			//Envia a nova lista de musicas para o client
			if err := conn.WriteJSON(newMusics); err != nil {
				delete(clients, conn)
				log.Fatal(err)
				break
			}
		}
	}
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

	// Baixar musicas
	youtubev2services.DownloadPlaylist(linkPlaylist.Link, youtubev2services.GetClient(), &ListMusics, &Broadcast)
	c.JSON(http.StatusOK, gin.H{"Message": "Download iniciado para " + linkPlaylist.Link})
}

func getDownloadMusic(c *gin.Context) {
	link := c.Param("musicName")
	music := link + ".mp3"
	filePath := "musics/" + music

	c.Header("Content-Disposition", "attachment; filename="+music)
	c.Header("Content-Type", "audio/mpeg")
	c.File(filePath)

}
