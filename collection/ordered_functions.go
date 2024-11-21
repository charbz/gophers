// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// ordered_functions.go defines all the package functions that operate only on an OrderedCollection.

package collection

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
func Corresponds[T, K any](s1 OrderedCollection[T], s2 OrderedCollection[K], f func(T, K) bool) bool {
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
func Drop[T any](s OrderedCollection[T], n int) OrderedCollection[T] {
	if n <= 0 {
		return s
	} else if n >= s.Length() {
		return s.NewOrdered()
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
func DropRight[T any](s OrderedCollection[T], n int) OrderedCollection[T] {
	if n <= 0 {
		return s
	} else if n >= s.Length() {
		return s.NewOrdered()
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
func DropWhile[T any](s OrderedCollection[T], f func(T) bool) OrderedCollection[T] {
	count := 0
	for v := range s.Values() {
		if !f(v) {
			break
		}
		count++
	}
	return s.Slice(count, s.Length())
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
//	2, 3
func Find[T any](s OrderedCollection[T], f func(T) bool) (index int, value T) {
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
func FindLast[T any](s OrderedCollection[T], f func(T) bool) (index int, value T) {
	for i, v := range s.Backward() {
		if f(v) {
			return i, v
		}
	}
	return -1, *new(T)
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
func Head[T any](s OrderedCollection[T]) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	return s.At(0), nil
}

// Init returns a collection containing all elements excluding the last one.
//
// example usage:
//
//	c := NewSequence([]int{1,2,3,4,5,6})
//	c.Init()
//
// output:
//
//	[1,2,3,4,5]
func Init[T any](s OrderedCollection[T]) OrderedCollection[T] {
	if s.Length() == 0 {
		return s
	}
	return s.Slice(0, s.Length()-1)
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
func Last[T any](s OrderedCollection[T]) (T, error) {
	if s.Length() == 0 {
		return *new(T), EmptyCollectionError
	}
	return s.At(s.Length() - 1), nil
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
func ReduceRight[T, K any](s OrderedCollection[T], f func(K, T) K, init K) K {
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
func Reverse[T any](s OrderedCollection[T]) OrderedCollection[T] {
	c := s.NewOrdered()
	for _, v := range s.Backward() {
		c.Add(v)
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
func ReverseMap[T, K any](s OrderedCollection[T], f func(T) K) OrderedCollection[K] {
	r := s.NewOrdered().(OrderedCollection[K])
	for _, v := range s.Backward() {
		r.Add(f(v))
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
func SplitAt[T any](s OrderedCollection[T], n int) (OrderedCollection[T], OrderedCollection[T]) {
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
func Tail[T any](s OrderedCollection[T]) OrderedCollection[T] {
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
func Take[T any](s OrderedCollection[T], n int) OrderedCollection[T] {
	if n <= 0 {
		return s.NewOrdered()
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
func TakeRight[T any](s OrderedCollection[T], n int) OrderedCollection[T] {
	if n <= 0 {
		return s.NewOrdered()
	}
	return s.Slice(max(s.Length()-n, 0), s.Length())
}

// StartsWith checks if the elements of the second collection (s2) match the
// initial elements of the first collection (s1) in order.
//
// Example usage:
//
//	c1 := NewSequence([]int{1, 2, 3, 4, 5})
//	c2 := NewSequence([]int{1, 2})
//	StartsWith(c1, c2)
//
// Output:
//
//	true
func StartsWith[T comparable](s1 OrderedCollection[T], s2 OrderedCollection[T]) bool {
	if s1.Length() < s2.Length() {
		return false
	}

	for i, v := range s2.All() {
		if v != s1.At(i) {
			return false
		}
	}
	return true
}

// EndsWith checks if the elements of the second collection (s2) match the
// final elements of the first collection (s1) in reverse order.
//
// Example usage:
//
//	c1 := NewSequence([]int{1, 2, 3, 4, 5})
//	c2 := NewSequence([]int{4, 5})
//	EndsWith(c1, c2)
//
// Output:
//
//	true
func EndsWith[T comparable](s1 OrderedCollection[T], s2 OrderedCollection[T]) bool {
	// If s2 is longer than s1, s1 cannot end with s2
	if s1.Length() < s2.Length() {
		return false
	}

	offset := s1.Length() - s2.Length()

	for i, v := range s2.All() {
		if s1.At(offset+i) != v {
			return false
		}
	}
	return true
}
