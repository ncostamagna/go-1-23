package main_test


import (
	"fmt"
	"iter"
		"testing"
)

func Backward[E any](s []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(s[i]) {
				// Where clean-up code goes
				return
			}
		}
	}
}

func Backward2[E comparable](s []E, v E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := len(s)-1; i >= 0; i-- {
			// if we add i+3 for 3 elements, we'll see 3, 4 and 5 indexes
			if !yield(i, s[i]) {
				// Where clean-up code goes
				return
			}
		}
	}
}


func PrintAllPush(seq []int) int{
	num := 0
	n, s := iter.Pull(Backward(seq))
	for  {
		v, b := n()
		if !b {
			s()
			break
		}
		num += v
	}
	return num
}

func PrintAll2(seq []int) {
	n := 0
	for _, v := range Backward2(seq, 2) {
		n += v
	}
	fmt.Println(n)
}


var s []int

func init(){
	i := 0

	for {

		if i > 99999999 {
			break
		}
		if i % 4 == 0 {
			s = append(s, 1)
		}else if i % 5 == 0 {
			s = append(s, 3)
		}else{
			s = append(s, 4)
		}
		i++
	}

}


// more performance 
func BenchmarkFor(b *testing.B) {

	var n int64 
	count := 0
	fmt.Println("len ", len(s))
    for i := 0; i < 1; i++ {
		for v := range s {
			n += int64(v)
			count++
		}
    }
	fmt.Println("count: ", count)
	fmt.Println("resutl: ", n)
}

// less performance
func BenchmarkPull(b *testing.B) {
	n, stop := iter.Pull(Backward(s))
	var num int64
	fmt.Println("len ", len(s))
	count := 0
    for i := 0; i < 1; i++ {
       for  {
		v, b := n()
		if !b {
			stop()
			break
		}
		num += int64(v)
		count++
	}
    }
	fmt.Println("count: ", count)
	fmt.Println("resutl: ", num)
}