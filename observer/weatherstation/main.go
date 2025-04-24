package observer

func main() {

	station := NewWeatherStation()

	phone := &PhoneDisplay{id: "iphone"}
	web := &PhoneDisplay{id: "web"}
	tv := &PhoneDisplay{id: "tv"}

	station.Register(phone)
	station.Register(web)
	station.Register(tv)

	station.SetMeasurements(25.5, 60.2)
	station.Deregister(web)
	station.SetMeasurements(34.5, 75.8)

}
