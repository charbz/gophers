// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package collections implements support for generic Collections of data.
// A collection wraps an underlying Go slice and provides convenience methods
// and synthatic sugar on top of it.
package collection

import (
	"fmt"
	"iter"
)

// Collection is a generic interface that must be implemented by all collection sub-types.
// At a minimum, collections must support the methods defined below.
type Collection[T any] interface {
	Append(T)
	Length() int
	New(s ...[]T) Collection[T]
	Random() T
	Values() iter.Seq[T]
}

type CollectionError struct {
	code int
	msg  string
}

func (e *CollectionError) Error() string {
	return fmt.Sprintf("error %d: %s", e.code, e.msg)
}

var (
	EmptyCollectionError = &CollectionError{
		code: 100, msg: "invalid operation on an empty collection",
	}
	ValueNotFoundError = &CollectionError{
		code: 101, msg: "value not found",
	}
	IndexOutOfBoundsError = &CollectionError{
		code: 102, msg: "index out of bounds",
	}
)
