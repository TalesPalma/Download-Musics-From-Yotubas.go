package main

import (
	"github.com/TalesPalma/internal/DonwloadServices/ui"
	"github.com/kkdai/youtube/v2"
)

// "github.com/TalesPalma/internal/api"

func main() {
	// api.Handler()
	ui.LoopInterface(&youtube.Client{})
}
