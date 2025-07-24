package pubsub

import "fmt"

type Subscriber interface {
	ReceiveEvent(event *Event)
}

type PoliticalNewSubscriber struct {
	Name string
}

func (p *PoliticalNewSubscriber) ReceiveEvent(event *Event) {
	fmt.Printf("[Political] %s received: %s\n", p.Name, event.Data)
}

type SportsNewsSubscriber struct {
	Name string
}

func (s *SportsNewsSubscriber) ReceiveEvent(event *Event) {
	fmt.Printf("[Sports] %s received : %s\n", s.Name, event.Data)
}
