package benchmarking

import "fmt"

func Execute() {
	sorted := bubbleSort([]int{20, 19, 3, 76, 1, 7, 4, 17})
	fmt.Println(sorted)
}

func bubbleSort(in []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < len(in); i++ {
			if in[i-1] > in[i] {
				in[i-1], in[i] = in[i], in[i-1]
				sorted = false
			}
		}
	}

	return in
}
