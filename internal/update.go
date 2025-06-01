package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (self GameState) Update() GameState {
	return GameState{self.AState.update(self.BState), self.BState.update(self.AState)}
}

func (self TeamState) update(other TeamState) TeamState {
	new := self

	if self.IsLocal && rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.GetMousePosition().X < 800 {
		new.Units = append(new.Units, Unit{Type: Knight, Hp: 300, Position: rl.GetMousePosition()})
	}

	for i := range len(new.Units) {
		if rl.Vector2Distance(new.Units[i].Position, other.Tower.Position) > 40 {
			new.Units[i].Position = rl.Vector2MoveTowards(new.Units[i].Position, other.Tower.Position, 4)
		}
	}

	for _, u := range other.Units {
		if rl.Vector2Distance(u.Position, self.Tower.Position) < 50 {
			new.Tower.Hp -= 10
		}
	}

	if self.Tower.Hp <= 0 {
		rl.CloseWindow()
	}

	return new
}
