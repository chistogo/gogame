package models

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"github.com/chistogo/gogame/gogame/utils"
)

type Star struct {
	PosX float32
	PosY float32
	VelX float32
	VelY float32
}


type StarBackgroundModel struct {
	Height int32
	Width int32
	Pos utils.VectorI
	Parent *RenderableModel
	StartCount int
	points []sdl.Point
	SkipFrame int32
	skipFrameCount int32
}


func (m *StarBackgroundModel) GetPos()utils.VectorI{
	return m.Pos
}

func (m *StarBackgroundModel) GetDisplayType() ModelDisplayType{
	return 0
}

func (m *StarBackgroundModel) GetParent() *RenderableModel {
	return m.Parent
}

func (m *StarBackgroundModel)Draw(renderer *sdl.Renderer)  {

	// Check to see if Sensible Defaults were set

	if(m.StartCount == 0){
		m.StartCount = 100
	}

	if(m.points == nil){
		m.points = []sdl.Point{}
	}



	if m.skipFrameCount == 0 {

		for i:= 0; i < m.StartCount; i++{

			// This Index is Bigger than the length of the points and thus you need to grow the slice
			if(i >= len(m.points)){
				point := &sdl.Point{rand.Int31n(m.Width/5) + 2*m.Width/5,rand.Int31n(m.Height/5) + 2*m.Height/5}
				m.points = append(m.points,*point)
			}

			// Modify Star
			m.points[i].X +=  int32(float32(m.points[i].X - m.Width/2)/10)+ rand.Int31n(2)
			m.points[i].Y +=  int32(float32(m.points[i].Y - m.Height/2)/10)+ rand.Int31n(2)

			// Check if its out of points and reset the point to middle
			if(m.points[i].Y > m.Height || m.points[i].Y < 0 || m.points[i].X > m.Width || m.points[i].X < 0){
				m.points[i].X = rand.Int31n(m.Width/9) + 4*m.Width/9
				m.points[i].Y = rand.Int31n(m.Height/9) + 4*m.Height/9
			}


		}
		m.skipFrameCount = m.SkipFrame
	}else{
		m.skipFrameCount--
	}


	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoints(m.points)

}