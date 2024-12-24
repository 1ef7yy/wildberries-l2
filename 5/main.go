package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для поиска множеств анаграмм
func findAnagramSets(words []string) map[string][]string {
	// Мапа для хранения групп анаграмм
	anagrams := make(map[string][]string)

	// Обрабатываем каждое слово
	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем символы
		wordLower := strings.ToLower(word)
		sortedWord := sortString(wordLower)

		// Добавляем слово в соответствующую группу анаграмм
		anagrams[sortedWord] = append(anagrams[sortedWord], wordLower)
	}

	// Результат будет содержать только те группы, которые имеют больше одного элемента
	result := make(map[string][]string)

	for _, group := range anagrams {
		// Если в группе больше одного слова, то добавляем в результат
		if len(group) > 1 {
			// Сортируем группу по возрастанию
			sort.Strings(group)
			// Добавляем в мапу ключ, который соответствует первому слову в группе
			result[group[0]] = group
		}
	}

	return result
}

// Функция для сортировки строки по символам
func sortString(s string) string {
	// Преобразуем строку в слайс рун, сортируем и преобразуем обратно в строку
	slice := []rune(s)
	sort.Sort(sortRunes(slice))
	return string(slice)
}

// Для сортировки рун
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
	// Пример данных
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "катеп", "пятак"}

	// Находим множества анаграмм
	anagramSets := findAnagramSets(words)

	// Выводим результат
	for key, group := range anagramSets {
		fmt.Printf("Множество анаграмм для слова '%s': %v\n", key, group)
	}
}
