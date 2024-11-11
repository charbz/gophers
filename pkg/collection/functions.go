// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// functions.go defines all the package functions that operate on a Collection but could not
// be defined as methods on the Collection struct due to the limitation of Generics in Go.
//
// Unfortunately Go does not allow Generic type parameters to be defined directly on struct methods,
// Given that the Collection struct is bound to 1 generic argument [T any] representing the underlying type,
// operations that operations that map into a different type altogether such as f(T) -> K must be defined as functions.
// and used as follows:
//
//	Map(collection, func(t T) K {
//	     ...
//	     return k
//	})

package collection

// Filter returns a new collection containing only the elements that
// satisfy the predicate function.
//
// example usage:
//
//	numbers := NewSequence([]int{1,2,3,4,5,6})
//	Filter(numbers, func(t int) bool { return t % 2 == 0 })
//
// output:
//
//	[2,4,6]
func Filter[T any](s Collection[T], f func(T) bool) Collection[T] {
	result := s.New()
	for v := range s.Values() {
		if f(v) {
			result.Append(v)
		}
	}
	return result
}

// FilterNot returns the complement of the Filter function.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	FilterNot(c, func(t int) bool { return t % 2 == 0 })
//
// output:
//
//	[1,3,5]
func FilterNot[T any](s Collection[T], f func(T) bool) Collection[T] {
	return Filter(s, func(t T) bool { return !f(t) })
}

// Find returns the index and value of the first element
// that satisfies a predicate, otherwise returns -1 and the zero value.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	Find(c, func(i int) bool {
//	  return (i + 3) > 5
//	})
//
// output
//
//	3
func Find[T any](s Collection[T], f func(T) bool) (index int, value T) {
	for i, v := range s.All() {
		if f(v) {
			return i, v
		}
	}
	return -1, *new(T)
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
func Map[T, K any](s Collection[T], f func(T) K) Collection[K] {
	r := s.New().(Collection[K])
	for v := range s.Values() {
		r.Append(f(v))
	}
	return r
}

// Partition takes a partitioning function as input and returns two collections,
// the first one contains the elements that match the partitioning condition,
// the second one contains the rest of the elements.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	Partition(c, func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[2,4,6], [1,3,5]
func Partition[T any](s Collection[T], f func(T) bool) (Collection[T], Collection[T]) {
	match := s.New()
	noMatch := s.New()
	for v := range s.Values() {
		if f(v) {
			match.Append(v)
		} else {
			noMatch.Append(v)
		}
	}
	return match, noMatch
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
func Reduce[T, K any](s Collection[T], f func(K, T) K, init K) K {
	accumulator := init
	for v := range s.Values() {
		accumulator = f(accumulator, v)
	}
	return accumulator
}
