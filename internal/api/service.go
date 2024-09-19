package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func Handler() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/download/", fileHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("Listening on prot 8080")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := path.Base(r.URL.Path)
	filePath := "musics/" + path

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filePath)
	http.ServeFile(w, r, filePath)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Home Page<h1>")
}
