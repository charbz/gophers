package collections

import (
	"iter"
	"slices"

	"github.com/charbz/gophers/pkg/utils"
)

func (c *Collection[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

func (c *Collection[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}

func (c *Collection[T]) Filter(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.Filter(c.elements, f),
	}
}

func (c *Collection[T]) FilterNot(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.FilterNot(c.elements, f),
	}
}

func (c *Collection[T]) Partition(f func(T) bool) (*Collection[T], *Collection[T]) {
	left, right := utils.Partition(c.elements, f)
	return &Collection[T]{left}, &Collection[T]{right}
}

func (c *Collection[T]) ForEach(f func(T)) *Collection[T] {
	for v := range c.Values() {
		f(v)
	}
	return c
}

func Map[T any, K any](s *Collection[T], f func(T) K) *Collection[K] {
	return &Collection[K]{
		utils.Map(s.elements, f),
	}
}

func Reduce[T any, K any](s *Collection[T], f func(K, T) K, init K) K {
	return utils.Reduce(s.elements, f, init)
}
