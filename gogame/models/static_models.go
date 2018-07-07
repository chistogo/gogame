package models

import "github.com/veandco/go-sdl2/sdl"


type ModelDisplayType int

const (
	STATIC ModelDisplayType = iota
	RELATIVE
)

type StaticModel struct {
	PosX int32
	PosY int32
	Height int32
	Weight int32
	Sprite *sdl.Texture
	DisplayType ModelDisplayType
}

func (m *StaticModel)Draw(renderer *sdl.Renderer)  {

	rect := sdl.Rect{m.PosX, m.PosY, 200, 200}
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

}