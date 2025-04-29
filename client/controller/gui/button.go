package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	buttonX      = 250
	buttonY      = 200
	buttonWidth  = 140
	buttonHeight = 50
)

type ButtonField struct {
	buttonColor color.Color
	isPressed   bool
}

func (receiver *ButtonField) Update() error {
	// 检测鼠标位置和按钮是否被点击
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
			receiver.isPressed = true
			receiver.buttonColor = color.RGBA{0x00, 0xff, 0x00, 0xff} // 绿色
		} else {
			receiver.isPressed = false
			receiver.buttonColor = color.RGBA{0xff, 0x00, 0x00, 0xff} // 红色
		}
	} else {
		receiver.isPressed = false
		receiver.buttonColor = color.RGBA{0xff, 0x00, 0x00, 0xff} // 红色
	}
	return nil
}

func (receiver *ButtonField) Draw(screen *ebiten.Image) {

	/*	// 绘制圆角按钮
		drawRoundedRect(screen, buttonX, buttonY, buttonWidth, buttonHeight, cornerRadius, g.buttonColor)
		// 绘制文字
		ebitenutil.DebugPrintAt(screen, "Click Me!", buttonX+10, buttonY+10)*/

	// 绘制按钮
	ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), float64(buttonWidth), float64(buttonHeight), receiver.buttonColor)
	ebitenutil.DrawCircle(screen, float64(buttonX+150), float64(buttonY+150), float64(buttonHeight), receiver.buttonColor)
	// 绘制文字
	ebitenutil.DebugPrintAt(screen, "Click Me!", buttonX+10, buttonY+10)
}
