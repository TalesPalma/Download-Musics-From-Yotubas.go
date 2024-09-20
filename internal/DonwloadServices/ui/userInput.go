package ui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	managerfiles "github.com/TalesPalma/internal/DonwloadServices/managerFiles"
	youtubev2services "github.com/TalesPalma/internal/DonwloadServices/youtubev2Services"
	"github.com/kkdai/youtube/v2"
)

func LoopInterface(client *youtube.Client) {
	for {
		if UserInterface(client) {
			continue
		} else {
			break
		}
	}
}

func UserInterface(client *youtube.Client) bool {
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
	case 4:
		managerfiles.MoveFilesToFolderInHomeSystem("musics")
		fmt.Println("Moving files ...")
	default:
		fmt.Println("Close app")
		clearScreen()
		return false
	}
	clearScreen()
	return true
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
	fmt.Println("3. Clean all files from musics folder")
	fmt.Println("4. Move File to download")
	fmt.Println("Para sair pressione qualquer outra tecla")
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

func clearScreen() {
	var cmd *exec.Cmd

	// Verifica o sistema operacional
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	// Executa o comando
	cmd.Stdout = os.Stdout
	cmd.Run()
}
