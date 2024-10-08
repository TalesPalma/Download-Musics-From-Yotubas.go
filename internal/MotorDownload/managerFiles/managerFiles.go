package managerfiles

import (
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/TalesPalma/internal/MotorDownload/converters"
	"github.com/kkdai/youtube/v2"
)

func SaveVideoMp3File(video *youtube.Video, response io.ReadCloser, fileName *string) {

	//Cria a pasta musics caso ela não exista
	err := os.MkdirAll("musics/", os.ModePerm)
	if err != nil {
		log.Fatalf("Error with create folder : %v", err)
	}

	//Sanitizar o nome do arquivo
	sanitizeFileName(fileName)

	//Caso a pasta musics exista, salva o arquivo
	file, err := os.Create("musics/" + *fileName)
	if err != nil {
		log.Fatalf("Error with create file : %v", err)
	}
	defer file.Close()

	//Escreve no arquivo
	_, err = file.ReadFrom(response)
	if err != nil {
		log.Fatalf("Error with read file : %v", err)
	}

	log.Println("Salvou com sucesso o arquivo" + *fileName)

	converters.ConvertMp4ToMp3(*fileName)
}

func sanitizeFileName(fileName *string) {
	invalidChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
	for _, chars := range invalidChars {
		*fileName = strings.ReplaceAll(*fileName, chars, "_")
	}
}

func CleanVideoMp3Folder() {
	err := filepath.WalkDir("musics/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error with clean folder : %v", err)
	}

}
