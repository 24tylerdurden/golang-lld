package pubsub

import "fmt"

type Subscriber interface {
	Events() <-chan *Event
	ReceiveEvent(event *Event)
}

type PoliticalNewSubscriber struct {
	Name    string
	eventCh chan *Event
}

func NewPoliticalNewsSubscriber(name string) *PoliticalNewSubscriber {
	return &PoliticalNewSubscriber{
		Name:    name,
		eventCh: make(chan *Event, 10),
	}
}

func (p *PoliticalNewSubscriber) ReceiveEvent(event *Event) {
	p.eventCh <- event
}

func (p *PoliticalNewSubscriber) Listen() {
	go func() {
		for newsEvent := range p.eventCh {
			fmt.Printf("[Political] %s received: %s\n", p.Name, newsEvent.Data)
		}
	}()
}

func (p *PoliticalNewSubscriber) Events() <-chan *Event {
	return p.eventCh
}

type SportsNewsSubscriber struct {
	Name    string
	eventCh chan *Event
}

func NewSportsNewsSubscriber(name string) *SportsNewsSubscriber {
	return &SportsNewsSubscriber{
		Name:    name,
		eventCh: make(chan *Event, 10),
	}
}

func (s *SportsNewsSubscriber) ReceiveEvent(event *Event) {
	s.eventCh <- event
}

func (s *SportsNewsSubscriber) Events() <-chan *Event {
	return s.eventCh
}

func (s *SportsNewsSubscriber) Listen() {
	go func() {
		for sportsNews := range s.eventCh {
			fmt.Printf("[Sports] %s received : %s\n", s.Name, sportsNews.Data)
		}
	}()
}
