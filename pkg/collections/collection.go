// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package collections provides an implementation of a generic Collection.
// A collection wraps an underlying Go slice and provides convenience methods.
// The design of this package is inspired by Scala's collections library.
package collections

import (
	"fmt"
	"slices"

	"github.com/charbz/gophers/pkg/utils"
)

type Collection[T any] struct {
	elements []T
}

// NewCollection is a constructor for collections it takes
// a variadic argument of slices of any type as input and converts
// them to a single flat collection.
func NewCollection[T any](s ...[]T) Collection[T] {
	if len(s) == 0 {
		return Collection[T]{}
	}
	return Collection[T]{elements: slices.Concat(s...)}
}

// String implements the Stringer interface so fmt
// can print the underlying slice
func (c *Collection[T]) String() string {
	return fmt.Sprintf("%v", c.elements)
}

// zero returns the zero value of the underlying slice
// used internally when an operation must return an empty collection.
func (c *Collection[T]) zero() *Collection[T] {
	var z Collection[T]
	return &z
}

// zeroT returns the zero value of the underlying slice's type.
func (c *Collection[T]) zeroT() T {
	var z T
	return z
}

// At returns the element at the specified index.
func (c *Collection[T]) At(index int) T {
	return c.elements[index]
}

// ToSlice returns the underlying slice.
func (c *Collection[T]) ToSlice() []T {
	return c.elements
}

// IsEmpty returns true if the Collection contains 0 elements.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// NonEmpty returns true if the Collection contains at
// least 1 element.
func (c *Collection[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Length returns the number of elements in the Collection.
func (c *Collection[T]) Length() int {
	return len(c.elements)
}

// Head returns the first element in the Collection and nil
// error if non-empty, otherwise it returns zero value and error.
func (c *Collection[T]) Head() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), EmptyCollectionError
	}
	return c.elements[0], nil
}

// Last returns the final element in the Collection
// and nil error if non-empty, or zero value and error otherwise.
func (c *Collection[T]) Last() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), EmptyCollectionError
	}
	return c.elements[len(c.elements)-1], nil
}

// Tail returns a collection containing all elements except the first.
func (c *Collection[T]) Tail() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[1:],
	}
}

// Init returns a collection containing all elements except the last.
func (c *Collection[T]) Init() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[0 : len(c.elements)-1],
	}
}

// Take returns a collection containing the first n elements.
func (c *Collection[T]) Take(n int) *Collection[T] {
	if n <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0:min(n, c.Length())],
	}
}

// TakeRight returns a collection containing the last n elements.
func (c *Collection[T]) TakeRight(n int) *Collection[T] {
	if n <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[max(c.Length()-n, 0):],
	}
}

// Drop returns a collection with the first n elements removed.
func (c *Collection[T]) Drop(n int) *Collection[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[n:],
	}
}

// DropRight returns a collection with the last n elements removed.
func (c *Collection[T]) DropRight(n int) *Collection[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0 : c.Length()-n],
	}
}

// Filter takes a filtering function as input and returns
// the resulting collection after applying the filter to each element.
func (c *Collection[T]) Filter(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.Filter(c.elements, f),
	}
}

// FilterNot takes a filtering function as input and returns
// the collection containing all elements that do not satisfy
// the filtering function provided.
func (c *Collection[T]) FilterNot(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.FilterNot(c.elements, f),
	}
}

// Partition takes a partitioning function as input and returns
// two collections, the first one containing the elements that
// match the partitioning condition, and the other one contains
// the rest of the elements.
func (c *Collection[T]) Partition(f func(T) bool) (*Collection[T], *Collection[T]) {
	left, right := utils.Partition(c.elements, f)
	return &Collection[T]{left}, &Collection[T]{right}
}

// ForEach takes a function as input and applies the
// function to each element in the collection.
func (c *Collection[T]) ForEach(f func(T)) *Collection[T] {
	for v := range c.Values() {
		f(v)
	}
	return c
}
