package observer

import "fmt"

type Observer interface {
	Update(news string)
	GetID() string
}

// Concrete Observer
type NewsSubscriber struct {
	id string
}

func (s *NewsSubscriber) Update(news string) {
	fmt.Printf("Subscriber %s received news: %s\n", s.id, news)
}

func (s *NewsSubscriber) GetID() string {
	return s.id
}
