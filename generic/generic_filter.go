package generic

import "fmt"

func filter[T any](items []T, fx func(T) bool) []T {
	var filtred []T
	for _, item := range items {
		if fx(item) {
			filtred = append(filtred, item)
		}
	}

	return filtred
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
