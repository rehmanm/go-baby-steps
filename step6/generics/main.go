package main

import "fmt"

func main() {
	fmt.Println("Hello Generics")

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non Generic Sums: %v and %v \n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v \n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums Infered: %v and %v \n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums With Constraint: %v and %v \n", SumNumbers(ints), SumNumbers(floats))
}
