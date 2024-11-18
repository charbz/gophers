// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// functions.go defines all the package functions that operate on a Collection.
// These functions apply to both Collection and OrderedCollection types.

// Unfortunately Go does not allow Generic type parameters to be defined directly on struct methods,
// Given that the Collection struct is bound to 1 generic argument [T any] representing the underlying type,
// operations that map into a different type altogether such as f(T) -> K must be defined as functions.
// and used as follows:
//
//	Map(collection, func(t T) K {
//	     ...
//	     return k
//	})

package collection

import (
	"cmp"
)

// Count returns the number of elements in the collection that satisfy the predicate function.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	Count(c, func(i int) bool { return i % 2 == 0 })
//
// output:
//
//	3
func Count[T any](s Collection[T], f func(T) bool) int {
	count := 0
	for v := range s.Values() {
		if f(v) {
			count++
		}
	}
	return count
}

// Diff returns a new collection containing elements that are present in the first collection but not in the second.
//
// example usage:
//
//	c1 := NewSequence([]int{1,2,3,4,5,6})
//	c2 := NewSequence([]int{2,4,6,8,10,12})
//	Diff(c1, c2)
//
// output:
//
//	[1,3,5]
func Diff[T comparable](s1 Collection[T], s2 Collection[T]) Collection[T] {
	return FilterNot(s1, func(t T) bool {
		for v := range s2.Values() {
			if v == t {
				return true
			}
		}
		return false
	})
}

// Distinct returns a new collection containing only the unique elements of the collection.
//
// example usage:
//
//	c := NewSequence([]int{1,1,1,4,5,1,2,2})
//	Distinct(c, func(i int, i2 int) bool { return i == i2 })
//
// output:
//
//	[1,4,5,2]
func Distinct[T any](s Collection[T], f func(T, T) bool) Collection[T] {
	s2 := s.New()
	for v := range s.Values() {
		match := false
		for v2 := range s2.Values() {
			if f(v, v2) {
				match = true
				break
			}
		}
		if !match {
			s2.Add(v)
		}
	}
	return s2
}

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
			result.Add(v)
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

// ForAll tests whether a predicate holds for all elements of this sequence.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	ForAll(c, func(i int) bool { return i < 10 })
//
// output:
//
//	true
func ForAll[T any](s Collection[T], f func(T) bool) bool {
	for v := range s.Values() {
		if !f(v) {
			return false
		}
	}
	return true
}

// GroupBy takes a collection and a grouping function as input and returns a map
// where the key is the result of the grouping function and the value is a collection
// of elements that satisfy the predicate.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	GroupBy(c, func(i int) int { return i % 2 })
//
// output:
//
//	{0:[2,4,6], 1:[1,3,5]}
func GroupBy[T any, K comparable](s Collection[T], f func(T) K) map[K]Collection[T] {
	m := make(map[K]Collection[T])
	for v := range s.Values() {
		k := f(v)
		if _, ok := m[k]; !ok {
			m[k] = s.New()
		}
		m[k].Add(v)
	}
	return m
}

// Intersect returns a new collection containing elements that are present in both input collections.
//
// example usage:
//
//	c1 := NewSequence([]int{1,2,3,4,5,6})
//	c2 := NewSequence([]int{2,4,6,8,10,12})
//	Intersect(c1, c2)
//
// output:
//
//	[2,4,6]
func Intersect[T comparable](s1 Collection[T], s2 Collection[T]) Collection[T] {
	return Filter(s1, func(t T) bool {
		for v := range s2.Values() {
			if v == t {
				return true
			}
		}
		return false
	})
}

// Map takes a collection of type T and a mapping function func(T) K,
// applies the mapping function to each element and returns a slice of type K.
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
func Map[T, K any](s Collection[T], f func(T) K) []K {
	k := make([]K, 0, s.Length())
	for v := range s.Values() {
		k = append(k, f(v))
	}
	return k
}

// MaxBy returns the element in the collection that has the maximum value
// according to a comparison function.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	MaxBy(c, func(a int, b int) int { return a - b })
//
// output:
//
//	6
func MaxBy[T any, K cmp.Ordered](s Collection[T], f func(T) K) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	maxElement := s.Random()
	maxValue := f(maxElement)
	for v := range s.Values() {
		if f(v) > maxValue {
			maxElement = v
			maxValue = f(v)
		}
	}
	return maxElement, nil
}

// MinBy returns the element in the collection that has the minimum value
// according to a comparison function.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	MinBy(c, func(a int, b int) int { return a - b })
//
// output:
//
//	1
func MinBy[T any, K cmp.Ordered](s Collection[T], f func(T) K) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	minElement := s.Random()
	minValue := f(minElement)
	for v := range s.Values() {
		if f(v) < minValue {
			minElement = v
			minValue = f(v)
		}
	}
	return minElement, nil
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
			match.Add(v)
		} else {
			noMatch.Add(v)
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
