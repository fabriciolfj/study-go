package generic

import "fmt"

type Numeric interface {
	~int8 | int16 | int32 | int64 | float32 | float64
}

type Smallint int8

func doubler[T Numeric](value T) T {
	return value * 2
}

func filterPositive[T Numeric](items []T) []T {
	var filtered []T
	for _, v := range items {
		if v > 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func ExecuteRestriction() {
	var four Smallint = 4
	fmt.Println(doubler(four))

	ints := []int8{-4, -3, -2, -1, 0, 1, 2, 3, 4}
	ints = filterPositive[int8](ints)
	fmt.Println(ints)

	floats := []float32{-4.5, -3.5, -2.5, -1.5, 0.5, 1.5, 2.5, 3.5, 4.5}
	floats = filterPositive[float32](floats)
	fmt.Println(floats)
}
