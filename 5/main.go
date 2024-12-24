package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordLower := strings.ToLower(word)
		sortedWord := sortString(wordLower)

		anagrams[sortedWord] = append(anagrams[sortedWord], wordLower)
	}

	result := make(map[string][]string)

	for _, group := range anagrams {
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}

	return result
}

// Функция для сортировки строки по символам
func sortString(s string) string {
	slice := []rune(s)
	sort.Sort(sortRunes(slice))
	return string(slice)
}

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "катеп", "пятак"}

	anagramSets := findAnagramSets(words)

	for key, group := range anagramSets {
		fmt.Printf("Множество анаграмм для слова '%s': %v\n", key, group)
	}
}
