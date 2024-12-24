package main

import (
	"fmt"
)

var (
	ErrInvalidString = fmt.Errorf("invalid string")
)

func main() {
	input := "45"

	output, err := decodeString(input)

	if err != nil {
		fmt.Printf("invalid string: %s\n", err.Error())
		return
	}

	fmt.Printf("%s => %s\n", input, output)
}

func decodeString(s string) (string, error) {
	var result []rune
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if i+1 < len(s) && s[i+1] == '\\' {
				result = append(result, '\\')
				i++
			} else if i+1 < len(s) && s[i+1] == '4' {
				result = append(result, '4', '4', '4', '4')
				i++
			} else {
				return "", ErrInvalidString
			}
		} else {
			result = append(result, rune(s[i]))
		}
	}
	return string(result), nil
}
