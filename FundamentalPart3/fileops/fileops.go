package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(balance float64, filename string) {
	floatText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(floatText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	valueText := string(data)
	value, convErr := strconv.ParseFloat(valueText, 64)
	if convErr != nil {
		return 0, errors.New("failed to parse text")
	}
	return value, nil
}
