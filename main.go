package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {
	url := "https://www.youtube.com/watch?v=HVgRjMGRtsg&list=PLdKWRMG0YzNsTedetJAYPDkblxzY014ZS"
	download(url)
}

func download(url string) {

	client := youtube.Client{}

	video, err := client.GetVideo(url)

	if err != nil {
		log.Fatalf("Error with get video : %v", err)
	}

	formats := video.Formats.WithAudioChannels()

	response, _, error := client.GetStream(video, &formats[0])

	if error != nil {
		log.Fatalf("Error with get stream : %v", error)
	}
	defer response.Close()

	fileName := "video.mp3"
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Error with create file : %v", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(response)
	if err != nil {
		log.Fatalf("Error with read file : %v", err)
	}

	fmt.Println("Deu certo acho")

}

func inputUserUrl() string {
	var url string
	_, erro := fmt.Scanf("%s", &url)

	if erro != nil {
		log.Fatalf("Error with read url : %v", erro)
		return ""
	}

	return strings.TrimSpace(url)
}
