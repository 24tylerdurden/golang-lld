package pubsub

import "sync"

type PubSub struct {
	mu          sync.Mutex
	topics      map[int]*Topic
	subscribers map[int][]Subscriber // topicId --> subscribers
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics:      make(map[int]*Topic),
		subscribers: make(map[int][]Subscriber),
	}
}

func (p *PubSub) AddToTopic(topic *Topic) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.topics[topic.Id] = topic
}

func (p *PubSub) AddSubscriberToTopic(top *Topic, sub Subscriber) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.subscribers[top.Id] = append(p.subscribers[top.Id], sub)
}

func (p *PubSub) PublishToTopic(topic *Topic, event *Event) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, sub := range p.subscribers[topic.Id] {
		go sub.ReceiveEvent(event)
	}
}
