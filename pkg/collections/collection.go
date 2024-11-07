// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package collections implements support for a generic Collection.
// A collection wraps an underlying Go slice and provides convenience methods
// and synthatic sugar on top of it.
package collections

import (
	"fmt"
	"slices"
)

// NOTE: this file only contains basic definitions.
//   - collection methods can be found in methods.go
//   - package functions can be found in functions.go
//
// it's important to keep this file lightweight

type Collection[T any] struct {
	elements []T
}

// NewCollection is a constructor for collections.
// It can be initialized with any number slices.
//
// example:
//
//	c := NewCollection([]int{1,2,3})
func NewCollection[T any](s ...[]T) Collection[T] {
	if len(s) == 0 {
		return Collection[T]{}
	}
	return Collection[T]{elements: slices.Concat(s...)}
}

// String implements the Stringer interface to enable fmt to print the underlying slice.
func (c *Collection[T]) String() string {
	return fmt.Sprintf("%v", c.elements)
}

// zero returns the zero value of the underlying slice
func (c *Collection[T]) zero() *Collection[T] {
	var z Collection[T]
	return &z
}

// zeroT returns the zero value of the underlying slice's type.
func (c *Collection[T]) zeroT() T {
	var z T
	return z
}

// collectionError type definition
type collectionError struct {
	code int
	msg  string
}

// Error implements the error interface
func (e *collectionError) Error() string {
	return fmt.Sprintf("error %d: %s", e.code, e.msg)
}

// error values used throughout this package
var (
	emptyCollectionError = &collectionError{
		code: 100, msg: "invalid operation on an empty collection",
	}
)
