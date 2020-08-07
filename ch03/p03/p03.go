package p03

import (
	"fmt"

	"github.com/alcortesm/ctci/ch03"
)

type StackOfStacks struct {
	stacks  []*ch03.Stack
	maxSize int
}

func NewStackOfStacks(maxSize int) (*StackOfStacks, error) {
	if maxSize < 1 {
		return nil, fmt.Errorf("invalid maximum size (%d)", maxSize)
	}

	return &StackOfStacks{
		stacks:  nil,
		maxSize: maxSize,
	}, nil
}

func (s *StackOfStacks) Dump() [][]int {
	result := [][]int{}

	for _, s := range s.stacks {
		inner := make([]int, len(*s))
		copy(inner, *s)
		result = append(result, inner)
	}

	return result
}

func (s *StackOfStacks) Push(v int) {
	if len(s.stacks) == 0 {
		s.stacks = []*ch03.Stack{{}}
	}

	topStack := s.stacks[len(s.stacks)-1]

	if topStack.Len() == s.maxSize {
		topStack = &ch03.Stack{}
		s.stacks = append(s.stacks, topStack)
	}

	topStack.Push(v)
}

func (s *StackOfStacks) Pop() (int, bool) {
	if len(s.stacks) == 0 {
		return 0, false
	}

	topStack := s.stacks[len(s.stacks)-1]
	result, _ := topStack.Pop()

	if topStack.Len() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return result, true
}

func (s *StackOfStacks) Peek() (int, bool) {
	if len(s.stacks) == 0 {
		return 0, false
	}

	topStack := s.stacks[len(s.stacks)-1]
	return topStack.Peek()
}

func (s *StackOfStacks) Len() int {
	result := 0

	for _, s := range s.stacks {
		result += s.Len()
	}

	return result
}