package main

import (
	"PlebeianClash/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "Plebeian Clash")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	state := internal.DefaultState
	for !rl.WindowShouldClose() {
		state = state.Update()

		rl.BeginDrawing()
		state.Render()
		rl.EndDrawing()
	}
}
