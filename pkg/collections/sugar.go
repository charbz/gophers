// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// sugar.go this file just contains synthatic sugar for the Go slices package,
// the methods defined here enable function chaining i.e.:
//
//   collection.Reverse().Concat(c2).All()

package collections

import (
	"iter"
	"slices"
)

// All returns an interator over all (index,value) pairs of the underlying slice.
func (c *Collection[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

// Backward returns an iterator over index-value pairs in the collection,
// traversing it backward with descending indices.
func (c *Collection[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(c.elements)
}

// Clone returns a copy of the collection. The elements are copied using assignment,
// so this is a shallow clone.
func (c *Collection[T]) Clone() *Collection[T] {
	return &Collection[T]{
		slices.Clone(c.elements),
	}
}

// Concat returns a new collection concatenating the passed in collections.
func (c *Collection[T]) Concat(collections ...Collection[T]) *Collection[T] {
	e := c.elements
	for _, col := range collections {
		e = slices.Concat(e, col.elements)
	}
	return &Collection[T]{e}
}

// Distinct takes a higher order "equality" function as an argument
// and returns a new collection containing all the unique elements
// from the original collection.
//
// example:
//
//	c := NewCollection([]int{1,2,2,3,3,3})
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
func (c *Collection[T]) Distinct(f func(T, T) bool) *Collection[T] {
	return &Collection[T]{
		slices.CompactFunc(c.elements, f),
	}
}

// Values returns an iterator over all values of the underlying slice.
func (c *Collection[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}
