// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package sequence implements support for a generic ordered sequence.
// A Sequence is a Collection that wraps an underlying Go slice and provides
// convenience methods and synthatic sugar on top of it.
//
// Compared to a List, a Sequence allows for efficient O(1) access to arbitrary elements
// but slower insertion and removal time, making it ideal for situations where fast random access is needed.
//
// for comparable types it is recommended to use
// ComparableSequence which provides additional methods
package sequence

import (
	"fmt"
	"iter"
	"math/rand"
	"slices"

	"github.com/charbz/gophers/pkg/collection"
)

type Sequence[T any] struct {
	elements []T
}

func NewSequence[T any](s ...[]T) *Sequence[T] {
	seq := new(Sequence[T])
	if len(s) == 0 {
		return seq
	}
	return &Sequence[T]{elements: slices.Concat(s...)}
}

// The following methods implement
// the Collection interface.

// Append appends an element to the sequence.
func (c *Sequence[T]) Append(v T) {
	c.elements = append(c.elements, v)
}

// Length returns the number of elements in the sequence.
func (c *Sequence[T]) Length() int {
	return len(c.elements)
}

// New is a constructor for a generic sequence.
func (c *Sequence[T]) New(s ...[]T) collection.Collection[T] {
	return NewSequence(s...)
}

// Random returns a random element from the sequence.
func (c *Sequence[T]) Random() T {
	if len(c.elements) == 0 {
		return *new(T)
	}
	return c.elements[rand.Intn(len(c.elements))]
}

// Values returns an iterator over all values of the underlying slice.
func (c *Sequence[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}

// The following methods implement
// the OrderedCollection interface.

// At returns the element at the given index.
func (c *Sequence[T]) At(index int) T {
	if index < 0 || index >= len(c.elements) {
		panic(collection.IndexOutOfBoundsError)
	}
	return c.elements[index]
}

// All returns an iterator over all elements of the sequence.
func (c *Sequence[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

// Backward returns an iterator over all elements of the sequence in reverse order.
func (c *Sequence[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(c.elements)
}

// Slice returns a new sequence containing the elements from the start index to the end index.
func (c *Sequence[T]) Slice(start, end int) collection.OrderedCollection[T] {
	return &Sequence[T]{
		c.elements[start:end],
	}
}

// NewOrdered returns a new ordered collection.
func (c *Sequence[T]) NewOrdered(s ...[]T) collection.OrderedCollection[T] {
	return NewSequence(s...)
}

// The following methods are mostly synthatic sugar
// wrapping Collection functions to enable function chaining:
// i.e. sequence.Filter(f).Take(n)

// Clone returns a copy of the collection. This is a shallow clone.
func (c *Sequence[T]) Clone() *Sequence[T] {
	return &Sequence[T]{
		slices.Clone(c.elements),
	}
}

// Count is an alias for collection.Count
func (c *Sequence[T]) Count(f func(T) bool) int {
	return collection.Count(c, f)
}

// Concat returns a new sequence concatenating the passed in sequences.
func (c *Sequence[T]) Concat(sequences ...Sequence[T]) *Sequence[T] {
	e := c.elements
	for _, col := range sequences {
		e = slices.Concat(e, col.elements)
	}
	return &Sequence[T]{e}
}

// Contains tests whether a predicate holds for at least one element of this sequence.
func (c *Sequence[T]) Contains(f func(T) bool) bool {
	i, _ := collection.Find(c, f)
	return i > -1
}

// Corresponds is an alias for collection.Corresponds
func (c *Sequence[T]) Corresponds(s *Sequence[T], f func(T, T) bool) bool {
	return collection.Corresponds(c, s, f)
}

// Dequeue removes and returns the first element of the sequence.
func (c *Sequence[T]) Dequeue() (T, error) {
	if len(c.elements) == 0 {
		return *new(T), collection.EmptyCollectionError
	}
	element := c.elements[0]
	c.elements = c.elements[1:]
	return element, nil
}

// Distinct takes an "equality" function as an argument
// and returns a new sequence containing all the unique elements
// If you prefer not to pass an equality function use a ComparableSequence.
func (c *Sequence[T]) Distinct(f func(T, T) bool) *Sequence[T] {
	return &Sequence[T]{
		slices.CompactFunc(c.elements, f),
	}
}

// Drop is an alias for collection.Drop
func (c *Sequence[T]) Drop(n int) *Sequence[T] {
	return collection.Drop(c, n).(*Sequence[T])
}

// DropWhile is an alias for collection.DropWhile
func (c *Sequence[T]) DropWhile(f func(T) bool) *Sequence[T] {
	return collection.DropWhile(c, f).(*Sequence[T])
}

// DropRight is an alias for collection.DropRight
func (c *Sequence[T]) DropRight(n int) *Sequence[T] {
	return collection.DropRight(c, n).(*Sequence[T])
}

// Enqueue appends an element to the sequence.
func (c *Sequence[T]) Enqueue(v T) {
	c.elements = append(c.elements, v)
}

// Equals takes a sequence and an equality function as an argument
// and returns true if the two sequences are equal.
// If you prefer not to pass an equality function use a ComparableSequence.
func (c *Sequence[T]) Equals(c2 *Sequence[T], f func(T, T) bool) bool {
	return slices.EqualFunc(c.elements, c2.elements, f)
}

// Exists is an alias for Contains
func (c *Sequence[T]) Exists(f func(T) bool) bool {
	return c.Contains(f)
}

// Filter is an alias for collection.Filter
func (c *Sequence[T]) Filter(f func(T) bool) *Sequence[T] {
	return collection.Filter(c, f).(*Sequence[T])
}

// FilterNot is an alias for collection.FilterNot
func (c *Sequence[T]) FilterNot(f func(T) bool) *Sequence[T] {
	return collection.FilterNot(c, f).(*Sequence[T])
}

// Find is an alias for collection.Find
func (c *Sequence[T]) Find(f func(T) bool) (int, T) {
	return collection.Find(c, f)
}

// FindLast is an alias for collection.FindLast
func (c *Sequence[T]) FindLast(f func(T) bool) (int, T) {
	return collection.FindLast(c, f)
}

// ForEach is an alias for collection.ForEach
func (c *Sequence[T]) ForEach(f func(T)) *Sequence[T] {
	return collection.ForEach(c, f).(*Sequence[T])
}

// Head is an alias for collection.Head
func (c *Sequence[T]) Head() (T, error) {
	return collection.Head(c)
}

// Init is an alias for collection.Init
func (c *Sequence[T]) Init() *Sequence[T] {
	return collection.Init(c).(*Sequence[T])
}

func (c *Sequence[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// Last is an alias for collection.Last
func (c *Sequence[T]) Last() (T, error) {
	return collection.Last(c)
}

// returns true if the sequence is not empty.
func (c *Sequence[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Pop removes and returns the last element of the sequence.
func (c *Sequence[T]) Pop() (T, error) {
	if len(c.elements) == 0 {
		return *new(T), collection.EmptyCollectionError
	}
	element := c.elements[len(c.elements)-1]
	c.elements = c.elements[:len(c.elements)-1]
	return element, nil
}

// Push appends an element to the sequence.
func (c *Sequence[T]) Push(v T) {
	c.elements = append(c.elements, v)
}

// Partition is an alias for collection.Partition
func (c *Sequence[T]) Partition(f func(T) bool) (*Sequence[T], *Sequence[T]) {
	left, right := collection.Partition(c, f)
	return left.(*Sequence[T]), right.(*Sequence[T])
}

// Reverse is an alias for collection.Reverse
func (c *Sequence[T]) Reverse() *Sequence[T] {
	return collection.Reverse(c).(*Sequence[T])
}

// String implements the Stringer interface.
func (c *Sequence[T]) String() string {
	return fmt.Sprintf("Seq(%T) %v", *new(T), c.elements)
}

// SplitAt is an alias for collection.SplitAt
func (c *Sequence[T]) SplitAt(n int) (*Sequence[T], *Sequence[T]) {
	left, right := collection.SplitAt(c, n)
	return left.(*Sequence[T]), right.(*Sequence[T])
}

// Take is an alias for collection.Take
func (c *Sequence[T]) Take(n int) *Sequence[T] {
	return collection.Take(c, n).(*Sequence[T])
}

// TakeRight is an alias for collection.TakeRight
func (c *Sequence[T]) TakeRight(n int) *Sequence[T] {
	return collection.TakeRight(c, n).(*Sequence[T])
}

// Tail is an alias for collection.Tail
func (c *Sequence[T]) Tail() *Sequence[T] {
	return collection.Tail(c).(*Sequence[T])
}

// ToSlice returns the underlying slice.
func (c *Sequence[T]) ToSlice() []T {
	return c.elements
}
