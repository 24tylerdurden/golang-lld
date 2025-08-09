package main

type ICareTacker interface {
	Add(Memento)
	Pop() Memento
	ListAllSavePoints() []Memento
}
