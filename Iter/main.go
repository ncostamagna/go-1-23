package main

import (
	"fmt"
	"iter"
)

func Backward[E any](s []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(s[i]) {
				// Where clean-up code goes
				return
			}
		}
	}
}

func Backward2[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(i, s[i]) {
				// Where clean-up code goes
				return
			}
		}
	}
}

func PrintAll(seq []int) {
	for v := range Backward(seq) {
		fmt.Println(v)
	}
}

func PrintAllPush(seq []int) {
	n, s := iter.Pull(Backward(seq))
	for {
		v, b := n()
		if !b {
			s()
			break
		}
		fmt.Println(v, b)
	}
}

func PrintAllPush2(seq []int) {
	n, s := iter.Pull2(Backward2(seq))
	for {
		k, v, b := n()
		if !b {
			s()
			break
		}
		fmt.Println(k, v, b)
	}
}

func PrintAll2(seq []int) {
	for _, v := range Backward2(seq) {
		fmt.Println(v)
	}
}

func main() {
	v := []int{3, 2, 2, 8, 5}
	PrintAll(v)

	PrintAllPush(v)

	PrintAllPush2(v)
}
