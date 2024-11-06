// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// funcs.go defines all the collection operations that take a collection of type T
// and a higher order function as input, and produces some generic output of type K
// such as Map and Reduce.
//
// Note: unfortunately Go does not allow Generic type parameters to be defined
// on methods, therefore all operations of the form f(T) -> K must be declared
// as package functions and used as follow:
//
// Map(collection, func(t T) K {return k})

package collections

import (
	"github.com/charbz/gophers/pkg/utils"
)

// Map takes a collection of generic type T and a mapping function func(T) K
// applies the mapping function and returns a collection of type K
//
// example usage:
//
//	names := NewCollection([]string{"Alice", "Bob", "Charlie"})
//
//	Map(names, func(name string) int {
//	  return len(name)
//	})
func Map[T any, K any](s *Collection[T], f func(T) K) *Collection[K] {
	return &Collection[K]{
		utils.Map(s.elements, f),
	}
}

// Reduce takes a collection of generic type T, a reducing function func(K, T) K,
// and an initial value of type K as parameters. It applies the reducing
// function to each element and returns the resulting value K
//
// example usage:
//
//	numbers := NewCollection([]int{1,2,3,4,5,6})
//
//	Reduce(numbers, func(accumulator int, number int) int {
//	  return accumulator + number
//	}, 0)
func Reduce[T any, K any](s *Collection[T], f func(K, T) K, init K) K {
	return utils.Reduce(s.elements, f, init)
}
