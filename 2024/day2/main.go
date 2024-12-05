package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// --- Day 2: Red-Nosed Reports ---
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func isReportSafe(report []int) bool {
	isSafeFlag := true
	flow := report[0] - report[1]

	for ind := 0; ind < len(report)-1; ind++ {
		localDiff := report[ind] - report[ind+1]
		// Check if 1 <= diff <= 3
		if 1 > abs(localDiff) || abs(localDiff) > 3 {
			isSafeFlag = false
			break
		}
		// Check if strictly increasing or strictly decreasing
		if flow*localDiff <= -1 {
			isSafeFlag = false
			break
		}
	}
	return isSafeFlag
}

// Part 1
func Part1(fileSource string) int {
	reportsArr := ProcessInput(fileSource)
	safeReportCount := 0
	for _, report := range reportsArr {

		if isReportSafe(report) {
			safeReportCount++
		}
	}
	fmt.Println(safeReportCount)
	return safeReportCount
}

// Part 2
func Part2(fileSource string) int {
	reportArr := ProcessInput(fileSource)
	safeReportCount := 0

	for _, report := range reportArr {

		if isReportSafe(report) {
			safeReportCount++
			continue
		}

		for ind := 0; ind < len(report); ind++ {
			temp_report := make([]int, len(report))
			copy(temp_report, report)
			temp_report = append(temp_report[0:ind], temp_report[ind+1:]...)
			if isReportSafe(temp_report) {
				safeReportCount++
				break
			}
		}
	}

	fmt.Println(safeReportCount)
	return safeReportCount
}

func ProcessInput(fileSrc string) [][]int {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		str := fmt.Sprintf("Unable to read input file : %v", err)
		panic(str)
	}

	fileString := string(fileBytes)

	reports := strings.Split(fileString, "\n")
	reportsArray := [][]int{}

	for ind := 0; ind < len(reports)-1; ind++ {
		levels := strings.Split(reports[ind], " ")
		levelsArr := []int{}

		for _, level := range levels {
			level, err := strconv.Atoi(level)
			if err != nil {
				panic("Error processing input")
			}
			levelsArr = append(levelsArr, level)
		}
		reportsArray = append(reportsArray, levelsArr)
	}
	return reportsArray
}
