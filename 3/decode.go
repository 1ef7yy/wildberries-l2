package main

import "fmt"

func main() {
	input := "45" // -> aaaabccddddde

	output, err := decodeString(input)

	if err != nil {
		fmt.Printf("invalid string: %s\n", err.Error())
		return
	}

	fmt.Printf("%s => %s\n", input, output)
}

func decodeString(s string) (string, error) {
	var res string

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			if i == 0 || isDigit(s[i-1]) {
				return "", fmt.Errorf("invalid string")
			}
		}

		if isLetter(s[i]) {
			if i+1 < len(s) && isDigit(s[i+1]) {
				res += repeatRune(s[i], int(s[i+1]-'0'))
				i++ // skip digit
			} else {
				res += string(s[i])
			}
		}

		if s[i] == '\\' {
			if s[i+1] == '\\' {
				if i+2 < len(s) && isDigit(s[i+2]) {
					res += repeatRune(s[i+1], int(s[i+2]-'0'))
				} else if i+2 < len(s) && isLetter(s[i+2]) {
					res += string(s[i+1])
				}
				i += 2 // skip digit
			}

			if isDigit(s[i+1]) {
				if i+2 < len(s) && isDigit(s[i+2]) {
					res += repeatRune(s[i+1], int(s[i+2]-'0'))
				} else if i+2 < len(s) && isLetter(s[i+2]) {
					res += string(s[i+1])
				}
			}
		}
	}

	return res, nil
}

func repeatRune(r byte, n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += string(r)
	}
	return res
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r byte) bool {
	return r >= 'a' && r <= 'z'
}
