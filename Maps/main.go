package main

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

func main() {
	mall1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	mall2 := map[string]int{
		"one":   10,
		"three": 3,
	}

	// inserts
	maps.Insert(mall2, maps.All(mall1))
	fmt.Println("m2 is:", mall2)

	// collect
	scol1 := []string{"zero", "one", "two", "three"}
	mcol1 := maps.Collect(slices.All(scol1))
	fmt.Println("m1 is:", mcol1)

	mcol1 = maps.Collect(Backward(scol1))
	fmt.Println("m1 is:", mcol1)

	// Insert
	m1 := map[int]string{
		1000: "THOUSAND",
	}
	s1 := []string{"zero", "one", "two", "three"}
	maps.Insert(m1, slices.All(s1))
	fmt.Println("m1 is:", m1)

	m1 = map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	keys := slices.Sorted(maps.Keys(m1))
	fmt.Println(keys)

	m1 = map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	values := slices.Sorted(maps.Values(m1))
	fmt.Println(values)
}

func Backward[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i < len(s); i++ {
			if !yield(i*2, s[i]) {
				return
			}
		}
	}
}
