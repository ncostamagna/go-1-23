package main

import (
	"fmt"
	"slices"
)

func main() {

	fmt.Printf("\n### All ###\n")
	// func All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
	// All returns an iterator over index-value pairs in the slice in the usual order.
	names := []string{"Alice", "Bob", "Vera"}
	for i, v := range slices.All(names) {
		fmt.Println(i, ":", v)
	}

	fmt.Printf("\n### AppendSeq ###\n")
	// func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice
	// AppendSeq appends the values from seq to the slice and returns the extended slice.
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}

	s := slices.AppendSeq([]int{1, 2}, seq)
	fmt.Println(s)

	fmt.Printf("\n### Backward ###\n")
	// func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]
	// Backward returns an iterator over index-value pairs in the slice, traversing it backward with descending indices.
	names2 := []string{"Alice", "Bob", "Vera"}
	for i, v := range slices.Backward(names2) {
		fmt.Println(i, ":", v)
	}

	fmt.Printf("\n### Chunk ###\n")
	// func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice]
	/*
	   Chunk returns an iterator over consecutive sub-slices of up to n elements of s.
	   All but the last sub-slice will have size n. All sub-slices are clipped to have no
	   capacity beyond the length. If s is empty, the sequence is empty: there is no empty
	   slice in the sequence. Chunk panics if n is less than 1.
	*/

	type Person struct {
		Name string
		Age  int
	}

	type People []Person

	people := People{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 5},
		{"Vera", 24},
		{"Zac", 15},
	}

	// Chunk people into []Person 2 elements at a time.
	for c := range slices.Chunk(people, 2) {
		fmt.Println(c)
	}

	fmt.Printf("\n### Colect ###\n")
	// func Collect[E any](seq iter.Seq[E]) []E
	// Collect collects values from seq into a new slice and returns it.
	seq2 := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}

	fmt.Println(slices.Collect(seq2))

	fmt.Printf("\n### Repeat ###\n")
	// func Repeat[S ~[]E, E any](x S, count int) S
	/*
		Repeat returns a new slice that repeats the provided slice the given number of times.
		The result has length and capacity (len(x) * count). The result is never nil.
		Repeat panics if count is negative or if the result of (len(x) * count) overflows.
	*/
	numbers := []int{0, 1, 2, 3}
	fmt.Println(slices.Repeat(numbers, 2))

	fmt.Printf("\n### Sorted ###\n")
	// func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E
	// Sorted collects values from seq into a new slice, sorts the slice, and returns it.
	seq4 := func(yield func(int) bool) {
		flag := -1
		for i := 0; i < 10; i += 2 {
			flag = -flag
			if !yield(i * flag) {
				return
			}
		}
	}

	s2 := slices.Sorted(seq4)
	fmt.Println(s2)
	fmt.Println(slices.IsSorted(s2))

	fmt.Printf("\n### Values ###\n")
	// func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]
	// Values returns an iterator that yields the slice elements in order.

	names3 := []string{"Alice", "Bob", "Vera"}
	for v := range slices.Values(names3) {
		fmt.Println(v)
	}

}
