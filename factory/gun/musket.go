package factory

type musket struct {
	Gun
}

// GetName implements IGun.
// Subtle: this method shadows the method (Gun).GetName of musket.Gun.
func (m *musket) GetName() string {
	panic("unimplemented")
}

// GetPower implements IGun.
// Subtle: this method shadows the method (Gun).GetPower of musket.Gun.
func (m *musket) GetPower() string {
	panic("unimplemented")
}

// SetName implements IGun.
// Subtle: this method shadows the method (Gun).SetName of musket.Gun.
func (m *musket) SetName(name string) {
	panic("unimplemented")
}

// SetPower implements IGun.
// Subtle: this method shadows the method (Gun).SetPower of musket.Gun.
func (m *musket) SetPower(name int) {
	panic("unimplemented")
}

func newMusket() IGun {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 4,
		},
	}
}
