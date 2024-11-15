# Gophers - generic collections library

Gophers is the generic collections library offering tons of features right out of the box.

## Quick Start

```go
import (
  "github.com/charbz/gophers/pkg/list"
  "github.com/charbz/gophers/pkg/set"
  "github.com/charbz/gophers/pkg/collection"
)

// some examples using a comparable list

nums := list.NewComparableList([]int{1, 2, 2, 3, 4, 5, 5, 6, 9, 10})

nums.Filter(func(i int) bool { return i%2 == 0 }) // List [2,2,4,6,10]

nums.FilterNot(func(i int) bool { return i%2 == 0 }) // List [1,3,5,5,9]

nums.Distinct() // List [1,2,3,4,5,6,9,10]

nums.Distinct().Drop(2) // List [3,4,5,6,9,10]

nums.PartitionAt(6) // List [1,2,3,4,5,5],  List [6,9,10]

nums.Partition(func(i int) bool { return i%2 == 0 }) // List [2,2,4,6,10] List [1,3,5,5,9]

nums.Diff(list.NewComparableList([]int{1, 2, 3, 4})) // List [5,5,6,9,10]

nums.Max() // 10

nums.Min() // 1

nums.Sum() // 47

nums.Reverse() // List [10,9,6,5,5,4,3,2,2,1]

nums.Count(func(i int) bool { return i>5}) // 3

collection.Map(nums, func(i int) int { return i * 2 }) // List [2,4,4,6,8,10,10,12,18,20]

collection.Reduce(nums, func(acc int, i int) float64 { return acc + i/2 }, 0) // 23.5


// some examples using a set

letters := set.NewSet([]string{"A", "B", "C", "A", "C", "A"}) // Set ["A", "B", "C"]

other := set.NewSet([]string{"A", "B", "D"}) // Set ["A", "B", "D"]

letters.Intersect(other) // Set ["A", "B"]

letters.Union(other) // Set ["A", "B", "C", "D"]

letters.Diff(other) // Set ["C"]

collection.Map(letters, func(s string) string { return s + "!" }) // Set ["A!", "B!", "C!"]

collection.Reduce(letters, func(acc string, i string) string { return acc + i }, "") // "ABC"

// There are many more features, check out the docs for more details.

```

## Core Features

- **Collection** : A generic collection interface that provides a common interface for all collections.
- **Sequence** : An ordered collection wrapping a Go slice. Great for fast random access.
- **List** : An ordered collection wrapping a linked list. Great for fast insertion and removal, implementing queues and stacks.
- **Set** : A hash set implementation.

### Sequence Operations

- `Add(element)` - Append element to sequence
- `At(index)` - Get element at index
- `All()` - Get iterator over all elements
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
- `ForEach(function)` - Apply function to each element
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
- `ForEach(function)` - Apply function to each element
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
- `Clone()` - Create shallow copy of set
- `Contains(value)` - Test if set contains value
- `ContainsFunc(predicate)` - Test if set contains element matching predicate
- `Count(predicate)` - Count elements matching predicate
- `Diff(set)` - Get elements in first set but not in second
- `Equals(set)` - Test set equality
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `ForAll(predicate)` - Test if predicate holds for all elements
- `ForEach(function)` - Apply function to each element
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
- `ForEach(function)` - Apply function to each element
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

## License

This project is licensed under the MIT License - see the LICENSE file for details.
