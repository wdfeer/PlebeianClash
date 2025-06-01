package internal

import rl "github.com/gen2brain/raylib-go/raylib"

func (self GameState) Update() GameState {
	if self.AState.Tower.Hp > 0 && self.BState.Tower.Hp > 0 {
		return GameState{self.AState.update(self.BState), self.BState.update(self.AState)}
	} else {
		return self
	}
}

func (self TeamState) update(other TeamState) TeamState {
	new := self

	new.Mana += 0.01

	if self.IsLocal {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.GetMousePosition().X < 800 && self.Mana >= 1 {
			new.Units = append(new.Units, Unit{Type: Knight, Hp: 300, Position: rl.GetMousePosition()})
			new.Mana -= 1
		}
	} else {
		// TODO: spawn units
	}

	for i := range len(new.Units) {
		target := other.Tower.Position
		if rl.Vector2Distance(new.Units[i].Position, target) > 40 {
			new.Units[i].Position = rl.Vector2MoveTowards(new.Units[i].Position, target, 4)
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
