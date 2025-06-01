package internal

func (self GameState) Update() GameState {
	return GameState{self.AState.update(), self.BState.update()}
}

func (self TeamState) update() TeamState {
	// TODO: update unit positions, tower hp and stuff
	return self
}
