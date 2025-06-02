package internal

import rl "github.com/gen2brain/raylib-go/raylib"

type GameState struct {
	AState       TeamState
	BState       TeamState
	showControls bool
}

type TeamState struct {
	IsLocal     bool
	Mana        float32
	Tower       Tower
	Units       []Unit
	Projectiles []Projectile
}

type Tower struct {
	Position rl.Vector2
	Hp       int
}

type Unit struct {
	Position rl.Vector2
	Hp       int
	Type     UnitType
}

type UnitType uint8

const (
	Knight UnitType = iota
)

type Projectile struct {
	Position rl.Vector2
	Velocity rl.Vector2
}

var DefaultState = GameState{
	AState: TeamState{
		IsLocal: true,
		Tower: Tower{
			Position: rl.Vector2{X: 160, Y: 450},
			Hp:       1000,
		},
	},
	BState: TeamState{
		Tower: Tower{
			Position: rl.Vector2{X: 1440, Y: 450},
			Hp:       1000,
		},
	},
	showControls: true,
}
