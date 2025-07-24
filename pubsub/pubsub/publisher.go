package pubsub

type Publisher interface {
	PublishToTopic(topic *Topic, event *Event)
}

type PoliticalNewsPublisher struct {
	Pubsub *PubSub
}

func (p *PoliticalNewsPublisher) PublishToTopic(topic *Topic, event *Event) {
	p.Pubsub.PublishToTopic(topic, event)
}

type SportsNewsPublisher struct {
	Pubsub *PubSub
}

func (s *SportsNewsPublisher) PublishToTopic(topic *Topic, event *Event) {
	s.Pubsub.PublishToTopic(topic, event)
}
