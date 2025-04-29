package gui

import (
	"flag"
	_ "image/png"
	"os"

	"golang.org/x/text/language"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	lang       language.Tag
	cpup       *os.File
	inputText  string
	outputText string
	buttonRect [4]int // x, y, width, height
}

var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	mute       = flag.Bool("mute", false, "mute")
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}
