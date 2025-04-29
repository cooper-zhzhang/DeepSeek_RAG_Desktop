package controller

import (
	"dp_client/controller/gui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type TestGUI struct {
}

func NewTestGUI() TestInterface {
	return &TestGUI{}
}

func (receiver *TestGUI) Run() {
	game := &gui.EbitenGame{}
	//ebiten.SetWindowSize(200, 300)
	ebiten.SetWindowTitle("Ebiten Input Box")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
