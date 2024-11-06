// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package collections provides an implementation of a generic Collection.
// Collection wraps an underlying Go slice and provides convenient methods.
// The design of this package is inspired by Scala's collections library.
// This file contains the constructor, and basic introspection methods.
package collections

import (
	"fmt"
	"slices"
)

// Collection represents a generic collection of elements of type T.
type Collection[T any] struct {
	elements []T
}

// NewCollection is a constructor for collections it takes
// a variadic argument of slices of type T as input and converts
// them to a single flat collection.
func NewCollection[T any](s ...[]T) Collection[T] {
	if len(s) == 0 {
		return Collection[T]{}
	}
	return Collection[T]{elements: slices.Concat(s...)}
}

func (c *Collection[T]) String() string {
	return fmt.Sprintf("%v", c.elements)
}

// zero returns the zero value of the underlying slice
// this is only used internally when an operation must
// return an empty collection.
func (c *Collection[T]) zero() *Collection[T] {
	var z Collection[T]
	return &z
}

// zeroT returns the zero value of
// the underlying slice's type.
func (c *Collection[T]) zeroT() T {
	var z T
	return z
}

// At returns the element at the specified
// index in the Collection.
func (c *Collection[T]) At(index int) T {
	return c.elements[index]
}

// ToSlice returns the underlying slice.
func (c *Collection[T]) ToSlice() []T {
	return c.elements
}

// IsEmpty returns true if the Collection contains
// 0 elements.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// NonEmpty returns true if the Collection contains
// at least 1 element.
func (c *Collection[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Length returns the number of elements in the Collection.
func (c *Collection[T]) Length() int {
	return len(c.elements)
}

// Head returns the first element in the Collection
// and nil error if non-empty, or zero value and
// error otherwise.
func (c *Collection[T]) Head() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), EmptyCollectionError
	}
	return c.elements[0], nil
}

// Last returns the final element in the Collection
// and nil error if non-empty, or zero value and
// error otherwise.
func (c *Collection[T]) Last() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), EmptyCollectionError
	}
	return c.elements[len(c.elements)-1], nil
}

// Tail returns a slice containing all elements except the first.
// If the Collection is empty, returns an empty slice.
func (c *Collection[T]) Tail() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[1:],
	}
}

// Init returns a slice containing all elements except the last.
// If the Collection is empty, returns an empty slice.
func (c *Collection[T]) Init() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[0 : len(c.elements)-1],
	}
}

// Take returns a slice containing the first q elements.
// If q is negative or zero, returns an empty slice.
// If q exceeds the Collection length, returns all elements.
func (c *Collection[T]) Take(q int) *Collection[T] {
	if q <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0:min(q, c.Length())],
	}
}

// TakeRight returns a slice containing the last q elements.
// If q is negative or zero, returns an empty slice.
// If q exceeds the Collection length, returns all elements.
func (c *Collection[T]) TakeRight(q int) *Collection[T] {
	if q <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[max(c.Length()-q, 0):],
	}
}

// Drop returns a slice with the first q elements removed.
// If q is negative or zero, returns all elements.
// If q exceeds the Collection length, returns an empty slice.
func (c *Collection[T]) Drop(q int) *Collection[T] {
	if q <= 0 {
		return c
	} else if q >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[q:],
	}
}

// DropRight returns a slice with the last q elements removed.
// If q is negative or zero, returns all elements.
// If q exceeds the Collection length, returns an empty slice.
func (c *Collection[T]) DropRight(q int) *Collection[T] {
	if q <= 0 {
		return c
	} else if q >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0 : c.Length()-q],
	}
}
