// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package collection

import "iter"

// OrderedCollection is a generic interface for collections who's underlying
// data structure is index-based, and the order of elements matters.
type OrderedCollection[T any] interface {
	Collection[T]
	At(index int) T
	All() iter.Seq2[int, T]
	Backward() iter.Seq2[int, T]
	Slice(start, end int) OrderedCollection[T]
	NewOrdered(s ...[]T) OrderedCollection[T]
}
