package main

import (
	"github.com/TalesPalma/src/ui"
	"github.com/kkdai/youtube/v2"
)

func main() {
	ui.UserInterface(&youtube.Client{})
}
