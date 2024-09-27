package main

import (
	api "github.com/TalesPalma/internal/ApiServices"
	// "github.com/TalesPalma/internal/DonwloadServices/ui"
	"github.com/gin-gonic/gin"
	// "github.com/kkdai/youtube/v2"
)

// "github.com/TalesPalma/internal/api"

func main() {
	runApi()
	// ui.LoopInterface(&youtube.Client{})
}

func runApi() {
	r := gin.Default()
	api.Handler(r)
	r.Run(":8080")
}
