// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// functions.go defines all the package functions that operate on a Collection but could not
// be defined as methods on the Collection struct due to the limitation of Generics in Go.
//
// Unfortunately Go does not allow Generic type parameters to be defined directly on struct methods,
// Given that the Collection struct is bound to 1 generic argument [T any] representing the underlying type,
// operations that require a specific constraint, such as [T comparable], and operations that map into a
// different type altogether such as f(T) -> K must be defined as functions. and used as follows:
//
//	Map(collection, func(t T) K {
//	     ...
//	     return k
//	})

package collections

import (
	"github.com/charbz/gophers/pkg/utils"
)

// Distinct takes a collection of type T comparable, and returns a new collection
// containing all the unique elements from the original collection.
//
// example:
//
//	c := NewCollection([]int{1,2,2,2,3,3})
//	Distinct(c)
//
// output:
//
//	[1,2,3]
func Distinct[T comparable](s *Collection[T]) *Collection[T] {
	return &Collection[T]{
		utils.Distinct(s.elements),
	}
}

// Map takes a collection of type T and a mapping function func(T) K,
// applies the mapping function to each element and returns a collection of type K.
//
// example usage:
//
//	names := NewCollection([]string{"Alice", "Bob", "Charlie"})
//	Map(names, func(name string) int {
//	  return len(name)
//	})
//
// output:
//
//	[5,3,6]
func Map[T, K any](s *Collection[T], f func(T) K) *Collection[K] {
	return &Collection[K]{
		utils.Map(s.elements, f),
	}
}

// Reduce takes a collection of type T, a reducing function func(K, T) K,
// and an initial value of type K as parameters. It applies the reducing
// function to each element and returns the resulting value K.
//
// example usage:
//
//	numbers := NewCollection([]int{1,2,3,4,5,6})
//
//	Reduce(numbers, func(accumulator int, number int) int {
//	  return accumulator + number
//	}, 0)
//
// output:
//
//	21
func Reduce[T, K any](s *Collection[T], f func(K, T) K, init K) K {
	return utils.Reduce(s.elements, f, init)
}
