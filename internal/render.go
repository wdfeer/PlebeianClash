package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state GameState) Render() {
	rl.ClearBackground(rl.Black)

	rl.DrawLine(800, 0, 800, 900, rl.DarkGray)

	state.AState.render(rl.DarkGreen)
	state.BState.render(rl.DarkPurple)
}

func (team TeamState) render(color rl.Color) {
	rl.DrawCircle(int32(team.Tower.Position.X), int32(team.Tower.Position.Y), 32, color)
	for _, u := range team.Units {
		rl.DrawCircle(int32(u.Position.X), int32(u.Position.Y), 20, color)
	}
}
