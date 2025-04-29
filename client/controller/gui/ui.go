package gui

import "github.com/hajimehoshi/ebiten/v2"

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (receiver *UI) Run() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
