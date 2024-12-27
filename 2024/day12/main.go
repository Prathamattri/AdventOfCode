package day12

import (
	"fmt"
	"os"
	"strings"
)

var cache = map[[2]int]int{}

func processInput(fileSrc string) [][]rune {
	file_bytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	file_string := string(file_bytes)
	plot_map := [][]rune{}
	lines := strings.Split(file_string, "\n")
	for _, line := range lines[:len(lines)-1] {
		plot_map = append(plot_map, []rune(line))
	}
	return plot_map
}

func isValidIndex(index, low_bound int, up_bound *int) bool {
	if index >= low_bound && index < *up_bound {
		return true
	}
	return false
}

func checkSurrounding(visited_arr *[][]bool, plot_map [][]rune, plant_type rune, plant_y, plant_x int, plot_size *int) (bool, int, int) {
	copy_visited_arr := *visited_arr
	if !isValidIndex(plant_y, 0, plot_size) || !isValidIndex(plant_x, 0, plot_size) || (plant_type != plot_map[plant_y][plant_x]) {
		return false, 0, 0
	}
	if copy_visited_arr[plant_y][plant_x] {
		return true, 0, 0
	}
	copy_visited_arr[plant_y][plant_x] = true
	*visited_arr = copy_visited_arr

	fence_count := 4

	// checking top
	visited, top, p_count_top := checkSurrounding(visited_arr, plot_map, plant_type, plant_y-1, plant_x, plot_size)
	if visited {
		fence_count--
	}
	// checking right
	visited, right, p_count_right := checkSurrounding(visited_arr, plot_map, plant_type, plant_y, plant_x+1, plot_size)
	if visited {
		fence_count--
	}
	// checking bottom
	visited, bottom, p_count_bottom := checkSurrounding(visited_arr, plot_map, plant_type, plant_y+1, plant_x, plot_size)
	if visited {
		fence_count--
	}
	// checking left
	visited, left, p_count_left := checkSurrounding(visited_arr, plot_map, plant_type, plant_y, plant_x-1, plot_size)
	if visited {
		fence_count--
	}

	total_fence_count := fence_count + top + right + bottom + left
	total_plant_count := p_count_top + p_count_right + p_count_bottom + p_count_left + 1

	return true, total_fence_count, total_plant_count
}

func Part1(fileSrc string) int {
	plot_map := processInput(fileSrc)
	total_fencing_cost := 0

	visited_plant_arr := make([][]bool, len(plot_map))

	for ind, row := range visited_plant_arr {
		row = make([]bool, len(plot_map))
		visited_plant_arr[ind] = row
	}

	plot_size := len(plot_map)
	// only using size for both col and row size is enough since plot is a square

	for row_num := 0; row_num < plot_size; row_num++ {
		for col_num := 0; col_num < plot_size; col_num++ {
			if visited_plant_arr[row_num][col_num] {
				continue
			}
			plant_type := plot_map[row_num][col_num]
			_, fence_count, plant_count := checkSurrounding(&visited_plant_arr, plot_map, plant_type, row_num, col_num, &plot_size)

			total_fencing_cost += fence_count * plant_count

		}
	}
	fmt.Printf("\nCost of fencing : %d\n", total_fencing_cost)
	return total_fencing_cost
}

func Part2(fileSrc string) {

}
