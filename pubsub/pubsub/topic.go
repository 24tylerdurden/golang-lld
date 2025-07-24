package pubsub

type Topic struct {
	Id   int
	Desc string
}

func NewTopic(id int, desc string) *Topic {
	return &Topic{
		Id:   id,
		Desc: desc,
	}
}
