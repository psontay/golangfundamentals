package practice

import (
	"errors"
	"strings"
)

var ErrBadWord = errors.New("bad word")

func AnalyzeText(text string) (map[string]int, error) {
	substr := "badword"
	if strings.Contains(text, substr) {
		return nil, ErrBadWord
	}
	words := strings.Fields(text)
	ans := make(map[string]int, len(text))
	for _, w := range words {
		lower := strings.ToLower(w)
		if len(lower) < 3 {
			continue
		}
		ans[lower]++
	}
	return ans, nil
}
