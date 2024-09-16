package youtubev2services

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/TalesPalma/internal/converters"
	"github.com/kkdai/youtube/v2"
)

// Download a playlist
func DownloadPlaylist(url string, client *youtube.Client) {
	playlist, err := client.GetPlaylist(url)
	if err != nil {
		log.Fatalf("Error with get playlist : %v", err)
	}

	for _, item := range playlist.Videos {
		video, err := client.VideoFromPlaylistEntry(item)
		if err != nil {
			log.Fatalf("Error with get video : %v", err)
		}
		SingleVideoDownload(video, client)
	}
}

// Download a single video
func SingleVideoDownload(video *youtube.Video, client *youtube.Client) {
	fmt.Println("Downloading ", video.Title, "...")

	formats := video.Formats.WithAudioChannels()
	response, _, error := client.GetStream(video, &formats[0])
	if error != nil {
		log.Fatalf("Error with get stream : %v", error)
	}
	defer response.Close()
	saveVideoMp3(video, response)

	fmt.Println("Download completed!")
}

// Save the video
func saveVideoMp3(video *youtube.Video, response io.ReadCloser) {
	fileName := video.Title + ".mp4"

	file, err := os.Create("musics/" + fileName)
	if err != nil {
		log.Fatalf("Error with create file : %v", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(response)
	if err != nil {
		log.Fatalf("Error with read file : %v", err)
	}

	converters.ConvertMp4ToMp3(fileName) // Convert the mp4 file to mp3 using ffmpeg

	time.Sleep(5 * time.Second) // wiat 5 seconds ( Prevent the YouTube server from boring me
}
