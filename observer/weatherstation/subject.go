package observer

// Subject interface
type WeatherStation interface {
	Register(observer WeatherObserver)
	Deregister(observer WeatherObserver)
	NotifyAll()
	SetMeasurements(temp float32, humidity float32)
}

// Concrete subject
type BasicWeatherStation struct {
	observers map[string]WeatherObserver
	temp      float32
	humidity  float32
}

func NewWeatherStation() *BasicWeatherStation {
	return &BasicWeatherStation{
		observers: make(map[string]WeatherObserver),
	}
}

func (b *BasicWeatherStation) Register(observer WeatherObserver) {
	b.observers[observer.GetID()] = observer
}

func (b *BasicWeatherStation) Deregister(observer WeatherObserver) {
	delete(b.observers, observer.GetID())
}

func (b *BasicWeatherStation) NotifyAll() {
	for _, ob := range b.observers {
		ob.Update(b.temp, b.humidity)
	}
}

func (b *BasicWeatherStation) SetMeasurements(temp, humidity float32) {
	b.temp = temp
	b.humidity = humidity
}
