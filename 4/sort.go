package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	kFlag := flag.Int("k", 1, "Column number to sort by (default is 1)")
	nFlag := flag.Bool("n", false, "Sort numerically")
	rFlag := flag.Bool("r", false, "Sort in reverse order")
	uFlag := flag.Bool("u", false, "Unique lines only")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Usage: go run sort.go [-k column] [-n] [-r] [-u] < input.txt > output.txt")
		return
	}

	inputFile := flag.Arg(0)
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.SliceStable(lines, func(i, j int) bool {
		colsI := strings.Fields(lines[i])
		colsJ := strings.Fields(lines[j])

		colI := colsI[*kFlag-1]
		colJ := colsJ[*kFlag-1]

		// если сортировка по числовому значению
		if *nFlag {
			numI, errI := strconv.Atoi(colI)
			numJ, errJ := strconv.Atoi(colJ)

			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}

		// лексикографическая сортировка
		if *rFlag {
			return colI > colJ
		}
		return colI < colJ
	})

	if *uFlag {
		lines = removeDuplicates(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func removeDuplicates(lines []string) []string {
	uniqueLines := make([]string, 0)
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			uniqueLines = append(uniqueLines, line)
			seen[line] = true
		}
	}
	return uniqueLines
}
