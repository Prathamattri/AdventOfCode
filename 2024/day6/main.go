package day6

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func processInput(fileSrc string) [][]string {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	fileString := string(fileBytes)
	// Breaking lines from newline character into array of strings
	//The rules are divided from printing order by an empty line
	//so finding the empty line and breaking the array into 2 parts is the target here
	parts := strings.Split(fileString, "\n")

	mapped_area := [][]string{}
	for _, part := range parts {
		mapped_area = append(mapped_area, strings.Split(part, ""))
	}

	return mapped_area[:len(mapped_area)-1]
}

func isGuard(char string) bool {
	return char == ">" || char == "<" || char == "v" || char == "^"
}

// GLOBALS
var guardCharArr = []string{"^", ">", "v", "<"}
var guardDirMap = map[string][]int{
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
	"<": {-1, 0},
}
var guardDir = -1

// -------------------------

// Returns x_coord,  y_coord
func getGuardCurrPos(mappedArea [][]string) (int, int, string) {
	for row_ind := 0; row_ind < len(mappedArea); row_ind++ {

		for col_ind := 0; col_ind < len(mappedArea[0]); col_ind++ {
			g_char := mappedArea[row_ind][col_ind]
			if isGuard(g_char) {
				return row_ind, col_ind, g_char
			}
		}
	}

	return -1, -1, "."
}

func isInsideMap(mappedArea *[][]string, g_x, g_y int) bool {
	firstRow := *mappedArea
	colLen := len(firstRow[0])
	rowLen := len(firstRow)
	return (g_x >= 0 && g_x < colLen) || (g_y >= 0 && g_y < rowLen)
}

func moveGuard(mappedArea *[][]string, g_x, g_y *int, g_char string) {
	temp_map_area := *mappedArea
	temp_map_area[*g_x][*g_y] = "X"

	*g_x += guardDirMap[g_char][1]
	*g_y += guardDirMap[g_char][0]

	if *g_x != -1 && *g_y != len(temp_map_area[0]) && *g_y != -1 && *g_x != len(temp_map_area) {
		temp_map_area[*g_x][*g_y] = g_char
	}
	*mappedArea = temp_map_area
}

func turnGuardDir(mappedArea *[][]string, g_x, g_y int, g_char *string) {
	temp_map_area := *mappedArea
	guardDir = (guardDir + 1) % 4
	temp_map_area[g_x][g_y] = guardCharArr[guardDir]
	*g_char = guardCharArr[guardDir]
	*mappedArea = temp_map_area
}

func isHittingObstacle(mappedArea [][]string, g_x, g_y int, g_char string) bool {
	g_x += guardDirMap[g_char][1]
	g_y += guardDirMap[g_char][0]
	if mappedArea[g_x][g_y] == "#" {
		return true
	}
	return false
}

func isGoingOut(mappedArea [][]string, g_x, g_y int, g_char string) bool {
	g_x += guardDirMap[g_char][1]
	g_y += guardDirMap[g_char][0]
	if g_x == -1 || g_x == len(mappedArea[0]) || g_y == -1 || g_y == len(mappedArea) {
		return true
	}

	return false
}

func printMap(mappedArea *[][]string) {
	for _, row := range *mappedArea {
		fmt.Println(row)
	}
	fmt.Println()
}

func Part1(fileSrc string) int {
	mapped_area := processInput(fileSrc)
	g_x, g_y, g_char := getGuardCurrPos(mapped_area)
	guardDir = slices.Index(guardCharArr, g_char)
	distinct_visited_cells := 0

	for ; isInsideMap(&mapped_area, g_x, g_y); moveGuard(&mapped_area, &g_x, &g_y, g_char) {

		if isGoingOut(mapped_area, g_x, g_y, g_char) {
			moveGuard(&mapped_area, &g_x, &g_y, g_char)
			break
		}
		if isHittingObstacle(mapped_area, g_x, g_y, g_char) {
			turnGuardDir(&mapped_area, g_x, g_y, &g_char)
		}
	}

	for row_ind := 0; row_ind < len(mapped_area); row_ind++ {

		for col_ind := 0; col_ind < len(mapped_area[0]); col_ind++ {
			if mapped_area[row_ind][col_ind] == "X" {
				distinct_visited_cells++
			}
		}
	}
	printMap(&mapped_area)
	fmt.Println(distinct_visited_cells)

	return distinct_visited_cells
}

func isRepeatingPos(mappedArea [][]string, g_x, g_y int) bool {
	if mappedArea[g_x][g_y] == "#" {
		return true
	}
	return false
}

func Part2(fileSrc string) int {
	mapped_area := processInput(fileSrc)
	g_x, g_y, g_char := getGuardCurrPos(mapped_area)
	guardDir = slices.Index(guardCharArr, g_char)
	distinct_visited_cells := 0
	isNotInLoopCheck := true

	for ; isInsideMap(&mapped_area, g_x, g_y); moveGuard(&mapped_area, &g_x, &g_y, g_char) {

		if isNotInLoopCheck && isRepeatingPos(mapped_area, g_x, g_y) {
			isNotInLoopCheck = false
			turnGuardDir(&mapped_area, g_x, g_y, &g_char)
		}
		if isGoingOut(mapped_area, g_x, g_y, g_char) {
			moveGuard(&mapped_area, &g_x, &g_y, g_char)
			break
		}
		if isHittingObstacle(mapped_area, g_x, g_y, g_char) {
			turnGuardDir(&mapped_area, g_x, g_y, &g_char)
		}
	}

	for row_ind := 0; row_ind < len(mapped_area); row_ind++ {

		for col_ind := 0; col_ind < len(mapped_area[0]); col_ind++ {
			if mapped_area[row_ind][col_ind] == "X" {
				distinct_visited_cells++
			}
		}
	}
	printMap(&mapped_area)
	fmt.Println(distinct_visited_cells)

	return distinct_visited_cells
}
