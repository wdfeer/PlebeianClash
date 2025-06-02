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
		util.DrawTextCentered("You Lost!", 800, 400, 64, rl.White)
		util.DrawTextCentered("Press ESC to Exit or R to restart.", 800, 500, 64, rl.White)
	} else if state.BState.Tower.Hp <= 0 {
		util.DrawTextCentered("You Won!", 800, 400, 64, rl.White)
		util.DrawTextCentered("Press ESC to Exit or R to restart.", 800, 500, 64, rl.White)
	}

	if state.showControls {
		util.DrawTextCentered("Left Click: summon knight (1 Mana)", 400, 350, 40, rl.White)
	}
}

func (team TeamState) render(color rl.Color) {
	rl.DrawCircle(int32(team.Tower.Position.X), int32(team.Tower.Position.Y), 32, color)

	for _, u := range team.Units {
		rl.DrawCircle(int32(u.Position.X), int32(u.Position.Y), 20, color)
	}

	for _, p := range team.Projectiles {
		rl.DrawCircle(int32(p.Position.X), int32(p.Position.Y), 10, color)
	}

	if team.IsLocal {
		text := strconv.FormatFloat(float64(team.Mana), 'f', 1, 32) + " Mana"
		util.DrawTextCentered(text, 800, 850, 40, rl.SkyBlue)
	}
}
