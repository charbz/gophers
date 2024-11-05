// Package collections provides utilities for working with generic collections of elements
package collections

import (
	"fmt"
	"slices"
)

type Sliceable interface {
	ToSlice() int
}

// Collection represents a generic collection of elements of type T
type Collection[T any] struct {
	elements []T
}

// NewCollection creates a new Collection from zero or more slices of type T.
// If no slices are provided, returns an empty Collection.
// If multiple slices are provided, they are concatenated.
func NewCollection[T any](s ...[]T) Collection[T] {
	if len(s) == 0 {
		return Collection[T]{}
	}
	return Collection[T]{elements: slices.Concat(s...)}
}

// zero returns an empty slice of type T.
func (c *Collection[T]) zero() []T {
	var z []T
	return z
}

// zeroT returns the zero value for type T.
func (c *Collection[T]) zeroT() T {
	var z T
	return z
}

// At returns the element at the specified index in the Collection.
func (c *Collection[T]) At(index int) T {
	return c.elements[index]
}

// ToSlice returns the underlying slice containing all elements in the Collection.
func (c *Collection[T]) ToSlice() []T {
	return c.elements
}

// IsEmpty returns true if the Collection contains no elements.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// NonEmpty returns true if the Collection contains at least one element.
func (c *Collection[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Length returns the number of elements in the Collection.
func (c *Collection[T]) Length() int {
	return len(c.elements)
}

// Head returns the first element in the Collection and nil error if non-empty,
// or zero value and error if empty.
func (c *Collection[T]) Head() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), fmt.Errorf("head of empty list")
	}
	return c.elements[0], nil
}

// Last returns the final element in the Collection and nil error if non-empty,
// or zero value and error if empty.
func (c *Collection[T]) Last() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), fmt.Errorf("last of empty list")
	}
	return c.elements[len(c.elements)-1], nil
}

// Tail returns a slice containing all elements except the first.
// If the Collection is empty, returns an empty slice.
func (c *Collection[T]) Tail() []T {
	if c.IsEmpty() {
		return c.elements
	}
	return c.elements[1:]
}

// Init returns a slice containing all elements except the last.
// If the Collection is empty, returns an empty slice.
func (c *Collection[T]) Init() []T {
	if c.IsEmpty() {
		return c.elements
	}
	return c.elements[0 : len(c.elements)-1]
}

// Take returns a slice containing the first q elements.
// If q is negative or zero, returns an empty slice.
// If q exceeds the Collection length, returns all elements.
func (c *Collection[T]) Take(q int) []T {
	if q <= 0 {
		return c.zero()
	}
	return c.elements[0:min(q, c.Length())]
}

// TakeRight returns a slice containing the last q elements.
// If q is negative or zero, returns an empty slice.
// If q exceeds the Collection length, returns all elements.
func (c *Collection[T]) TakeRight(q int) []T {
	if q <= 0 {
		return c.zero()
	}
	return c.elements[max(c.Length()-q, 0):]
}

// Drop returns a slice with the first q elements removed.
// If q is negative or zero, returns all elements.
// If q exceeds the Collection length, returns an empty slice.
func (c *Collection[T]) Drop(q int) []T {
	if q <= 0 {
		return c.elements
	} else if q >= c.Length() {
		return c.zero()
	}
	return c.elements[q:]
}

// DropRight returns a slice with the last q elements removed.
// If q is negative or zero, returns all elements.
// If q exceeds the Collection length, returns an empty slice.
func (c *Collection[T]) DropRight(q int) []T {
	if q <= 0 {
		return c.elements
	} else if q >= c.Length() {
		return c.zero()
	}
	return c.elements[0 : c.Length()-q]
}
