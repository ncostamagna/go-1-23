package main

import (
	"fmt"
	"testing"
)

var s []string

func init(){
	i := 0

	for {

		if i > 100000 {
			break
		}
		if i % 4 == 0 {
			s = append(s, "a")
		}else if i % 5 == 0 {
			s = append(s, "b")
		}else{
			s = append(s, "c")
		}
		i++
	}

}

// GOEXPERIMENT=rangefunc go run main.go
// GOEXPERIMENT=rangefunc go test -bench=.
func Backward[E comparable](s []E, v E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			//fmt.Println(yield)
			//fmt.Println(i, s[i])
			if (s[i] != v) {
				continue
			}
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
	//fmt.Print(Backward(s))
	for i, el := range Backward(s, "b") {
		// we will see the values that we executed in yield
		// if we execute yield 2 times, we'll see each value 2 times
		fmt.Println(i, el, "for")
	}

}



func BenchmarkBackward(b *testing.B) {
	var n []string
    for i := 0; i < b.N; i++ {
        for _, el := range Backward(s, "b") {
			n = append(n, el)
		}
    }
}

func BenchmarkNormal(b *testing.B) {
	var n []string
    for i := 0; i < b.N; i++ {
        for _, el := range s {
			if el == "b" {
				n = append(n, el)
			}
		}
    }
}

// result: more or less the sema, there aren't difference between both

// command to run the test with the benchmark
// go test -bench=.