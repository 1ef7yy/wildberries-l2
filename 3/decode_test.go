package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	tests := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"":         "",
		`qwe\4\5`:  "qwe45",
		`qwe\45`:   "qwe44444",
		`qwe\\5`:   `qwe\\\\\`,
	}

	for input, expected := range tests {
		output, err := decodeString(input)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}
		if output != expected {
			t.Errorf("input: %s, expected: %s, got: %s", input, expected, output)
		}
	}
}

func TestDecodeError(t *testing.T) {
	tests := []string{
		"45",
	}

	for _, input := range tests {
		val, err := decodeString(input)
		if err == nil {
			t.Errorf("input: %s, expected: error, got: val - %s, err - %s", input, val, err)
		}
	}
}
