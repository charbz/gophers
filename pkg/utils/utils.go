// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// The utils package contains several helpful utilities for manipulating slices
// that are not part of the standard slices library.
// This includes Filter, Map, Reduce, and others.

package utils

import (
	"slices"
)

// Distinct returns a new slice containing only unique elements from the input slice.
// The input type must be comparable to enable map-based deduplication.
func Distinct[S ~[]T, T comparable](s S) S {
	m := make(map[T]interface{})
	r := make([]T, 0)
	for v := range slices.Values(s) {
		_, ok := m[v]
		if !ok {
			r = append(r, v)
			m[v] = true
		}
	}
	return r
}

// Filter returns a new slice containing only the elements that satisfy the predicate function f.
// The predicate f returns true for elements that should be included in the result.
func Filter[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// FilterNot returns a new slice containing only the elements that do not satisfy the predicate function f.
// The predicate f returns true for elements that should be excluded from the result.
func FilterNot[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Map applies the function f to each element in the slice and returns a new slice
// containing the results. The function f transforms elements of type T to type K.
func Map[T any, K any](s []T, f func(T) K) []K {
	s2 := make([]K, 0, len(s))
	for v := range slices.Values(s) {
		s2 = append(s2, f(v))
	}
	return s2
}

// Partition splits the input slice into two slices based on the predicate function f.
// The first returned slice contains elements for which f returns true,
// and the second contains elements for which f returns false.
func Partition[S ~[]T, T any](s S, f func(T) bool) (S, S) {
	match := make([]T, 0)
	noMatch := make([]T, 0)
	for v := range slices.Values(s) {
		if f(v) {
			match = append(match, v)
		} else {
			noMatch = append(noMatch, v)
		}
	}
	return match, noMatch
}

// Reduce applies the reduction function f to each element in the slice,
// accumulating a single result. The init parameter provides the initial value
// for the accumulator.
func Reduce[S ~[]T, T any, K any](s S, f func(K, T) K, init K) K {
	acc := init
	for v := range slices.Values(s) {
		acc = f(acc, v)
	}
	return acc
}

// Find returns the index and value of the first element in the slice that satisfies
// the predicate function f. If no element is found, returns -1 and the zero value
// of type T.
func Find[S ~[]T, T any](s S, f func(T) bool) (index int, value T) {
	for i, v := range slices.All(s) {
		if f(v) {
			return i, v
		}
	}
	return -1, *new(T)
}
