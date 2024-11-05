// The utils package contains several helpful utilities for manipulating slices.
// These include common operations such as Filter, Map, and Reduce.
package utils

import (
	"slices"
)

// Filter takes a slice of any type T and a filtering function as parameters.
// It returns a new slice containing all the elements from the original slice that satisfy the filter.
// Note that the provided function must take an element T as input and return true if T satisfies the filter.
func Filter[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// FilterNot takes a slice of any type T and a filtering function as parameters.
// It returns a new slice containing all the elements from the original slice that do not satisfy the filter condition.
// Note that the provided function must take an element T as input and return true if T satisfies the filter condition.
func FilterNot[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

// MapFunc takes a slice of type T, and a mapping function as parameters.
// It returns a new slice of type K with each value from []T mapped into a new value in []K
// Note that the provided function must take an element T as input and produce K.
func Map[T any, K any](s []T, f func(T) K) []K {
	s2 := make([]K, 0, len(s))
	for v := range slices.Values(s) {
		s2 = append(s2, f(v))
	}
	return s2
}

// ReduceFunc takes a slice of type T, a reduction function, and an initial value K as parameters.
// It returns the accumulated value we get by applying the reduction function to every element of the slice.
// Note that the provided function must take an accumulator value of type K, and an element of type T
// and returns the output of combining K with T
func Reduce[S ~[]T, T any, K any](s S, f func(K, T) K, init K) K {
	acc := init
	for v := range slices.Values(s) {
		acc = f(acc, v)
	}
	return acc
}

// Partition takes a slice of type T, and a partition function as parameters.
// It returns 2 new slices of type T, the first one contains all the elements that satisfy the partition condition,
// The second list contains the elements that do not satisfy the partition condition.
// Note that the provided function must take an element T as input and return true if T satisifes the partition condition.
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

// Distinct takes a slice of type T comparable as a parameter
// and returns a new slice with all the distinct elements
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
