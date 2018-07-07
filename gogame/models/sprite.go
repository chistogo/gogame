package models

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	"github.com/chistogo/gogame/gogame/utils"
)


type SpriteModel struct {
	Pos utils.VectorI
	Height int32
	Width int32
	Parent *RenderableModel
	Sprite *sdl.Texture
	DisplayType ModelDisplayType
}




func (m *SpriteModel) GetPos( ) utils.VectorI{
	return m.Pos
}

func (m *SpriteModel) GetParent( ) *RenderableModel{
	return m.Parent
}

func (m *SpriteModel) GetDisplayType( ) ModelDisplayType{
	return 0
}

func (m *SpriteModel) VerifyDefaultValues( ){
	if(m.Pos == nil){
		m.Pos = utils.NewVectorI(0,0)
	}
}

func NewSprite(r *sdl.Renderer,imgSrc string, x, y int32 )*SpriteModel{

	texture , err := img.LoadTexture(r,imgSrc)

	if err != nil{
		// TODO: Load a default texture
		panic(err)
	}

	_,_, width, height ,err := texture.Query()

	if err != nil{
		panic(err)
	}

	return &SpriteModel{Height:height,Width:width,Sprite:texture,Pos:utils.NewVectorI(50,50)}



}


func (m *SpriteModel)Draw(renderer *sdl.Renderer)  {

	m.VerifyDefaultValues()
	rect := sdl.Rect{m.Pos[0], m.Pos[1], m.Width, m.Height}
	renderer.Copy(m.Sprite,nil,&rect)
	renderer.DrawRect(&rect)

}