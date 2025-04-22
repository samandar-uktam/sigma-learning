package factory

type Ak47 struct {
	Gun
}

// GetName implements IGun.
// Subtle: this method shadows the method (Gun).GetName of Ak47.Gun.
func (a *Ak47) GetName() string {
	panic("unimplemented")
}

// GetPower implements IGun.
// Subtle: this method shadows the method (Gun).GetPower of Ak47.Gun.
func (a *Ak47) GetPower() string {
	panic("unimplemented")
}

// SetName implements IGun.
// Subtle: this method shadows the method (Gun).SetName of Ak47.Gun.
func (a *Ak47) SetName(name string) {
	panic("unimplemented")
}

// SetPower implements IGun.
// Subtle: this method shadows the method (Gun).SetPower of Ak47.Gun.
func (a *Ak47) SetPower(name int) {
	panic("unimplemented")
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
