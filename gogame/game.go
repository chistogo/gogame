package gogame

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)


var CurrentGameConfig *GameConfig
var window *sdl.Window
var previousKeyState []uint8
var CurrentScene Scene
var RunGame bool
var mainfunc = make(chan func())


type Scene interface {
	LogicLoop()
	DrawLoop()
	PollEvents()
}

func CheckGlobalEventLister(event sdl.Event){
	var keyState []uint8

	switch /* t := */ event.(type) {
	case *sdl.QuitEvent:
		RunGame = false
	}

	keyState = sdl.GetKeyboardState()
	if(previousKeyState != nil) {


		if(keyState[sdl.SCANCODE_ESCAPE] == 1){
			RunGame = false
		}

		if (keyState[sdl.SCANCODE_F11] == 1){


			if(CurrentGameConfig.FullScreen){
				fmt.Println("EXIT FULL SCREEN")
				window.SetFullscreen(0)
			}else{
				fmt.Println("FULL SCREEN")
				fmt.Println(window)
				window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
			}
			CurrentGameConfig.FullScreen = !CurrentGameConfig.FullScreen


		}
	}
	previousKeyState = keyState
	
	
}



type GameConfig struct {
	Title string
	WindowWidth int32
	WindowHeight int32
	FullScreen bool
	FramesPerSecond uint32
}

func Run(conf *GameConfig) int{
	fmt.Println("Work")

	CurrentGameConfig = conf

	// Setup the Render Loop
	var renderer *sdl.Renderer

	w, err := sdl.CreateWindow(conf.Title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		conf.WindowWidth, conf.WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer w.Destroy()
	window = w

	renderer, err = sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	renderer.Clear()

	RunGame = true
	CurrentScene := Menu{renderer:renderer}

	if(conf.FullScreen){
		window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	}

	go CurrentScene.LogicLoop()
	go CurrentScene.DrawLoop()

	for f := range mainfunc {
		f()
		if !RunGame {
			return 0
		}
	}

	return 0
}


/*
	https://github.com/golang/go/wiki/LockOSThread
	Queue of work to run in main thread.
	Takes a function that you want to run and makes sure it runs in the main thread.
 */

func do(f func()) {

	done := make(chan bool, 1)

	// Wrap Function so we get notified through channel when the function exits
	mainfunc <- func() {
		f()
		done <- true
	}
	// Will Block until done is received.
	<-done
}

