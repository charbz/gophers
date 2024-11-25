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

// Contains returns true if the sequence contains the given value.
func (c *ComparableSequence[T]) Contains(v T) bool {
	return slices.Contains(c.elements, v)
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

// DistinctIterator is an alias for collection.DistinctIterator
func (c *ComparableSequence[T]) DistinctIterator() iter.Seq[T] {
	return collection.DistinctIterator(c)
}

// Diff is an alias for collection.Diff
func (c *ComparableSequence[T]) Diff(s *ComparableSequence[T]) *ComparableSequence[T] {
	return collection.Diff(c, s).(*ComparableSequence[T])
}

// DiffIterator is an alias for collection.DiffIterator
func (c *ComparableSequence[T]) DiffIterator(s *ComparableSequence[T]) iter.Seq[T] {
	return collection.DiffIterator(c, s)
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
func (c *ComparableSequence[T]) IntersectIterator(s *ComparableSequence[T]) iter.Seq[T] {
	return collection.IntersectIterator(c, s)
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
