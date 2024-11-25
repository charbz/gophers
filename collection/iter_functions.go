// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// iter_functions implements functions that take a collection as input, and
// return an iterator to the result instead of a new collection.

package collection

import "iter"

// Concatenated returns an iterator that yields the elements of s1 and s2.
//
// example usage:
//
//	a := NewList([]int{1,2})
//	b := NewList([]int{3,4})
//	for v := range Concatenated(a, b) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	2
//	3
//	4
func Concatenated[T any](s1, s2 Collection[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s1.Values() {
			yield(v)
		}
		for v := range s2.Values() {
			yield(v)
		}
	}
}

// Diffed returns an iterator that yields the elements of s1 that are not present in s2.
//
// example usage:
//
//	a := NewList([]int{1,2,3,4,5,6})
//	b := NewList([]int{2,4,6,8,10,12})
//	for v := range Diffed(a, b) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	3
//	5
func Diffed[T comparable](s1 OrderedCollection[T], s2 OrderedCollection[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s1.Values() {
			i, _ := Find(s2, func(t T) bool { return t == v })
			if i == -1 {
				yield(v)
			}
		}
	}
}

// DiffedFunc is similar to Diffed but applies to non-comparable types.
// It takes two collections (s1, s2) and an "equality" function as an argument such as
// func(a T, b T) bool {return a == b}
// and returns an iterator that yields the elements of s1 that are not present in s2.
//
// example usage:
//
//	a := NewList([]int{1,2,3,4,5,6})
//	b := NewList([]int{2,4,6,8,10,12})
//	for v := range DiffedFunc(a, b, func(a int, b int) bool { return a == b }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	3
//	5
func DiffedFunc[T any](s1 OrderedCollection[T], s2 OrderedCollection[T], f func(T, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s1.Values() {
			i, _ := Find(s2, func(t T) bool { return f(v, t) })
			if i == -1 {
				yield(v)
			}
		}
	}
}

// Distincted returns an iterator that yields the unique elements of s.
//
// example usage:
//
//	a := NewList([]int{1,1,1,2,2,3})
//	for v := range Distincted(a) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	2
//	3
func Distincted[T comparable](s Collection[T]) iter.Seq[T] {
	seen := make(map[T]bool)
	return func(yield func(T) bool) {
		for v := range s.Values() {
			if !seen[v] {
				seen[v] = true
				yield(v)
			}
		}
	}
}

// DistinctedFunc is similar to Distincted but applies to non-comparable types.
// It takes a collection (s) and an "equality" function as an argument such as
// func(a T, b T) bool {return a == b}
// and returns an iterator that yields the unique elements of s.
//
// example usage:
//
//	a := NewList([]int{1,1,1,2,2,3})
//	for v := range DistinctedFunc(a, func(a int, b int) bool { return a == b }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	2
//	3
func DistinctedFunc[T any](s Collection[T], f func(T, T) bool) iter.Seq[T] {
	s2 := s.New()
	return func(yield func(T) bool) {
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
				yield(v)
			}
		}
	}
}

// Filtered returns an iterator that yields the elements of s
// that satisfy the predicate function f.
//
// example usage:
//
//	a := NewList([]int{1,2,3,4,5,6})
//	for v := range Filtered(a, func(i int) bool { return i % 2 == 0 }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	2
//	4
//	6
func Filtered[T any](s Collection[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.Values() {
			if f(v) {
				yield(v)
			}
		}
	}
}

// Intersected returns an iterator that yields the elements of s1
// that are also present in s2.
//
// example usage:
//
//	a := NewList([]int{1,3,4,5,6})
//	b := NewList([]int{2,4,6,8,10,12})
//	for v := range Intersected(a, b) {
//		fmt.Println(v)
//	}
//
// output:
//
//	4
//	6
func Intersected[T comparable](s1 Collection[T], s2 Collection[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s1.Values() {
			for v2 := range s2.Values() {
				if v == v2 {
					yield(v)
				}
			}
		}
	}
}

// IntersectedFunc is similar to Intersected but applies to non-comparable types.
// It takes two collections (s1, s2) and an "equality" function as an argument such as
// func(a T, b T) bool {return a == b}
// and returns an iterator that yields the elements of s1 that are also present in s2.
//
// example usage:
//
//	a := NewList([]int{1,3,4,5,6})
//	b := NewList([]int{2,4,6,8,10,12})
//	for v := range IntersectedFunc(a, b, func(a int, b int) bool { return a == b }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	4
//	6
func IntersectedFunc[T any](s1 Collection[T], s2 Collection[T], f func(T, T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s1.Values() {
			for v2 := range s2.Values() {
				if f(v, v2) {
					yield(v)
				}
			}
		}
	}
}

// Mapped returns an iterator that yields the elements of s
// transformed by the function f.
//
// example usage:
//
//	a := NewList([]int{1,2,3})
//	for v := range Mapped(a, func(i int) int { return i * 2 }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	2
//	4
//	6
func Mapped[T, K any](s Collection[T], f func(T) K) iter.Seq[K] {
	return func(yield func(K) bool) {
		for v := range s.Values() {
			yield(f(v))
		}
	}
}

// Rejected returns an iterator that yields the elements of s
// that do not satisfy the predicate function f.
//
// example usage:
//
//	a := NewList([]int{1,2,3,4,5,6})
//	for v := range Rejected(a, func(i int) bool { return i % 2 == 0 }) {
//		fmt.Println(v)
//	}
//
// output:
//
//	1
//	3
//	5
func Rejected[T any](s Collection[T], f func(T) bool) iter.Seq[T] {
	return Filtered(s, func(t T) bool { return !f(t) })
}
