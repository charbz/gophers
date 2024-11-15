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

func (s *Set[T]) Append(v T) {
	s.elements[v] = struct{}{}
}

func (s *Set[T]) Length() int {
	return len(s.elements)
}

func (s *Set[T]) Random() T {
	for v := range s.elements {
		return v
	}
	panic(collection.EmptyCollectionError)
}

func (s *Set[T]) New(s2 ...[]T) collection.Collection[T] {
	return NewSet(s2...)
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
