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

func Filter[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func FilterNot[S ~[]T, T any](s S, f func(T) bool) S {
	r := make([]T, 0, len(s))
	for v := range slices.Values(s) {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Map[T any, K any](s []T, f func(T) K) []K {
	s2 := make([]K, 0, len(s))
	for v := range slices.Values(s) {
		s2 = append(s2, f(v))
	}
	return s2
}

func Reduce[S ~[]T, T any, K any](s S, f func(K, T) K, init K) K {
	acc := init
	for v := range slices.Values(s) {
		acc = f(acc, v)
	}
	return acc
}

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
