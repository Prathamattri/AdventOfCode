package day10

// Utility function for checking surrounding of a cell for part 1
func checkSurrounding2(topographic_map [][]int, row_ind, col_ind int, row_max, col_max *int, score *int, prev_height int) {

	if row_ind == -1 || col_ind == -1 || row_ind == *row_max || col_ind == *col_max || topographic_map[row_ind][col_ind] != prev_height+1 {
		return
	}

	if prev_height+1 == 9 {
		*score++
		return
	}

	prev_height = topographic_map[row_ind][col_ind]
	// Check top
	checkSurrounding2(topographic_map, row_ind-1, col_ind, row_max, col_max, score, prev_height)
	// Check right
	checkSurrounding2(topographic_map, row_ind, col_ind+1, row_max, col_max, score, prev_height)
	// Check bottom
	checkSurrounding2(topographic_map, row_ind+1, col_ind, row_max, col_max, score, prev_height)
	// Check left
	checkSurrounding2(topographic_map, row_ind, col_ind-1, row_max, col_max, score, prev_height)

}
func getScore2(topographic_map [][]int, row_max, col_max *int, trailhead_pos [2]int, score *int) {
	row_ind, col_ind := trailhead_pos[0], trailhead_pos[1]
	checkSurrounding2(topographic_map, row_ind, col_ind, row_max, col_max, score, -1)
}
