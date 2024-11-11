// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package sequence implements support for generic Sequences of data.
// A sequence wraps an underlying Go slice and provides convenience methods
// and synthatic sugar on top of it.

// for comparable types it is recommended to use ComparableSequence,
// which provides additional methods for comparable types.
package sequence

import (
	"fmt"
	"iter"
	"slices"

	"github.com/charbz/gophers/pkg/collection"
)

type Sequence[T any] struct {
	elements []T
}

// NewSequence is a constructor for a generic sequence.
func NewSequence[T any](s ...[]T) *Sequence[T] {
	seq := new(Sequence[T])
	if len(s) == 0 {
		return seq
	}
	return &Sequence[T]{elements: slices.Concat(s...)}
}

// At returns the element at the given index.
func (c *Sequence[T]) At(index int) T {
	return c.elements[index]
}

// Append adds a value to the end of the sequence.
func (c *Sequence[T]) Append(v T) {
	c.elements = append(c.elements, v)
}

// All returns an (index, value) iterator for the underlying slice.
func (c *Sequence[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

// Backward returns an (index, value) backwards iterator
// for the underlying slice.
func (c *Sequence[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(c.elements)
}

// Clone returns a copy of the collection. This is a shallow clone.
func (c *Sequence[T]) Clone() *Sequence[T] {
	return &Sequence[T]{
		slices.Clone(c.elements),
	}
}

// Concat returns a new sequence concatenating the passed in sequences.
func (c *Sequence[T]) Concat(sequences ...Sequence[T]) *Sequence[T] {
	e := c.elements
	for _, col := range sequences {
		e = slices.Concat(e, col.elements)
	}
	return &Sequence[T]{e}
}

// Contains tests whether a predicate holds for at least
// one element of this sequence.
//
// example:
//
//	c := NewSequence([]string{"Alice", "bilLy", "JOel"})
//	c.Contains(func (i string) bool {
//	  return strings.ToLower(i) == "joel"
//	})
//
// output:
//
//	true
func (c *Sequence[T]) Contains(f func(T) bool) bool {
	i, _ := collection.Find(c, f)
	return i > -1
}

// Distinct takes a higher order "equality" function as an argument
// and returns a new sequence containing all the unique elements
// from the original sequence.
//
// example:
//
//	c := NewSequence([]int{1,2,2,3,3,3})
//	c.DistinctFunc(func (a int, b int) bool {
//	  return a == b
//	})
//
// output:
//
//	[1,2,3]
//
// If you prefer not to pass an equality function check out
// Distinct() in functions.go
func (c *Sequence[T]) Distinct(f func(T, T) bool) *Sequence[T] {
	return &Sequence[T]{
		slices.CompactFunc(c.elements, f),
	}
}

// Drop returns a new sequence with the first n elements removed.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Drop(2)
//
// output:
//
//	[3,4,5,6]
func (c *Sequence[T]) Drop(n int) *Sequence[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return new(Sequence[T])
	}
	return &Sequence[T]{
		c.elements[n:],
	}
}

// DropRight returns a sequence with the last n elements removed.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.DropRight(2)
//
// output:
//
//	[1,2,3,4]
func (c *Sequence[T]) DropRight(n int) *Sequence[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return new(Sequence[T])
	}
	return &Sequence[T]{
		c.elements[0 : c.Length()-n],
	}
}

// Exists is an alias for Contains
func (c *Sequence[T]) Exists(f func(T) bool) bool {
	return c.Contains(f)
}

// Filter takes a filtering function as input and returns a new sequence
// containing all the elements that match the filter.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Filter(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[2,4,6]
func (c *Sequence[T]) Filter(f func(T) bool) *Sequence[T] {
	return collection.Filter(c, f).(*Sequence[T])
}

// FilterNot takes a filtering function as input and returns a new sequence
// containing all the elements that do not match the filter.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.FilterNot(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[1,3,5]
func (c *Sequence[T]) FilterNot(f func(T) bool) *Sequence[T] {
	return collection.FilterNot(c, f).(*Sequence[T])
}

// Find finds the first element of the sequence satisfying a predicate, if any.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Find(f(i int) bool {
//	  return (i + 3) > 5
//	})
//
// output
//
//	3
func (c *Sequence[T]) Find(f func(T) bool) (T, error) {
	i, v := collection.Find(c, f)
	if i > -1 {
		return v, nil
	}
	return v, collection.ValueNotFoundError
}

// FindWhere finds the index of the first element of the sequence satisfying a predicate.
// If the element is not found, -1 is returned
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.FindWhere(f(i int) int {
//	  return (i + 3) > 5
//	})
//
// output
//
//	2
func (c *Sequence[T]) FindWhere(f func(T) bool) int {
	i, _ := collection.Find(c, f)
	return i
}

// ForEach takes a function as input and applies the function
// to each element in the sequence.
//
// example usage:
//
//	c.ForEach(func(t Task) {
//	  t.run()
//	})
func (c *Sequence[T]) ForEach(f func(T)) *Sequence[T] {
	for v := range c.Values() {
		f(v)
	}
	return c
}

// Head returns the first element in a Sequence and a nil error.
// If the sequence is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewSequence([]string{"A","B","C"})
//	c.Head()
//
// output:
//
//	"A", nil
func (c *Sequence[T]) Head() (T, error) {
	if c.IsEmpty() {
		return *new(T), collection.EmptyCollectionError
	}
	return c.elements[0], nil
}

// Init returns a sequence containing all elements excluding the last one.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Tail()
//
// output:
//
//	[1,2,3,4,5]
func (c *Sequence[T]) Init() *Sequence[T] {
	if c.IsEmpty() {
		return c
	}
	return &Sequence[T]{
		c.elements[0 : len(c.elements)-1],
	}
}

// IsEmpty returns true if the Sequence contains 0 elements.
func (c *Sequence[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// Last returns the last element in the Sequence and a nil error.
// If the sequence is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewSequence([]string{"A","B","C"})
//	c.Last()
//
// output:
//
//	"C", nil
func (c *Sequence[T]) Last() (T, error) {
	if c.IsEmpty() {
		return *new(T), collection.EmptyCollectionError
	}
	return c.elements[len(c.elements)-1], nil
}

// Length returns the number of elements in the Sequence.
func (c *Sequence[T]) Length() int {
	return len(c.elements)
}

// New implements the Collection interface.
// this is useful for generic functions that need to create a new
// instance of the passed concrete collection type.
func (c *Sequence[T]) New(s ...[]T) collection.Collection[T] {
	return NewSequence(s...)
}

// NonEmpty returns true if the Sequence contains at least 1 element.
func (c *Sequence[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Partition takes a partitioning function as input and returns two sequences,
// the first one contains the elements that match the partitioning condition,
// the second one contains the rest of the elements.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Partition(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[2,4,6], [1,3,5]
func (c *Sequence[T]) Partition(f func(T) bool) (*Sequence[T], *Sequence[T]) {
	left, right := collection.Partition(c, f)
	return left.(*Sequence[T]), right.(*Sequence[T])
}

// Reverse returns a new sequence containing all elements in reverse
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Reverse
//
// output:
//
//	[6,5,4,3,2,1]
func (c *Sequence[T]) Reverse() *Sequence[T] {
	elements := make([]T, 0, len(c.elements))
	for i := len(c.elements) - 1; i >= 0; i-- {
		elements = append(elements, c.elements[i])
	}
	return &Sequence[T]{
		elements,
	}
}

// String implements the Stringer interface to
// enable fmt to print the underlying slice.
func (c *Sequence[T]) String() string {
	return fmt.Sprintf("Sequence <%T> %v", *new(T), c.elements)
}

// Take returns a new sequence containing the first n elements.
//
// example usage:
//
//	[c := NewSequence([]int{1,2,3,4,5,6})
//	c.Take(3)
//
// output:
//
//	[1,2,3]
func (c *Sequence[T]) Take(n int) *Sequence[T] {
	if n <= 0 {
		return new(Sequence[T])
	}
	return &Sequence[T]{
		c.elements[0:min(n, c.Length())],
	}
}

// TakeRight returns a new sequence containing the last n elements.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.TakeRight(3)
//
// output:
//
//	[4,5,6]
func (c *Sequence[T]) TakeRight(n int) *Sequence[T] {
	if n <= 0 {
		return new(Sequence[T])
	}
	return &Sequence[T]{
		c.elements[max(c.Length()-n, 0):],
	}
}

// Tail returns a new sequence containing all elements excluding the first one.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Tail()
//
// output:
//
//	[2,3,4,5,6]
func (c *Sequence[T]) Tail() *Sequence[T] {
	if c.IsEmpty() {
		return c
	}
	return &Sequence[T]{
		c.elements[1:],
	}
}

// ToSlice returns the underlying slice.
func (c *Sequence[T]) ToSlice() []T {
	return c.elements
}

// Values returns an iterator over all values of the underlying slice.
func (c *Sequence[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}
