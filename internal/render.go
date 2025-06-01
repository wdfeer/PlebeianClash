package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state GameState) Render() {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.White)
}
