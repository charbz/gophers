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

import "cmp"

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

// Corresponds tests whether every element of this sequence relates to the corresponding
// element of another sequence by satisfying a test predicate.
//
// example usage:
//
//	c1 := NewSequence([]int{1,2,3,4,5,6})
//	c2 := NewSequence([]int{2,4,6,8,10,12})
//	Corresponds(c1, c2, func(i int, j int) bool { return i*2 == j })
//
// output:
//
//	true
func Corresponds[T, K any](s1 Collection[T], s2 Collection[K], f func(T, K) bool) bool {
	if s1.Length() != s2.Length() {
		return false
	}
	for i, v := range s1.All() {
		if !f(v, s2.At(i)) {
			return false
		}
	}
	return true
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

// Drop returns a new sequence with the first n elements removed.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Drop(2)
//
// output:
//
//	[3,4,5,6]
func Drop[T any](s Collection[T], n int) Collection[T] {
	if n <= 0 {
		return s
	} else if n >= s.Length() {
		return s.New()
	}
	return s.Slice(n, s.Length())
}

// DropRight returns a sequence with the last n elements removed.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.DropRight(2)
//
// output:
//
//	[1,2,3,4]
func DropRight[T any](s Collection[T], n int) Collection[T] {
	if n <= 0 {
		return s
	} else if n >= s.Length() {
		return s.New()
	}
	return s.Slice(0, s.Length()-n)
}

// DropWhile returns a sequence with the first n elements that
// satisfy a predicate removed.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.DropWhile(func(i int) bool { return i < 4 })
//
// output:
//
//	[4,5,6]
func DropWhile[T any](s Collection[T], f func(T) bool) Collection[T] {
	count := 0
	for v := range s.Values() {
		if !f(v) {
			break
		}
		count++
	}
	return s.Slice(count, s.Length())
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

// FindLast returns the index and value of the last element
// that satisfies a predicate, otherwise returns -1 and the zero value.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	FindLast(c, func(i int) bool { return i < 6 })
//
// output:
//
//	4, 5
func FindLast[T any](s Collection[T], f func(T) bool) (index int, value T) {
	for i, v := range s.Backward() {
		if f(v) {
			return i, v
		}
	}
	return -1, *new(T)
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

// ForEach takes a function as input and applies the function to each element in the collection.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	ForEach(c, func(i int) {
//	  fmt.Println(i)
//	})
func ForEach[T any](s Collection[T], f func(T)) Collection[T] {
	for v := range s.Values() {
		f(v)
	}
	return s
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
		m[k].Append(v)
	}
	return m
}

// Head returns the first element in a Sequence and a nil error.
// If the sequence is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewSequence([]string{"A","B","C"})
//	c.Head()
//
// output:
//
//	"A", nil
func Head[T any](s Collection[T]) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	return s.At(0), nil
}

// Init returns a sequence containing all elements excluding the last one.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Init()
//
// output:
//
//	[1,2,3,4,5]
func Init[T any](s Collection[T]) Collection[T] {
	if s.Length() == 0 {
		return s
	}
	return s.Slice(0, s.Length()-1)
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

// Last returns the last element in the Sequence and a nil error.
// If the sequence is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewSequence([]string{"A","B","C"})
//	c.Last()
//
// output:
//
//	"C", nil
func Last[T any](s Collection[T]) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	return s.At(s.Length() - 1), nil
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
	maxValue := f(s.At(0))
	maxElement := s.At(0)
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
	minValue := f(s.At(0))
	minElement := s.At(0)
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

// ReduceRight takes a collection of type T, a reducing function func(K, T) K,
// and an initial value of type K as parameters. It applies the reducing
// function to each element in reverse order and returns the resulting value K.
//
// example usage:
//
//	c := NewSequence([]string{"A","B","C"})
//	ReduceRight(c, func(acc string, i string) string { return acc + i }, "")
//
// output:
//
//	"CBA"
func ReduceRight[T, K any](s Collection[T], f func(K, T) K, init K) K {
	accumulator := init
	for _, v := range s.Backward() {
		accumulator = f(accumulator, v)
	}
	return accumulator
}

// Reverse returns a new sequence with all elements in reverse order.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Reverse()
//
// output:
//
//	[6,5,4,3,2,1]
func Reverse[T any](s Collection[T]) Collection[T] {
	c := s.New()
	for _, v := range s.Backward() {
		c.Append(v)
	}
	return c
}

// ReverseMap takes a collection of type T and a mapping function func(T) K,
// applies the mapping function to each element in reverseand returns a collection of type K.
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
//	[6, 3, 5]
func ReverseMap[T, K any](s Collection[T], f func(T) K) Collection[K] {
	r := s.New().(Collection[K])
	for _, v := range s.Backward() {
		r.Append(f(v))
	}
	return r
}

// SplitAt returns two new sequences containing the first n elements and the rest of the elements.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	SplitAt(c, 3)
//
// output:
//
//	[1,2,3], [4,5,6]
func SplitAt[T any](s Collection[T], n int) (Collection[T], Collection[T]) {
	return s.Slice(0, n), s.Slice(n, s.Length())
}

// Tail returns a new sequence containing all elements excluding the first one.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Tail()
//
// output:
//
//	[2,3,4,5,6]
func Tail[T any](s Collection[T]) Collection[T] {
	if s.Length() == 0 {
		return s
	}
	return s.Slice(1, s.Length())
}

// Take returns a new sequence containing the first n elements.
//
// example usage:
//
//	[c := NewSequence([]int{1,2,3,4,5,6})
//	c.Take(3)
//
// output:
//
//	[1,2,3]
func Take[T any](s Collection[T], n int) Collection[T] {
	if n <= 0 {
		return s.New()
	}
	return s.Slice(0, min(n, s.Length()))
}

// TakeRight returns a new sequence containing the last n elements.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.TakeRight(3)
//
// output:
//
//	[4,5,6]
func TakeRight[T any](s Collection[T], n int) Collection[T] {
	if n <= 0 {
		return s.New()
	}
	return s.Slice(max(s.Length()-n, 0), s.Length())
}
