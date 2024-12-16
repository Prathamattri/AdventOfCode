package day10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInput(fileSrc string) [][]int {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	fileString := string(fileBytes)
	// Breaking lines from newline character into array of strings
	//The rules are divided from printing order by an empty line
	//so finding the empty line and breaking the array into 2 lines is the target here
	lines := strings.Split(fileString, "\n")

	lines = lines[:len(lines)-1]
	num_int_parent_arr := [][]int{}
	for _, line := range lines {
		num_int_arr := []int{}
		num_str_arr := strings.Split(line, "")
		for _, num_str := range num_str_arr {
			num_int, _ := strconv.Atoi(num_str)
			num_int_arr = append(num_int_arr, num_int)
		}
		num_int_parent_arr = append(num_int_parent_arr, num_int_arr)
	}
	return num_int_parent_arr
}

// Utility function for checking surrounding of a cell for part 1
func checkSurrounding(topographic_map [][]int, row_ind, col_ind int, row_max, col_max *int, score *int, prev_height int, visited_trailend *[][2]int) {

	if row_ind == -1 || col_ind == -1 || row_ind == *row_max || col_ind == *col_max || topographic_map[row_ind][col_ind] != prev_height+1 {
		return
	}

	//Checking if trail is already traversed
	temp_visited_topographic := *visited_trailend
	for _, pair := range temp_visited_topographic {
		if row_ind == pair[0] && col_ind == pair[1] {
			return
		}
	}

	if prev_height+1 == 9 {
		temp_visited_topographic = append(temp_visited_topographic, [2]int{row_ind, col_ind})
		*visited_trailend = temp_visited_topographic
		*score++
		return
	}

	prev_height = topographic_map[row_ind][col_ind]
	// Check top
	checkSurrounding(topographic_map, row_ind-1, col_ind, row_max, col_max, score, prev_height, visited_trailend)
	// Check right
	checkSurrounding(topographic_map, row_ind, col_ind+1, row_max, col_max, score, prev_height, visited_trailend)
	// Check bottom
	checkSurrounding(topographic_map, row_ind+1, col_ind, row_max, col_max, score, prev_height, visited_trailend)
	// Check left
	checkSurrounding(topographic_map, row_ind, col_ind-1, row_max, col_max, score, prev_height, visited_trailend)

}

func getScore(topographic_map [][]int, row_max, col_max *int, trailhead_pos [2]int, score *int) {
	row_ind, col_ind := trailhead_pos[0], trailhead_pos[1]
	visited_topographic_map := [][2]int{}
	checkSurrounding(topographic_map, row_ind, col_ind, row_max, col_max, score, -1, &visited_topographic_map)
}

func Part1(fileSrc string) int {
	topographic_map := processInput(fileSrc)
	total_score := 0

	row_len := len(topographic_map[0]) // also col max
	col_len := len(topographic_map)    // also row max

	// Finding list of all trailheads
	trailhead_positions := [][2]int{}
	for row_ind, row := range topographic_map {
		for col_ind, col := range row {
			if col == 0 {
				trailhead_positions = append(trailhead_positions, [2]int{row_ind, col_ind})
			}
		}
	}

	//Checking all trail routes using trailheads
	for _, trailhead_pos := range trailhead_positions {
		score := 0
		getScore(topographic_map, &col_len, &row_len, trailhead_pos, &score)
		total_score += score
	}

	fmt.Printf("Total Score : %d\n", total_score)
	return total_score
}

func Part2(fileSrc string) int {
	topographic_map := processInput(fileSrc)
	total_score := 0

	row_len := len(topographic_map[0]) // also col max
	col_len := len(topographic_map)    // also row max

	// Finding list of all trailheads
	trailhead_positions := [][2]int{}
	for row_ind, row := range topographic_map {
		for col_ind, col := range row {
			if col == 0 {
				trailhead_positions = append(trailhead_positions, [2]int{row_ind, col_ind})
			}
		}
	}

	//Checking all trail routes using trailheads
	for _, trailhead_pos := range trailhead_positions {
		score := 0
		getScore2(topographic_map, &col_len, &row_len, trailhead_pos, &score)
		total_score += score
	}

	fmt.Printf("Total Score : %d\n", total_score)
	return total_score
}
