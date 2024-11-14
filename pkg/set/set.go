package set

import (
	"fmt"
	"iter"

	"github.com/charbz/gophers/pkg/collection"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](s ...[]T) *Set[T] {
	set := new(Set[T])
	set.elements = make(map[T]struct{})
	for _, slice := range s {
		for _, v := range slice {
			set.elements[v] = struct{}{}
		}
	}
	return set
}

// Implement the Collection interface.

func (s *Set[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for v := range s.elements {
			if !yield(i, v) {
				break
			}
		}
	}
}

func (s *Set[T]) At(index int) T {
	if index < 0 || index >= len(s.elements) {
		panic(collection.IndexOutOfBoundsError)
	}
	for i, v := range s.All() {
		if i == index {
			return v
		}
	}
	panic(collection.IndexOutOfBoundsError)
}

func (s *Set[T]) Append(v T) {
	s.elements[v] = struct{}{}
}

func (s *Set[T]) Backward() iter.Seq2[int, T] {
	slice := s.ToSlice()
	return func(yield func(int, T) bool) {
		for i := len(slice) - 1; i >= 0; i-- {
			if !yield(i, slice[i]) {
				break
			}
		}
	}
}

func (s *Set[T]) Length() int {
	return len(s.elements)
}

func (s *Set[T]) New(s2 ...[]T) collection.Collection[T] {
	return NewSet(s2...)
}

func (s *Set[T]) Slice(start, end int) collection.Collection[T] {
	slice := s.ToSlice()[start:end]
	return NewSet(slice)
}

func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.elements {
			if !yield(v) {
				break
			}
		}
	}
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for v := range s.elements {
		slice = append(slice, v)
	}
	return slice
}

func (s *Set[T]) String() string {
	return fmt.Sprintf("Set(%T) %v", *new(T), s.ToSlice())
}
