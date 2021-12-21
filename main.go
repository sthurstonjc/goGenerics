package main

import (
	"constraints"
	"fmt"
	"os"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// SumNumbers sums the numbers of map m
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int{
		"first":  42,
		"second": 74,
		"third":  1,
	}

	int64s := map[string]int64{
		"1st": 420,
		"2nd": 740,
		"3rd": 10,
	}

	floats := map[string]float32{
		"one":   42.0,
		"two":   74.1,
		"three": 1.42,
	}

	float64s := map[string]float64{
		"ein":  420.69,
		"zwei": 74.31459,
		"drei": 10.10,
	}

	fmt.Printf("Sum ints: %v\n", SumNumbers(ints))
	fmt.Printf("Sum int64s: %v\n", SumNumbers(int64s))
	fmt.Printf("Sum floats: %v\n", SumNumbers(floats))
	fmt.Printf("Sum float64s: %v\n", SumNumbers(float64s))

	os.Exit(0)
}
