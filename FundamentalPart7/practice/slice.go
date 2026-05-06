package practice

func removeDuplicates(nums []int) []int {
	seen := map[int]bool{}
	ans := []int{}
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			ans = append(ans, v)
		}
	}
	return ans
}

func chunk(slice []int, size int) [][]int {
	ans := [][]int{}
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		ans = append(ans, slice[i:end])
	}
	return ans
}

func rotate(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	return append(nums[n-k:], nums[:n-k]...)
}

func ModifySlice(sl []int) {
	sl[0] = 99
}

func TriggerAppend(sl []int) {
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
	}
}

// pointer and slice
type Project struct {
	Name string
}

func (p Project) UpdateName(newName string) {
	p.Name = newName
}

func (p *Project) UpdateNameReal(newName string) {
	p.Name = newName
}

func updateListName(projects []*Project, newNames []string) {
	for i, p := range projects {
		if i < len(newNames) {
			p.UpdateNameReal(newNames[i])
		}
	}
}
