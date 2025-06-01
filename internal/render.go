package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (state GameState) Render() {
	rl.ClearBackground(rl.Black)

	if state.AState.Tower.Hp > 0 && state.BState.Tower.Hp > 0 {
		rl.DrawLine(800, 0, 800, 900, rl.DarkGray)
	}

	state.AState.render(rl.DarkGreen)
	state.BState.render(rl.DarkPurple)

	if state.AState.Tower.Hp < 0 {
		text := "You Lost! Press ESC to Exit."
		var textSize int32 = 64
		textWidth := rl.MeasureText(text, textSize)
		rl.DrawText(text, 800-(textWidth/2), 450, textSize, rl.White)
	} else if state.BState.Tower.Hp < 0 {
		text := "You Won! Press ESC to Exit."
		var textSize int32 = 64
		textWidth := rl.MeasureText(text, textSize)
		rl.DrawText(text, 800-(textWidth/2), 450, textSize, rl.White)
	}
}

func (team TeamState) render(color rl.Color) {
	rl.DrawCircle(int32(team.Tower.Position.X), int32(team.Tower.Position.Y), 32, color)
	for _, u := range team.Units {
		rl.DrawCircle(int32(u.Position.X), int32(u.Position.Y), 20, color)
	}
}
