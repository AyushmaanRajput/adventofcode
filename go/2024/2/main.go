package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	safeCount := 0
	safeCountDamped := 0
	for fileScanner.Scan() {
		reportLine := fileScanner.Text()
		report := strings.Fields(reportLine)
		// fmt.Println("Report", report)
		var intReport []int
		for _, val := range report {
			// convert each strign to int
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("Some line wasn't a int", err)
			}
			intReport = append(intReport, intVal)
		}
		if day2_part1(intReport) {
			safeCount += 1
		}
		if day2_part2(intReport) {
			safeCountDamped += 1
		}
	}
	fmt.Println("Safe count", safeCount)
	fmt.Println("Safe count damped", safeCountDamped)
}

func day2_part1(report []int) bool {
	// Figure out whether the report is safe or not
	// Rules:
	// 1. The levels are either all increasing or all decreasing.
	// 2. Any two adjacent levels differ by at least one and at most three.
	if len(report) < 2 {
		return true
	}

	// Determine the order
	order := ""
	if report[1] > report[0] {
		order = "asc"
	} else if report[1] < report[0] {
		order = "desc"
	} else {
		return false // If the first two elements are equal, it's invalid
	}

	for i := 1; i < len(report); i++ {
		diff := int(math.Abs(float64(report[i] - report[i-1])))

		// Rule 2: Difference must be between 1 and 3
		if diff < 1 || diff > 3 {
			return false
		}

		// Rule 1: Check order consistency
		if order == "asc" && report[i] < report[i-1] {
			return false
		}
		if order == "desc" && report[i] > report[i-1] {
			return false
		}
	}
	return true
}
func day2_part2(report []int) bool {
	// Figure out whether the report is safe or not
	// Rules:
	// 1. The levels are either all increasing or all decreasing.
	// 2. Any two adjacent levels differ by at least one and at most three.
	// 3. Atmost 1 bad value can be there

	// Approach
	// we first check if its safe if yes just return true
	// if not loop for all levels of the report and check whether a new array without containing hte element is safe or not if yes
	// this should only be possible for atmost 1 wrong level
	if day2_part1(report) {
		return true // Already safe, no need to modify
	}

	var safeCount int
	for i := range report {
		mod := append([]int{}, report[:i]...) // Copy elements before i
		mod = append(mod, report[i+1:]...)    // Skip the current element and append remaining elements

		if day2_part1(mod) {
			safeCount++
			break
		}
	}

	return safeCount==1 
}
