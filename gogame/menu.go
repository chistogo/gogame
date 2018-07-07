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
	player *models.Player
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
	m.player = models.NewPlayer(m.renderer, models.SPACE_SHIP)
	m.models = append(m.models, m.player )


	for{

		do(func() {
			m.PollEvents()

		})
		// Logic Runs at 2 Times the FPS
		sdl.Delay( (1000 / CurrentGameConfig.FramesPerSecond) / 2)
	}

}


func (m *Menu) PollEvents(){

	var keyState = m.previousKeyState

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		keyState = sdl.GetKeyboardState()
		m.previousKeyState = keyState
		CheckGlobalEventLister(event)
	}

	if(keyState[sdl.SCANCODE_DOWN] == 1){
		if(m.player.Pos[1]+3 > 0 && m.player.Pos[1]+3 < CurrentGameConfig.WindowHeight -32 ){
			m.player.Pos[1]+=3
		}
	}

	if(keyState[sdl.SCANCODE_UP] == 1){
		if(m.player.Pos[1]-3 > 0 && m.player.Pos[1]-3 < CurrentGameConfig.WindowHeight ){
			m.player.Pos[1]-=3
		}
	}

	if(keyState[sdl.SCANCODE_RIGHT] == 1){
		if(m.player.Pos[0]+3 > 0 && m.player.Pos[0]+3 < CurrentGameConfig.WindowWidth - 32){
			m.player.Pos[0]+=3
		}
	}

	if(keyState[sdl.SCANCODE_LEFT] == 1){
		if(m.player.Pos[0]-3 > 0 && m.player.Pos[0]-3 < CurrentGameConfig.WindowWidth){
			m.player.Pos[0]-=3
		}

	}






}
