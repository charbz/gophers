// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// wrappers.go contains collection methods that wrap existing functions
// provided by the go slices library to provide syntatic sugar and enable
// function chaining, for example:
//
// collection.Reverse().Concat(c2).Sort()

package collections

import (
	"iter"
	"slices"
)

func (c *Collection[T]) All() iter.Seq2[int, T] {
	return slices.All(c.elements)
}

func (c *Collection[T]) Values() iter.Seq[T] {
	return slices.Values(c.elements)
}
