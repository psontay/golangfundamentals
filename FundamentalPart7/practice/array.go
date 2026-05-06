package practice

import "fmt"

func bai1(arr [5]int) int {
	var ans int = 0
	for index, _ := range arr {
		ans += arr[index]
	}
	return ans
}

func reverseArray(arr [5]int) [5]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func createMatrix(arr [3][3]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(arr[i][j])
		}
	}
}
