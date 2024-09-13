package stackonlinkedlist

import "fmt"

type Stack struct {
	tail *node
}

type node struct {
	value rune
	prev  *node
}

func (s *Stack) isEmpty() bool {
	return s.tail == nil
}

func (s *Stack) pop() (rune, error) {
	if s.isEmpty() {
		return ' ', fmt.Errorf("stack is empty")
	}

	value := s.tail.value
	s.tail = s.tail.prev

	return value, nil
}

func (s *Stack) push(value rune) {
	node := &node{
		value: value,
		prev:  s.tail,
	}
	s.tail = node
}
