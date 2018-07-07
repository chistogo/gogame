package models

import "github.com/veandco/go-sdl2/sdl"

type RenderableModel interface{
	Draw(renderer *sdl.Renderer)
}

