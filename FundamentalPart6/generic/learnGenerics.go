package generic

// before Go 1.18
//func sumInt ( nums []int) int {
//	ans := 0
//	for _, num := range nums { ans += num }
//	return ans
//}
//func sumFloat( nums []float64) float64  {
//	ans := 0.0
//	for _, num := range nums { ans += num }
//	return ans
//}

// GENERICS

type Number interface {
	int | float64
}

func sum[T Number](nums []T) T {
	var ans T
	for _, num := range nums {
		ans += num
	}
	return ans
}

// comparable
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// struct with generics
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
