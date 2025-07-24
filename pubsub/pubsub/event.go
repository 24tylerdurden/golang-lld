package pubsub

type Event struct {
	Data string
}

func NewEvent(data string) *Event {
	return &Event{Data: data}
}
