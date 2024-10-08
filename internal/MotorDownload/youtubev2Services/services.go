package youtubev2services

import (
	"fmt"
	"io"
	"log"

	managerfiles "github.com/TalesPalma/internal/MotorDownload/managerFiles"
	"github.com/TalesPalma/internal/models"
	"github.com/kkdai/youtube/v2"
)

func GetClient() *youtube.Client {
	return &youtube.Client{}
}

// Download a playlist
func DownloadPlaylist(
	url string,
	client *youtube.Client,
	listMusics *[]models.Music,
	broadcast *chan []models.Music,
) {
	playlist, err := client.GetPlaylist(url)
	if err != nil {
		log.Fatalf("Error with get playlist : %v", err)
	}

	for _, item := range playlist.Videos {
		video, err := client.VideoFromPlaylistEntry(item)
		if err != nil {
			log.Fatalf("Error with get video : %v", err)
		}

		var fileName string
		SingleVideoDownload(video, client, &fileName)

		*listMusics = append(*listMusics, models.Music{
			Title: fileName,
		})

		*broadcast <- *listMusics // Isso quebrour a concorrencia e fez travar o donwload em 1 musicas

	}
}

// Download a single video
func SingleVideoDownload(video *youtube.Video, client *youtube.Client, fileName *string) {
	fmt.Println("Downloading ", video.Title, "...")

	formats := video.Formats.WithAudioChannels()
	response, _, error := client.GetStream(video, &formats[0])
	if error != nil {
		log.Fatalf("Error with get stream : %v", error)
	}
	defer response.Close()

	saveVideoMp3(video, response, fileName)

	fmt.Println("Download completed!")
}

// Save the video
func saveVideoMp3(video *youtube.Video, response io.ReadCloser, fileName *string) {
	*fileName = video.Title + ".mp4"                         // Fica de zoio
	managerfiles.SaveVideoMp3File(video, response, fileName) // Save the mp4 file

}
