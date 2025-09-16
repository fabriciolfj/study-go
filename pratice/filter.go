package pratice

func filter[T any](items []T, fix func(T) bool) (result []T) {
	for _, item := range items {
		if fix(item) {
			result = append(result, item)
		}
	}

	return result
}
