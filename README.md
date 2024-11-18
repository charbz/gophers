# Gophers - The generic collections library for Go

Gophers is an awesome collections library for Go offering tons of functionality right out of the box.

Gophers offers the following collections:
- Sequence
- ComparableSequence
- List
- ComparableList
- Set

## Installation
```bash
go get github.com/charbz/gophers
```

## Quick Start

### Using Generic Data Types

Here are some examples of what you can do:

```go
import (
  "github.com/charbz/gophers/pkg/list"
)

type Foo struct {
  a int
  b string
}

foos := list.NewList([]Foo{
  {a: 1, b: "one"}, 
  {a: 2, b: "two"}, 
  {a: 3, b: "three"}, 
  {a: 4, b: "four"}, 
  {a: 5, b: "five"},
})

foos.Filter(func(f Foo) bool { return f.a%2 == 0 }) 
// List[Foo] {{2 two} {4 four}}

foos.FilterNot(func(f Foo) bool { return f.a%2 == 0 }) 
// List[Foo] {{1 one} {3 three} {5 five}}

foos.Find(func(f Foo) bool { return f.a == 3 }) 
// {a: 3, b: "three"}

foos.Partition(func(f Foo) bool { return len(f.b) == 3 })
// List[Foo] {{1 one} {2 two}} , List[Foo] {{3 three} {4 four} {5 five}}

foos.SplitAt(3) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four}} , List[Foo] {{5 five}}

foos.Count(func(f Foo) bool { return f.a < 3 }) 
// 2

bars := foos.Concat(list.NewList([]Foo{{a: 1, b: "one"}, {a: 2, b: "two"}})) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four} {5 five} {1 one} {2 two}}

bars.Distinct(func(i Foo, j Foo) bool { return i.a == j.a }) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four} {5 five}}

foos.Apply(
  func(f Foo) Foo {
    f.a *= 2
    f.b += " * two"
    return f
  }
)
// List[Foo] {{2 one * two} {4 two * two} {6 three * two} {8 four * two} {10 five * two}}
```

### Comparable Collections

Comparable collections are collections with elements that can be compared to each other.

```go
import (
  "github.com/charbz/gophers/pkg/list"
)

nums := list.NewComparableList([]int{1, 1, 2, 2, 2, 3, 4, 5})

nums.Max() // 5

nums.Min() // 1

nums.Sum() // 20

nums.Distinct() // List[int] {1,2,3,4,5}

nums.Reverse() // List[int] {5,4,3,2,2,1,1}

nums.SplitAt(3) // List[int] {1,1,2,2}, List[int] {2,3,4,5}

nums.Take(3) // List[int] {1,1,2}

nums.TakeRight(3) // List[int] {3,4,5}

nums.Drop(3) // List[int] {2,2,3,4,5}

nums.DropRight(3) // List[int] {1,1,2,2,2}

nums.DropWhile(
  func(i int) bool { return i < 3 }, // List[int] {3,4,5}
)

nums.Diff(
  list.NewComparableList([]int{1, 2, 3}), // List[int] {4,5}
)

nums.Count(
  func(i int) bool { return i > 2 }, // 3
)
```

### Sets

Sets are collections of unique elements. Sets also implement the Collection interface, and offer additional methods for set operations.

```go
import (
  "github.com/charbz/gophers/pkg/set"
)

setA := set.NewSet([]string{"A", "B", "C", "A", "C", "A"}) // Set[string] {"A", "B", "C"}
setB := set.NewSet([]string{"A", "B", "D"})

setA.Intersection(setB) // Set[string] {"A", "B"}

setA.Union(setB) // Set[string] {"A", "B", "C", "D"}

setA.Diff(setB) // Set[string] {"C"}

setA.Apply(
  func(s string) string {
    s += "!"
    return s
  }, // Set[string] {"A!", "B!", "C!"}
)
```

### Map, Reduce, GroupBy...

You can use package functions such as Map, Reduce, GroupBy, etc on any concrete collection type.

```go
import (
  "github.com/charbz/gophers/pkg/collection"
  "github.com/charbz/gophers/pkg/list"
  "github.com/charbz/gophers/pkg/sequence"
)

foos := sequence.NewSequence([]Foo{
	{a: 1, b: "one"},
	{a: 2, b: "two"},
	{a: 3, b: "three"},
	{a: 4, b: "four"},
	{a: 5, b: "five"},
})

collection.Map(foos, func(f Foo) string { return f.b }) //  ["one", "two", "three", "four", "five"] 

collection.Reduce(foos, func(acc string, f Foo) string { return acc + f.b }, "") // "onetwothreefourfive"

collection.Reduce(foos, func(acc int, f Foo) int { return acc + f.a }, 0) // 15

collection.GroupBy(foos, func(f Foo) int { return f.a % 2 }) // Map[int][]Foo { 0: [{2 two}, {4 four}], 1: [{1 one}, {3 three}, {5 five}]}
```

## Core Features

- **Collection** : A generic collection interface providing common operations for all concrete collections.
- **Sequence** : An ordered collection wrapping a Go slice. Great for fast random access.
- **ComparableSequence** : An ordered collection with elements that can be compared to each other.
- **List** : An ordered collection wrapping a linked list. Great for fast insertion and removal, implementing queues and stacks.
- **ComparableList** : An ordered collection with elements that can be compared to each other.
- **Set** : A hash set implementation.

### Sequence Operations

- `Add(element)` - Append element to sequence
- `All()` - Get iterator over all elements
- `At(index)` - Get element at index
- `Apply(function)` - Apply function to each element
- `Backward()` - Get reverse iterator over elements
- `Clone()` - Create shallow copy of sequence
- `Concat(sequences...)` - Concatenate multiple sequences
- `Contains(predicate)` - Test if any element matches predicate
- `Corresponds(sequence, function)` - Test element-wise correspondence
- `Count(predicate)` - Count elements matching predicate
- `Dequeue()` - Remove and return first element
- `Distinct(function)` - Get unique elements using equality function
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `DropWhile(predicate)` - Drop elements while predicate is true
- `Enqueue(element)` - Add element to end
- `Equals(sequence, function)` - Test sequence equality using function
- `Exists(predicate)` - Test if any element matches predicate
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `Find(predicate)` - Find first matching element
- `FindLast(predicate)` - Find last matching element
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Head()` - Get first element
- `Init()` - Get all elements except last
- `IsEmpty()` - Test if sequence is empty
- `Last()` - Get last element
- `Length()` - Get number of elements
- `New(slices...)` - Create new sequence
- `NewOrdered(slices...)` - Create new ordered sequence
- `NonEmpty()` - Test if sequence is not empty
- `Partition(predicate)` - Split sequence based on predicate
- `Pop()` - Remove and return last element
- `Push(element)` - Add element to end
- `Random()` - Get random element
- `Reverse()` - Reverse order of elements
- `Slice(start, end)` - Get subsequence from start to end
- `SplitAt(n)` - Split sequence at index n
- `String()` - Get string representation
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Tail()` - Get all elements except first
- `ToSlice()` - Convert to Go slice
- `Values()` - Get iterator over values

### ComparableSequence Operations

Inherits all operations from Sequence, but with the following additional operations:

- `Contains(element)` - Test if sequence contains element
- `Distinct()` - Get unique elements using equality comparison
- `Diff(sequence)` - Get elements in first sequence but not in second
- `Equals(sequence)` - Test sequence equality using equality comparison
- `Exists(element)` - Test if sequence contains element
- `IndexOf(element)` - Get index of first occurrence of element
- `LastIndexOf(element)` - Get index of last occurrence of element
- `Max()` - Get maximum element
- `Min()` - Get minimum element
- `Sum()` - Get sum of all elements

### List Operations

- `Add(element)` - Add element to end
- `All()` - Get iterator over index/value pairs
- `Apply(function)` - Apply function to each element
- `At(index)` - Get element at index
- `Backward()` - Get reverse iterator over index/value pairs
- `Clone()` - Create shallow copy
- `Concat(lists...)` - Concatenate multiple lists
- `Contains(predicate)` - Test if any element matches predicate
- `Corresponds(list, function)` - Test element-wise correspondence
- `Count(predicate)` - Count elements matching predicate
- `Dequeue()` - Remove and return first element
- `Distinct(function)` - Get unique elements using equality function
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `DropWhile(predicate)` - Drop elements while predicate is true
- `Enqueue(element)` - Add element to end
- `Equals(list, function)` - Test list equality using function
- `Exists(predicate)` - Test if any element matches predicate
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `Find(predicate)` - Find first matching element
- `FindLast(predicate)` - Find last matching element
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Head()` - Get first element
- `Init()` - Get all elements except last
- `IsEmpty()` - Test if list is empty
- `Last()` - Get last element
- `Length()` - Get number of elements
- `New(slices...)` - Create new list
- `NewOrdered(slices...)` - Create new ordered list
- `NonEmpty()` - Test if list is not empty
- `Partition(predicate)` - Split list based on predicate
- `Pop()` - Remove and return last element
- `Push(element)` - Add element to end
- `Random()` - Get random element
- `Reverse()` - Reverse order of elements
- `Slice(start, end)` - Get sublist from start to end
- `SplitAt(n)` - Split list at index n
- `String()` - Get string representation
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Tail()` - Get all elements except first
- `ToSlice()` - Convert to Go slice
- `Values()` - Get iterator over values

### ComparableList Operations

Inherits all operations from List, but with the following additional operations:

- `Contains(value)` - Test if list contains value
- `Distinct()` - Get unique elements
- `Diff(list)` - Get elements in first list but not in second
- `Exists(value)` - Test if list contains value (alias for Contains)
- `Equals(list)` - Test list equality
- `IndexOf(value)` - Get index of first occurrence of value
- `LastIndexOf(value)` - Get index of last occurrence of value
- `Max()` - Get maximum element
- `Min()` - Get minimum element
- `Sum()` - Get sum of all elements


### Set Operations

- `Add(element)` - Add element to set
- `Apply(function)` - Apply function to each element
- `Clone()` - Create shallow copy of set
- `Contains(value)` - Test if set contains value
- `ContainsFunc(predicate)` - Test if set contains element matching predicate
- `Count(predicate)` - Count elements matching predicate
- `Diff(set)` - Get elements in first set but not in second
- `Equals(set)` - Test set equality
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Intersection(set)` - Get elements present in both sets
- `IsEmpty()` - Test if set is empty
- `Length()` - Get number of elements
- `New(slices...)` - Create new set
- `NonEmpty()` - Test if set is not empty
- `Partition(predicate)` - Split set based on predicate
- `Random()` - Get random element
- `Remove(element)` - Remove element from set
- `String()` - Get string representation
- `ToSlice()` - Convert to Go slice
- `Union(set)` - Get elements present in either set
- `Values()` - Get iterator over values


### Collection Operations

These operations are available on all collections, including Sequence, List, and Set.

- `Count(predicate)` - Count elements matching predicate
- `Diff(collection)` - Get elements in first collection but not in second
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `ForAll(predicate)` - Test if predicate holds for all elements
- `GroupBy(function)` - Group elements by key function
- `Head()` - Get first element
- `Init()` - Get all elements except last
- `Intersect(collection)` - Get elements present in both collections
- `Last()` - Get last element
- `Map(function)` - Transform elements using function
- `MaxBy(function)` - Get maximum element by comparison function
- `MinBy(function)` - Get minimum element by comparison function
- `Partition(predicate)` - Split collection based on predicate
- `Reduce(function, initial)` - Reduce collection to single value
- `ReduceRight(function, initial)` - Right-to-left reduction
- `Reverse()` - Reverse order of elements
- `ReverseMap(function)` - Map elements in reverse order
- `SplitAt(n)` - Split collection at index n
- `Tail()` - Get all elements except first
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `DropWhile(predicate)` - Drop elements while predicate is true

## Contributing

Contributions are welcome! Feel free to submit a Pull Request.

If you have any ideas for new features or improvements, or would like to chat,
Feel free to reach out on [Reddit r/gopherslib](https://www.reddit.com/r/gopherslib).
