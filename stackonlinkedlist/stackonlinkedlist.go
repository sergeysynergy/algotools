package stackonlinkedlist

import "fmt"

type Stack struct {
	tail *node
}

type node struct {
	value any
	prev  *node
}

func (s *Stack) IsEmpty() bool {
	return s.tail == nil
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return ' ', fmt.Errorf("stack is empty")
	}

	value := s.tail.value
	s.tail = s.tail.prev

	return value, nil
}

func (s *Stack) Push(value any) {
	node := &node{
		value: value,
		prev:  s.tail,
	}
	s.tail = node
}
