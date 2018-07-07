package gogame

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/chistogo/gogame/gogame/models"
)

type MenuOption struct {
	text string

}

type Menu struct {
	options []MenuOption
	renderer *sdl.Renderer
	previousKeyState []uint8
	models []models.RenderableModel
}



func (m *Menu) DrawLoop(){
	for{
		do(func() {

			m.renderer.SetDrawColor(0, 0, 0, 255)
			m.renderer.Clear()
			for _, model := range (m.models)  {
				model.Draw(m.renderer)
			}

			m.renderer.Present()


		})

		sdl.Delay(1000 / CurrentGameConfig.FramesPerSecond)
	}

}

func (m *Menu) LogicLoop(){



	// Load Models
	m.models = append(m.models, &models.StarBackgroundModel{SkipFrame:2,StartCount:100, Height:CurrentGameConfig.WindowHeight,Width:CurrentGameConfig.WindowWidth})

	for{

		do(func() {
			m.PollEvents()



		})
		// Logic Runs at 2 Times the FPS
		sdl.Delay(2* (1000 / CurrentGameConfig.FramesPerSecond))
	}

}


func (m *Menu) PollEvents(){

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		keyState := sdl.GetKeyboardState()

		if(keyState[sdl.SCANCODE_DOWN] == 1){
			m.models = m.models[:len(m.models)-1]
		}

		if(keyState[sdl.SCANCODE_UP] == 1){
				m.models = append(m.models, &models.StarBackgroundModel{SkipFrame:2,StartCount:100, Height:CurrentGameConfig.WindowHeight,Width:CurrentGameConfig.WindowWidth})
		}



		m.previousKeyState = keyState
		CheckGlobalEventLister(event)
	}

}
