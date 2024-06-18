package main

import (
	"fmt"
)

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			fmt.Println(yield)
			fmt.Println(i, s[i])
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(i+3, s[i]) {

				// Where clean-up code goes
				return
			}
		}
	}
}

func main(){

	s := []string{"a", "b", "c"}
	fmt.Print(Backward(s))
	for i, el := range Backward(s) {
		// we will see the values that we executed in yield
		// if we execute yield 2 times, we'll see each value 2 times
		fmt.Println(i, el, "for")
	}

}
