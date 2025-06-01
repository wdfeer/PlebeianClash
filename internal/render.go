package internal

import (
	"PlebeianClash/internal/util"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (state GameState) Render() {
	rl.ClearBackground(rl.Black)

	if state.AState.Tower.Hp > 0 && state.BState.Tower.Hp > 0 {
		rl.DrawLine(800, 0, 800, 900, rl.DarkGray)
	}

	state.AState.render(rl.DarkGreen)
	state.BState.render(rl.DarkPurple)

	if state.AState.Tower.Hp <= 0 {
		util.DrawTextCentered("You Lost! Press ESC to Exit.", 800, 450, 64, rl.White)
	} else if state.BState.Tower.Hp <= 0 {
		util.DrawTextCentered("You Won! Press ESC to Exit.", 800, 450, 64, rl.White)
	}
}

func (team TeamState) render(color rl.Color) {
	rl.DrawCircle(int32(team.Tower.Position.X), int32(team.Tower.Position.Y), 32, color)
	for _, u := range team.Units {
		rl.DrawCircle(int32(u.Position.X), int32(u.Position.Y), 20, color)
	}

	if team.IsLocal {
		text := strconv.FormatFloat(float64(team.Mana), 'f', 1, 32) + " Mana"
		util.DrawTextCentered(text, 800, 850, 40, rl.SkyBlue)
	}
}
