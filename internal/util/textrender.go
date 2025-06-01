package util

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawTextCentered(text string, x int32, y int32, fontSize int32, color rl.Color) {
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text, x-(textWidth/2), y-(fontSize/2), fontSize, color)
}
