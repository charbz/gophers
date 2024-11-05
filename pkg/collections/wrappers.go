package collections

import (
	"iter"
	"slices"

	"github.com/charbz/gophers/pkg/utils"
)

func (c *Collection[T]) Filter(f func(T) bool) []T {
	return utils.Filter(c.elements, f)
}

func (c *Collection[T]) FilterNot(f func(T) bool) []T {
	return utils.FilterNot(c.elements, f)
}

func (c *Collection[T]) Partition(f func(T) bool) ([]T, []T) {
	return utils.Partition(c.elements, f)
}

func (c *Collection[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

func (c *Collection[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}
