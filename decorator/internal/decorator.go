package internal

import "fmt"

// Componet interface
type Message interface {
	GetMessage() string
}

// Base component
type TextMessage struct {
	message string
}

func (t *TextMessage) GetMessage() string {
	return t.message
}

// skipping constructors methods here, doing this in client side
// decorator -> aggregation relationship
type MessageDecorator struct {
	message Message
}

func (d *MessageDecorator) GetMessage() string {
	return d.message.GetMessage()
}

type TimestampDecorator struct {
	MessageDecorator
}

func NewTimestampeDecorator(msg Message) *TimestampDecorator {
	return &TimestampDecorator{
		MessageDecorator: MessageDecorator{message: msg},
	}
}

func (t *TimestampDecorator) GetMessage() string {
	return fmt.Sprintf("[%s] %s", "2023-01-01 12:00:00", t.MessageDecorator.GetMessage())
}

func TestTimeStampDecorator() {
	msg := &TextMessage{message: "hello world"}

	tsd := NewTimestampeDecorator(msg)

	fmt.Println(tsd.GetMessage())
}

// throught composition or thorught interface
