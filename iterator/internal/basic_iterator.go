package internal

import (
	"errors"
	"fmt"
)

// Iterator interface
type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
}

// Iterable interface
type Iterable[T any] interface {
	CreateIterator() Iterator[T]
}

// concrete slice based list

type List[T any] struct {
	items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		items: make([]T, 0),
	}
}

// add
func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// get
func (l *List[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(l.items) {
		return zero, fmt.Errorf("index out of bounds")
	}

	return l.items[index], nil
}

// size
func (l *List[T]) Size() int {
	return len(l.items)
}

// Concrete Iterator -
type ListIterator[T any] struct {
	list     *List[T]
	position int
}

func NewListIterator[T any](list *List[T]) *ListIterator[T] {
	return &ListIterator[T]{
		list:     list,
		position: 0,
	}
}

func (li *ListIterator[T]) HasNext() bool {
	return li.position < li.list.Size()
}

func (li *ListIterator[T]) Next() (T, error) {
	var zero T
	if !li.HasNext() {
		return zero, errors.New("no more elements")
	}
	item := li.list.items[li.position]
	li.position++
	return item, nil
}

// Iterable Interface
func (l *List[T]) CreateIterator() Iterator[T] {
	return NewListIterator(l)
}

func TestBasicIterator() {
	list := NewList[string]()
	list.Add("First")
	list.Add("Second")
	list.Add("Third")

	// This is list Iterator
	itr := list.CreateIterator()

	for itr.HasNext() {
		item, err := itr.Next()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(item)
	}
}
