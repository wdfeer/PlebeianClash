package internal

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (self GameState) Update() GameState {
	if self.AState.Tower.Hp > 0 && self.BState.Tower.Hp > 0 {
		if len(self.AState.Units) > 0 {
			self.showControls = false
		}
		return GameState{self.AState.update(self.BState), self.BState.update(self.AState), self.showControls}
	} else if rl.IsKeyPressed(rl.KeyR) {
		state := DefaultState
		state.showControls = false
		return state
	} else {
		return self
	}
}

func (self TeamState) update(other TeamState) TeamState {
	new := self

	new.Mana += 0.01

	if self.IsLocal && rl.GetMousePosition().X < 800 {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && new.Mana >= 1 {
			unit := Unit{Type: Melee, Hp: 300, Position: rl.GetMousePosition()}
			new.Units = append(new.Units, unit)
			new.Mana -= 1
		}
		if rl.IsMouseButtonPressed(rl.MouseRightButton) && new.Mana >= 2 {
			unit := Unit{Type: Ranged, Hp: 200, Position: rl.GetMousePosition()}
			new.Units = append(new.Units, unit)
			new.Mana -= 2
		}
	} else if self.Mana > 1 {
		direction := rl.Vector2Normalize(rl.Vector2Subtract(other.Tower.Position, self.Tower.Position))
		position := rl.Vector2Add(self.Tower.Position, rl.Vector2Scale(rl.Vector2Rotate(direction, (rand.Float32()-0.5)), 500))
		unit := Unit{Type: Melee, Hp: 300, Position: position}

		new.Units = append(new.Units, unit)
		new.Mana -= 1
	}

	new = new.updateUnits(other)
	for i := range len(new.Projectiles) {
		new.Projectiles[i].Position = rl.Vector2Add(new.Projectiles[i].Position, new.Projectiles[i].Velocity)
	}

	return new
}

func (self TeamState) updateUnits(other TeamState) TeamState {
	new := self

	const attackDamage = 10

	for i := range len(new.Units) {
		// Move to target
		var attackRange float32
		if new.Units[i].Type == Melee {
			attackRange = 50
		} else {
			attackRange = 250
		}
		target := other.Tower.Position
		towerDistance := rl.Vector2Distance(new.Units[i].Position, target)
		targetDistance := towerDistance
		unitIndex, unitDistance := new.Units[i].closestEnemy(other)
		if unitDistance < towerDistance {
			target = other.Units[unitIndex].Position
			targetDistance = unitDistance
		}
		if targetDistance > attackRange-1 {
			new.Units[i].Position = rl.Vector2MoveTowards(new.Units[i].Position, target, 4)
		}

		// Ranged shoot projectiles
		if new.Units[i].Type == Ranged {
			if new.Units[i].AttackTimer >= 40 && targetDistance <= attackRange {
				direction := rl.Vector2Normalize(rl.Vector2Subtract(target, new.Units[i].Position))
				new.Projectiles = append(new.Projectiles, Projectile{
					new.Units[i].Position, rl.Vector2Scale(direction, 3),
				})
				new.Units[i].AttackTimer = 0
			} else {
				new.Units[i].AttackTimer++
			}
		}

		// Prevent friendlies going into each other
		for _, u := range new.Units {
			if rl.Vector2Distance(u.Position, new.Units[i].Position) < 35 {
				new.Units[i].Position = rl.Vector2MoveTowards(new.Units[i].Position, u.Position, -1.5)
			}
		}

		// Take damage from enemy projectiles
		for _, p := range other.Projectiles {
			if rl.Vector2Distance(new.Units[i].Position, p.Position) < 10 {
				new.Units[i].Hp -= 50
			}
			if rl.Vector2Distance(new.Tower.Position, p.Position) < 10 {
				new.Tower.Hp -= 5
			}
		}
	}

	// Take damage from enemy melee
	for _, u := range other.Units {
		if u.Type != Melee {
			continue
		}

		const attackRange = 50

		if rl.Vector2Distance(u.Position, self.Tower.Position) < attackRange {
			new.Tower.Hp -= attackDamage
		}

		index, dist := u.closestEnemy(self)
		if index != -1 && dist < attackRange {
			new.Units[index].Hp -= attackDamage
		}
	}

	// Delete dead units
	newUnits := []Unit{}
	for _, u := range new.Units {
		if u.Hp > 0 {
			newUnits = append(newUnits, u)
		}
	}
	new.Units = newUnits

	return new
}

func (self Unit) closestEnemy(enemy TeamState) (index int, distance float32) {
	index = -1
	minDist := float32(1e38)
	for i, u := range enemy.Units {
		dist := rl.Vector2Distance(self.Position, u.Position)
		if dist < minDist {
			minDist = dist
			index = i
		}
	}
	return index, minDist
}
