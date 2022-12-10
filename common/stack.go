package common

import (
	"errors"
	"fmt"
)

type Stack[T any] struct {
	contents []T
}

func (s *Stack[T]) Put(items []T) {
	s.contents = append(s.contents, items...)
}

func (s *Stack[T]) PutItem(item T) {
	s.contents = append(s.contents, item)
}

func (s *Stack[T]) PopItem() (*T, error) {
	if len(s.contents) == 0 {
		return nil, errors.New("empty stack")
	}
	idx := len(s.contents) - 1
	item := &s.contents[idx]
	s.contents = s.contents[:idx]
	return item, nil
}

func (s Stack[T]) Count() int {
	return len(s.contents)
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.contents) == 0
}

func (s *Stack[T]) Pop(n int, reverse bool) []T {
	elements := make([]T, 0)
	for i := 0; i < n; i++ {
		item, err := s.PopItem()
		if err == nil {
			if (reverse) {
				elements = append(elements, *item)
			} else {
				elements = append([]T{*item}, elements...)
			}
		}
	}	
	return elements
}

func (s Stack[T]) TopItem() T {
	return s.contents[len(s.contents)-1]
}

func (s Stack[T]) String() string {
	return fmt.Sprintf("%v", s.contents)
}