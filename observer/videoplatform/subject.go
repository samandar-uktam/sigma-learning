package observer

import "fmt"

type Channel interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

// Concrete subject
type VideoPlatform struct {
	subscribers map[string]Observer
	videos      string
}

func NewVideoPlatform() *VideoPlatform {
	return &VideoPlatform{
		subscribers: make(map[string]Observer),
	}
}

func (a *VideoPlatform) Register(o Observer) {
	a.subscribers[o.GetID()] = o
}

func (a *VideoPlatform) Deregister(o Observer) {
	delete(a.subscribers, o.GetID())
}

func (a *VideoPlatform) NotifyAll() {
	for _, sub := range a.subscribers {
		sub.Update(a.videos)
	}
}

func (a *VideoPlatform) PublishVideos(videos string) {
	fmt.Println("\n📰 Publishing videos: ", videos)
	a.videos = videos
	a.NotifyAll()
}
