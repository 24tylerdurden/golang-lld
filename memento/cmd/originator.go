package main

import "time"

// this is the state which we want to preserve
type Originator struct {
	state      string
	careTacker *CareTracker
}

func NewOriginator() Originator {
	return Originator{
		careTacker: NewCareTacker(),
	}
}

func (o *Originator) AddText(text string) {
	o.state += text
}

func (o *Originator) createMemento() Memento {
	return Memento{
		state:     o.state,
		timestamp: time.Now(),
	}
}

func (o *Originator) SavePoint() {
	mem := o.createMemento()
	o.careTacker.Add(mem)
}

func (o *Originator) RestoreSavePoint() {
	mem := o.careTacker.Pop()

	o.state = mem.state
}

func (o *Originator) ListSavePoints() []Memento {
	return o.careTacker.ListAllSavePoints()
}

func (o *Originator) GetState() string {
	return o.state
}

// create memento -> create memento

// store memento -> store create memento into caretacker

// restore memento state
