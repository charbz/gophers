// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package sequence

import (
	"cmp"
	"iter"
	"slices"

	"github.com/charbz/gophers/collection"
)

// ComparableSequence is a sequence of comparable types.
// it is similar to Sequence, but with additional methods that do not require a
// higher order function comparator to be provided as an argument:
// Max(), Min(), Sum(), Distinct(), Diff(c), and Exists(v).
type ComparableSequence[T cmp.Ordered] struct {
	Sequence[T]
}

func (c *ComparableSequence[T]) New(s ...[]T) collection.Collection[T] {
	return NewComparableSequence(s...)
}

func (c *ComparableSequence[T]) NewOrdered(s ...[]T) collection.OrderedCollection[T] {
	return NewComparableSequence(s...)
}

// NewComparableSequence is a constructor for a sequence of comparable types.
func NewComparableSequence[T cmp.Ordered](s ...[]T) *ComparableSequence[T] {
	seq := new(ComparableSequence[T])
	if len(s) == 0 {
		return seq
	}
	return &ComparableSequence[T]{Sequence[T]{elements: slices.Concat(s...)}}
}

// The following methods are mostly syntatic sugar
// wrapping Collection functions to enable function chaining:
// i.e. sequence.Filter(f).Take(n)

// Clone returns a copy of the collection. This is a shallow clone.
func (c *ComparableSequence[T]) Clone() *ComparableSequence[T] {
	return &ComparableSequence[T]{
		Sequence[T]{elements: slices.Clone(c.elements)},
	}
}

// Contains returns true if the sequence contains the given value.
func (c *ComparableSequence[T]) Contains(v T) bool {
	return slices.Contains(c.elements, v)
}

// Concat returns a new sequence concatenating the passed in sequences.
func (c *ComparableSequence[T]) Concat(sequences ...*ComparableSequence[T]) *ComparableSequence[T] {
	e := c.elements
	for _, col := range sequences {
		e = slices.Concat(e, col.elements)
	}
	return &ComparableSequence[T]{Sequence[T]{elements: e}}
}

// Concatenated is an alias for collection.Concatenated
func (c *ComparableSequence[T]) Concatenated(s *ComparableSequence[T]) iter.Seq[T] {
	return collection.Concatenated(c, s)
}

// Corresponds is an alias for collection.Corresponds
func (c *ComparableSequence[T]) Corresponds(s *ComparableSequence[T], f func(T, T) bool) bool {
	return collection.Corresponds(c, s, f)
}

// Distinct returns a new sequence containing only the unique elements from the original sequence.
func (c *ComparableSequence[T]) Distinct() *ComparableSequence[T] {
	m := make(map[T]interface{})
	r := &ComparableSequence[T]{}
	for v := range c.Values() {
		_, ok := m[v]
		if !ok {
			r.Add(v)
			m[v] = true
		}
	}
	return r
}

// Distincted is an alias for collection.Distincted
func (c *ComparableSequence[T]) Distincted() iter.Seq[T] {
	return collection.Distincted(c)
}

// Diff is an alias for collection.Diff
func (c *ComparableSequence[T]) Diff(s *ComparableSequence[T]) *ComparableSequence[T] {
	return collection.Diff(c, s).(*ComparableSequence[T])
}

// Diffed is an alias for collection.Diffed
func (c *ComparableSequence[T]) Diffed(s *ComparableSequence[T]) iter.Seq[T] {
	return collection.Diffed(c, s)
}

// Equals returns true if the two sequences are equal.
func (c *ComparableSequence[T]) Equals(c2 *ComparableSequence[T]) bool {
	return slices.Equal(c.elements, c2.elements)
}

// Exists returns true if the sequence contains the given value.
func (c *ComparableSequence[T]) Exists(v T) bool {
	return c.Contains(v)
}

// IndexOf returns the index of the first occurrence of the specified element in this sequence,
// or -1 if this sequence does not contain the element.
func (c *ComparableSequence[T]) IndexOf(v T) int {
	return slices.Index(c.elements, v)
}

// Intersect returns a new sequence containing the elements that are present in both sequences.
func (c *ComparableSequence[T]) Intersect(s *ComparableSequence[T]) *ComparableSequence[T] {
	return collection.Intersect(c, s).(*ComparableSequence[T])
}

// IntersectIterator is an alias for collection.IntersectIterator
func (c *ComparableSequence[T]) Intersected(s *ComparableSequence[T]) iter.Seq[T] {
	return collection.Intersected(c, s)
}

// LastIndexOf returns the index of the last occurrence of the specified element in this sequence,
// or -1 if this sequence does not contain the element.
func (c *ComparableSequence[T]) LastIndexOf(v T) int {
	for i, val := range c.Backward() {
		if val == v {
			return i
		}
	}
	return -1
}

// Max returns the maximum value in the sequence.
func (c *ComparableSequence[T]) Max() T {
	return slices.Max(c.elements)
}

// Min returns the minimum value in the sequence.
func (c *ComparableSequence[T]) Min() T {
	return slices.Min(c.elements)
}

// Sum returns the sum of the elements in the sequence.
func (c *ComparableSequence[T]) Sum() T {
	var sum T
	for _, v := range c.elements {
		sum += v
	}
	return sum
}

// StartsWith returns true if the sequence starts with the given sequence.
func (c *ComparableSequence[T]) StartsWith(other *ComparableSequence[T]) bool {
	return collection.StartsWith(c, other)
}

// EndsWith returns true if the sequence ends with the given sequence.
func (c *ComparableSequence[T]) EndsWith(other *ComparableSequence[T]) bool {
	return collection.EndsWith(c, other)
}
