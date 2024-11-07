// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// functions.go defines all the package functions that take a collection of type T as input
// and return some generic type K as output. Such as Map and Reduce.
//
// Unfortunately Go does not allow Generic type parameters to be defined directly on methods,
// otherwise all the code in here would be moved to methods.go. therefore all operations of
// the form f(T) -> K must be declared here and used as follows:
//
//	Map(collection, func(t T) K {
//	     ...
//	     return k
//	})

package collections

import (
	"github.com/charbz/gophers/pkg/utils"
)

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
func Map[T any, K any](s *Collection[T], f func(T) K) *Collection[K] {
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
func Reduce[T any, K any](s *Collection[T], f func(K, T) K, init K) K {
	return utils.Reduce(s.elements, f, init)
}
