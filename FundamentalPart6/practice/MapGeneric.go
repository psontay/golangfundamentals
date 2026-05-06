package practice

func Map[T, U any](slice []T, fn func(T) U) []U {
	ans := make([]U, len(slice))
	for index, value := range slice {
		ans[index] = fn(value)
	}
	return ans
}

func double[T int | float64](v T) T {
	return v * 2
}
