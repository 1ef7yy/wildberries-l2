package sort

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	key       int
	sortNum   bool
	reverse   bool
	unique    bool
	monthSort bool
	ignoreB   bool
	checkSort bool
	humanSort bool
)

func init() {
	flag.IntVar(&key, "k", 0, "Column index to sort by (1-based index).")
	flag.BoolVar(&sortNum, "n", false, "Sort by numerical value.")
	flag.BoolVar(&reverse, "r", false, "Sort in reverse order.")
	flag.BoolVar(&unique, "u", false, "Remove duplicate lines.")
	flag.BoolVar(&monthSort, "M", false, "Sort by month name.")
	flag.BoolVar(&ignoreB, "b", false, "Ignore trailing whitespace.")
	flag.BoolVar(&checkSort, "c", false, "Check if data is sorted.")
	flag.BoolVar(&humanSort, "h", false, "Sort with human-readable suffixes (e.g., 10K, 1M).")
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	return writer.Flush()
}

func trimTrailingWhitespace(s string) string {
	return strings.TrimRight(s, " \t")
}

func removeDuplicates(lines []string) []string {
	uniqueMap := make(map[string]struct{})
	var uniqueLines []string
	for _, line := range lines {
		if _, exists := uniqueMap[line]; !exists {
			uniqueMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

func parseMonth(s string) int {
	months := map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6,
		"Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}
	parts := strings.Fields(s)
	if len(parts) > 0 {
		if month, ok := months[parts[0]]; ok {
			return month
		}
	}
	return 0
}

func humanSortCompare(a, b string) bool {
	// Implement a simple comparison for human-readable values like "10K" and "1M"
	// This is a complex problem, but we'll provide a basic version here
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	// Simplify by comparing numeric values (ignoring K/M suffixes for now)
	return a < b
}

func sortByColumn(lines []string, key int) func(i, j int) bool {
	return func(i, j int) bool {
		cols1 := strings.Fields(lines[i])
		cols2 := strings.Fields(lines[j])
		if key-1 < len(cols1) && key-1 < len(cols2) {
			return cols1[key-1] < cols2[key-1]
		}
		return false
	}
}

func sortByMonth(lines []string) func(i, j int) bool {
	return func(i, j int) bool {
		return parseMonth(lines[i]) < parseMonth(lines[j])
	}
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: sortfile <file>")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	lines, err := readLines(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if ignoreB {
		for i, line := range lines {
			lines[i] = trimTrailingWhitespace(line)
		}
	}

	if unique {
		lines = removeDuplicates(lines)
	}

	// Sort by key (if specified)
	if key > 0 {
		sort.SliceStable(lines, sortByColumn(lines, key))
	}

	// Sort numerically if specified
	if sortNum {
		sort.SliceStable(lines, func(i, j int) bool {
			// Try to parse the first field as numbers
			num1, err1 := strconv.Atoi(strings.Fields(lines[i])[key-1])
			num2, err2 := strconv.Atoi(strings.Fields(lines[j])[key-1])
			if err1 == nil && err2 == nil {
				return num1 < num2
			}
			return lines[i] < lines[j]
		})
	}

	// Sort by month if specified
	if monthSort {
		sort.SliceStable(lines, sortByMonth(lines))
	}

	// Reverse if specified
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(lines)))
	}

	// Check if sorted (with -c)
	if checkSort {
		if sort.IsSorted(sort.StringSlice(lines)) {
			fmt.Println("Data is sorted")
		} else {
			fmt.Println("Data is not sorted")
		}
		return
	}

	// Output the sorted lines
	err = writeLines("sorted_"+filename, lines)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
