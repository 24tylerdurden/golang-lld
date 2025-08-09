package main

type CareTracker struct {
	mementos []Memento
}

func NewCareTacker() *CareTracker {
	return &CareTracker{
		mementos: make([]Memento, 0),
	}
}

func (c *CareTracker) Add(m Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *CareTracker) Pop() Memento {
	// get the last state, restore to that state
	n := len(c.mementos)
	lastMemento := c.mementos[n-1]
	c.mementos = c.mementos[:n-1]

	return lastMemento
}

func (c *CareTracker) ListAllSavePoints() []Memento {
	return c.mementos
}
