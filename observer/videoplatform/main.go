package observer

func main() {
	agency := NewVideoPlatform()
	sub1 := &NewsSubscriber{id: "Bob"}
	sub2 := &NewsSubscriber{id: "Alice"}

	agency.Register(sub1)
	agency.Register(sub2)

	agency.PublishVideos("Golang Course 2.0 is released!")
	agency.Deregister(sub1)
	agency.PublishVideos("Oracle Db Intro!")
}
