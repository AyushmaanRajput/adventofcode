package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 1. Open the file
	// 2. Initialize a new scanner and read by lines
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input file", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	list1 := []int{}
	list2 := []int{}
	for fileScanner.Scan() {
		lineText := fileScanner.Text()
		if strings.TrimSpace(lineText) == "" {
			continue // Skip empty lines
		}
		words := strings.Fields(lineText)
		firstNumber := convertStringToInt(words[0])
		secondNumber := convertStringToInt(words[1])

		list1 = append(list1, firstNumber)
		list2 = append(list2, secondNumber)
	}
	// 1. .Day1
	day1_part1(list1, list2)
	day1_part2(list1, list2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func convertStringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// fmt.Println("Integer value:", value)
	}
	return value
}

func day1_part1(list1, list2 []int) {
	sort.Ints(list1)
	sort.Ints(list2)
	// 2. loop over both list simultaneously and sum up the differences
	total := 0
	// minLen := min(len(list1), len(list2))
	for i := 0; i < len(list1); i++ {
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println("Output for Day 1 part 1", total)
}
func day1_part2(list1, list2 []int) {
	freq := make(map[int]int) // Initialize the map

	for _, val := range list1 {
		freq[val] = 0 // Assign any value (0 in this case), ensuring the key exists
	}

	for _, val := range list2 {
		if count, exists := freq[val]; exists {
			// already exits
			freq[val] = count + 1
		} else {
			freq[val] = 1
		}
	}
	var similiarityScore int
	for _, val := range list1 {
		similiarityScore += val * freq[val]
	}
	fmt.Println("Output for Day 1 part 2", similiarityScore)

}
