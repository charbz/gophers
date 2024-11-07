// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// wrappers.go contains collection methods that wrap functions
// provided by the Go slices library. This file is just synthatic sugar to
// enable function chaining i.e.:
//
//    collection.Reverse().Concat(c2).Sort()

package collections

import (
	"iter"
	"slices"
)

// All returns an interator over all (index,value) pairs of the underlying slice.
func (c *Collection[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

// Values returns an iterator over all values of the underlying slice.
func (c *Collection[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}
