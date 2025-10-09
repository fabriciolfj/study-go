package generic

import "fmt"

func filter[T any](items []T, predicate func(T) bool) []T {
	var filter []T
	for _, item := range items {
		if predicate(item) {
			filter = append(filter, item)
		}
	}

	return filter
}

func ExecutarFilter() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := filter[int](items, func(item int) bool {
		if item%2 == 0 {
			return true
		}

		return false
	})

	fmt.Println(result)
}
