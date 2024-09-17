package main

import (
	"github.com/TalesPalma/internal/ui"
	"github.com/kkdai/youtube/v2"
)

// Hello world
func main() {
	ui.UserInterface(&youtube.Client{})

}
