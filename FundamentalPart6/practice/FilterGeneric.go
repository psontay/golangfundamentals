package practice

func Filter[T any](slice []T, fn func(item T) bool) []T {
	ans := []T{}
	for _, v := range slice {
		if fn(v) {
			ans = append(ans, v)
		}
	}
	return ans
}

func isEven(item int) bool {
	return item%2 == 0
}

func isLongerThan3(item string) bool {
	return len(item) > 3
}
