package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (self GameState) Update() GameState {
	return GameState{self.AState.update(), self.BState.update()}
}

func (self TeamState) update() TeamState {
	new := self

	if self.IsLocal && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		new.Units = append(new.Units, Unit{Type: Knight, Hp: 300, Position: rl.GetMousePosition()})
	}

	// TODO: update unit positions, tower hp and stuff

	if self.Tower.Hp <= 0 {
		rl.CloseWindow()
	}

	return new
}
