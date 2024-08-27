package main

import (
	"fmt"
)

func Backward[E comparable](s []E, v E) func(func(E) bool) {
	return func(yield func(E) bool) {
		// for i := len(s) - 1; i >= 0; i-- {
		for i := 0; i < len(s); i++ {
			//fmt.Println(yield)
			//fmt.Println(i, s[i])
			//if s[i] != v {
			//	continue
			//}
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(s[i]) {

				// Where clean-up code goes
				return
			}
		}
	}
}

// values that we returs
func Backward2[E comparable](s []E, v E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != v {
				continue
			}

			// value that we return
			if !yield(i+1, s[i]) {

				fmt.Println("finish for")
				return
			}
		}
	}
}

func GetElements() []string {
	return []string{"a", "b", "c", "d", "e"}
}

func iterate[V any](f func() []V) func(func(V) bool) {

	return func(yield func(V) bool) {
		val := f()

		for _, v := range val {

			if !yield(v) {
				fmt.Println("finish iter")
				return
			}
		}
	}

}

// NO
func iterateParam(yield func(int) bool) {

	val := []int{1, 2, 3, 4, 5}
	for _, v := range val {

		fmt.Println(v)
		if !yield(v) {
			fmt.Println("finish iter")
			return
		}
	}
	fmt.Println("test")

}

func main() {

	s := []string{"a", "b", "c"}

	for v := range Backward(s, "b") {

		fmt.Println(v, "for")
	}

	for i, el := range Backward2(s, "b") {
		// we will see the values that we executed in yield
		// if we execute yield 2 times, we'll see each value 2 times
		fmt.Println(i, el, "for")

		// execute yield: break
	}

	fmt.Println("Iterate")
	for v := range iterate(GetElements) {

		fmt.Println(v)
		break
	}

	// NO
	fmt.Println("Iterate param")
	for v := range iterateParam {
		fmt.Println(v)
		break
	}

}
