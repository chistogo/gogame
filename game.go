package main

import (
	"fmt"
	"github.com/chistogo/gogame/gogame"
	"os"
)


func main() {
	fmt.Println("Go Game")

	// Todo: Read from file or args for config

	config := &gogame.GameConfig{
		Title: "Go Game Alpha",
		WindowHeight: 720,
		WindowWidth: 1280,
		FramesPerSecond: 60,
		FullScreen:false,
	}

	os.Exit( gogame.Run(config) )

}