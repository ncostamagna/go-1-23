package main

import (
	"unique"
	"fmt"
)

type testStruct struct {
	z float64
	b string
}

func main() {

	v1 := testStruct{ 12.2, "test1" }
	v2 := testStruct{ 12.2, "test1" }

	r1 := unique.Make(v1)
	r2 := unique.Make("v2")

	fmt.Println(r1.Value(), r2.Value())
	fmt.Println(r1.Value() == r2.Value())
	fmt.Println(v1 == v2)
}