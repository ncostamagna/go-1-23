package main_test

import (
	"fmt"
	"testing"
	"unique"
)

type testStruct struct {
	z float64
	b string
}

func main() {

	v1 := testStruct{12.2, "test1"}
	v2 := testStruct{12.2, "test1"}

	r1 := unique.Make(v1)
	r2 := unique.Make(v2)

	fmt.Println(r1.Value(), r2.Value())
	fmt.Println(r1.Value() == r2.Value())
	fmt.Println(v1 == v2)
}

// more performance
func BenchmarkNormal(b *testing.B) {
	v1 := testStruct{12.2, "test1"}
	v2 := testStruct{12.2, "test1"}
	count := 0
	for i := 0; i < b.N; i++ {
		if v1 == v2 {
			count++
		}
	}
	fmt.Println("count: ", count)
}

// less performance
func BenchmarkUnique(b *testing.B) {
	v1 := testStruct{12.2, "test1"}
	v2 := testStruct{12.2, "test1"}

	r1 := unique.Make(v1).Value()
	r2 := unique.Make(v2).Value()
	count := 0
	for i := 0; i < b.N; i++ {
		if r1 == r2 {
			count++
		}
	}
	fmt.Println("count: ", count)
}
