package factory

type IGun interface {
	SetName(name string)
	SetPower(name int)
	GetName() string
	GetPower() string
}
