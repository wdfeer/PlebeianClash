package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
			unit := Unit{Type: Knight, Hp: 300, Position: rl.GetMousePosition()}
			new.Units = append(new.Units, unit)
			new.Mana -= 1
		}
	} else if self.Mana > 1 {
		direction := rl.Vector2Normalize(rl.Vector2Subtract(other.Tower.Position, self.Tower.Position))
		position := rl.Vector2Add(self.Tower.Position, rl.Vector2Scale(rl.Vector2Rotate(direction, (rand.Float32()-0.5)), 500))
		unit := Unit{Type: Knight, Hp: 300, Position: position}

		new.Units = append(new.Units, unit)
		new.Mana -= 1
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
