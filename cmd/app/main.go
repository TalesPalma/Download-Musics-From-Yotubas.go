package main

import (
	// "github.com/TalesPalma/internal/api"
	"github.com/TalesPalma/internal/ui"
	"github.com/kkdai/youtube/v2"
)

// Hello world
func main() {
	// api.Handler()
	ui.UserInterface(&youtube.Client{})
}
