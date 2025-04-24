package observer

import "fmt"

// Observer interface
type WeatherObserver interface {
	Update(temp float32, humidity float32)
	GetID() string
}

// Concrete observer
type PhoneDisplay struct {
	id string
}

func (p *PhoneDisplay) Update(temp float32, humidity float32) {
	fmt.Printf("[%s] Phone updated: Temp %.1f°C, Humidity %.1f%%\n", p.id, temp, humidity)
}

func (p *PhoneDisplay) GetID() string {
	return p.id
}
