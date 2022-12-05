package common

import (
	"errors"
	"fmt"
)

type Stack struct {
	contents []string
}

func (s *Stack) Put(items []string) {
	s.contents = append(s.contents, items...)
}

func (s *Stack) PutItem(item string) {
	s.contents = append(s.contents, item)
}

func (s *Stack) PopItem() (string, error) {
	if len(s.contents) == 0 {
		return "", errors.New("empty stack")
	}
	idx := len(s.contents) - 1
	item := s.contents[idx]
	s.contents = s.contents[:idx]
	return item, nil
}

func (s Stack) Count() int {
	return len(s.contents)
}

func (s *Stack) Pop(n int, reverse bool) []string {
	elements := make([]string, 0)
	for i := 0; i < n; i++ {
		item, err := s.PopItem()
		if err == nil {
			if (reverse) {
				elements = append(elements, item)
			} else {
				elements = append([]string{item}, elements...)
			}
		}
	}	
	return elements
}

func (s Stack) TopItem() string {
	return s.contents[len(s.contents)-1]
}

func (s Stack) String() string {
	return fmt.Sprintf("%v", s.contents)
}

func NewStack(size int) Stack {
	return Stack{make([]string, size)}
}