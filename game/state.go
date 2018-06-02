package game

import "github.com/octomarat/heroes-go/units"

type State struct {
	selectedUnit *units.Unit
}

func (s *State) IsSelected(unit *units.Unit) bool {
	return s.selectedUnit == unit
}

func (s *State) SetSelectedUnit(unit *units.Unit) {
	s.selectedUnit = unit
}
