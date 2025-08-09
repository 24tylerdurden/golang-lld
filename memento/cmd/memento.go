package main

import "time"

// preserve the state
type Memento struct {
	state     string
	timestamp time.Time
}
