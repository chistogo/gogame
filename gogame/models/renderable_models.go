package models

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/chistogo/gogame/gogame/utils"
)


type ModelDisplayType int

const (
	STATIC ModelDisplayType = iota
	RELATIVE
)



type RenderableModel interface{
	Draw(renderer *sdl.Renderer)
	GetDisplayType() ModelDisplayType
	GetPos() utils.VectorI
	GetParent() *RenderableModel
}

/*
Will return the coordinates relative to the parent.
 */
func GetPos(m RenderableModel){
	if(m.GetDisplayType() == RELATIVE){

	}


}