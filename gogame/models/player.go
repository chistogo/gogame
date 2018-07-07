package models

import "github.com/veandco/go-sdl2/sdl"

type Player struct {
	PosX float32
	PosY float32
	VelX float32
	VelY float32
}

func (p *Player)Draw(renderer *sdl.Renderer)  {



}