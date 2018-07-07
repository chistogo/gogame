package models

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/chistogo/gogame/gogame/utils"
	"github.com/veandco/go-sdl2/img"
)


type PlayerType int

const (
	SPACE_SHIP PlayerType = iota
)

const (
	SPACE_SHIP_SPRITE  = "./assets/space_ship.png"
)



type Player struct {
	Parent *RenderableModel
	DisplayType ModelDisplayType
	Sprite *sdl.Texture
	Height int32
	Width int32
	Pos utils.VectorI
	Vel utils.VectorI
}

func NewPlayer(r *sdl.Renderer, playerType PlayerType) *Player {

	var err error
	var texture *sdl.Texture

	if(playerType == SPACE_SHIP){
		texture , err = img.LoadTexture(r,SPACE_SHIP_SPRITE)
	}

	if err != nil{
		panic(err)
	}

	return &Player{Height:32,Width:32,Sprite:texture,Pos:utils.NewVectorI(0,0),Vel:utils.NewVectorI(0,0)}

}




func (p *Player) GetPos() utils.VectorI{
	return p.Pos
}

func (p *Player) GetParent() *RenderableModel{
	return p.Parent
}

func (p *Player) GetDisplayType() ModelDisplayType{
	return p.DisplayType
}

func (p *Player)Draw(renderer *sdl.Renderer)  {

	renderer.SetDrawColor(0, 0, 0, 255)
	rect := sdl.Rect{p.Pos[0], p.Pos[1], p.Width, p.Height}
	renderer.Copy(p.Sprite,nil,&rect)
	renderer.DrawRect(&rect)
}