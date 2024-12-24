package main

import (
	"sort"
	"strings"
)

func SortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

}
