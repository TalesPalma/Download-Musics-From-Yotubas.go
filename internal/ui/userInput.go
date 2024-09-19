package ui

import (
	"fmt"
	"log"
	"strings"

	managerfiles "github.com/TalesPalma/internal/managerFiles"
	youtubev2services "github.com/TalesPalma/internal/youtubev2Services"
	"github.com/kkdai/youtube/v2"
)

func UserInterface(client *youtube.Client) {

	userJoice()
	options := InputUserOption()
	switch options {
	case 1:
		fmt.Println("Digite sua url:")
		handleDownloadSingleVideo(client)
	case 2:
		fmt.Println("Digite sua url:")
		handleDownloadPlaylist(client)
	case 3:
		handleCleanFiles()
		fmt.Println("Removing files ...")
	default:
		fmt.Println("Invalid option")
	}
}

func handleCleanFiles() {
	managerfiles.CleanVideoMp3Folder()
}

func handleDownloadSingleVideo(client *youtube.Client) {
	video, err := client.GetVideo(InputUserUrl())
	if err != nil {
		log.Fatalf("Error with get video : %v", err)
	}
	youtubev2services.SingleVideoDownload(video, client)
}

func handleDownloadPlaylist(client *youtube.Client) {
	youtubev2services.DownloadPlaylist(InputUserUrl(), client)
}

func userJoice() {
	fmt.Println("1. Download single video")
	fmt.Println("2. Download playlist")
	fmt.Println("3.Clean all files from musics folder")
}

func InputUserOption() int {
	var option int
	fmt.Scanf("%d", &option)
	return option
}

func InputUserUrl() string {
	var url string
	_, erro := fmt.Scanf("%s", &url)

	if erro != nil {
		log.Fatalf("Error with read url : %v", erro)
	}

	if url == "" {
		log.Fatalf("Url not informed")
	}

	return strings.TrimSpace(url)
}
