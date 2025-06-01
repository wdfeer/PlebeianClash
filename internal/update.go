package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (self GameState) Update() GameState {
	return GameState{self.AState.update(self), self.BState.update(self)}
}

func (self TeamState) update(game GameState) TeamState {
	new := self

	if self.IsLocal && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		new.Units = append(new.Units, Unit{Type: Knight, Hp: 300, Position: rl.GetMousePosition()})
	}

	var target rl.Vector2
	if self.IsLocal == game.AState.IsLocal {
		target = game.BState.Tower.Position
	} else {
		target = game.AState.Tower.Position
	}

	for i := range len(new.Units) {
		new.Units[i].Position = rl.Vector2MoveTowards(new.Units[i].Position, target, 4)
	}

	// TODO: update tower hp

	if self.Tower.Hp <= 0 {
		rl.CloseWindow()
	}

	return new
}
