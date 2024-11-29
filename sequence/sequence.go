// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package sequence implements support for a generic ordered sequence.
// A Sequence is a Collection that wraps an underlying Go slice and provides
// convenience methods and syntatic sugar on top of it.
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

	"github.com/charbz/gophers/collection"
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

// Add appends an element to the sequence.
func (c *Sequence[T]) Add(v T) {
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

// Apply applies a function to each element in the sequence.
func (c *Sequence[T]) Apply(f func(T) T) *Sequence[T] {
	for i := range c.elements {
		c.elements[i] = f(c.elements[i])
	}
	return c
}

// The following methods are mostly syntatic sugar
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

// Concatenated is an alias for collection.Concatenated
func (c *Sequence[T]) Concatenated(s *Sequence[T]) iter.Seq[T] {
	return collection.Concatenated(c, s)
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

// Diff is an alias for collection.Diff
func (c *Sequence[T]) Diff(s *Sequence[T], f func(T, T) bool) *Sequence[T] {
	return collection.DiffFunc(c, s, f).(*Sequence[T])
}

// Diffed is an alias for collection.Diffed
func (c *Sequence[T]) Diffed(s *Sequence[T], f func(T, T) bool) iter.Seq[T] {
	return collection.DiffedFunc(c, s, f)
}

// Distinct takes an "equality" function as an argument
// and returns a new sequence containing all the unique elements
// If you prefer not to pass an equality function use a ComparableSequence.
func (c *Sequence[T]) Distinct(f func(T, T) bool) *Sequence[T] {
	return collection.Distinct(c, f).(*Sequence[T])
}

// DistinctIterator is an alias for collection.DistinctIterator
func (c *Sequence[T]) Distincted(f func(T, T) bool) iter.Seq[T] {
	return collection.DistinctedFunc(c, f)
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

// FilterIterator is an alias for collection.FilterIterator
func (c *Sequence[T]) Filtered(f func(T) bool) iter.Seq[T] {
	return collection.Filtered(c, f)
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

// ForAll is an alias for collection.ForAll
func (c *Sequence[T]) ForAll(f func(T) bool) bool {
	return collection.ForAll(c, f)
}

// Head is an alias for collection.Head
func (c *Sequence[T]) Head() (T, error) {
	return collection.Head(c)
}

// Init is an alias for collection.Init
func (c *Sequence[T]) Init() *Sequence[T] {
	return collection.Init(c).(*Sequence[T])
}

// Intersect is an alias for collection.Intersect
func (c *Sequence[T]) Intersect(s *Sequence[T], f func(T, T) bool) *Sequence[T] {
	return collection.IntersectFunc(c, s, f).(*Sequence[T])
}

// IntersectIterator is an alias for collection.IntersectIterator
func (c *Sequence[T]) Intersected(s *Sequence[T], f func(T, T) bool) iter.Seq[T] {
	return collection.IntersectedFunc(c, s, f)
}

// IsEmpty returns true if the sequence is empty.
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

// SplitAt splits the sequence at the given index.
func (c *Sequence[T]) SplitAt(n int) (*Sequence[T], *Sequence[T]) {
	left := NewSequence(c.elements[:n+1])
	right := NewSequence(c.elements[n+1:])
	return left, right
}

// Reverse is an alias for collection.Reverse
func (c *Sequence[T]) Reverse() *Sequence[T] {
	return collection.Reverse(c).(*Sequence[T])
}

// Reject is an alias for collection.FilterNot
func (l *Sequence[T]) Reject(f func(T) bool) *Sequence[T] {
	return collection.FilterNot(l, f).(*Sequence[T])
}

// Rejected is an alias for collection.Rejected
func (c *Sequence[T]) Rejected(f func(T) bool) iter.Seq[T] {
	return collection.Rejected(c, f)
}

// String implements the Stringer interface.
func (c *Sequence[T]) String() string {
	return fmt.Sprintf("Seq(%T) %v", *new(T), c.elements)
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

func (c *Sequence[T]) Shuffle() *Sequence[T] {
	return collection.Shuffle(c).(*Sequence[T])
}
