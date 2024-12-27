package day14

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func checkError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func processInput(fileSrc string) ([][2]int, [][2]int, [2]int) {
	file_bytes, err := os.ReadFile(fileSrc)
	checkError(err, "Unable to read input file")

	//Regex for getting positions
	regex := regexp.MustCompile("p.+\\s")
	positions_temp := regex.FindAll(file_bytes, -1)

	//Regex for getting velocities
	regex = regexp.MustCompile("v.+")
	velocities_temp := regex.FindAll(file_bytes, -1)

	//Regex for getting map size
	regex = regexp.MustCompile("size.+")
	map_size := regex.FindAll(file_bytes, -1)

	//Regex for getting digits
	regex = regexp.MustCompile("-?[0-9]+")

	positions := [][2]int{}
	for _, position_bytes := range positions_temp {
		position_str := regex.FindAllString(string(position_bytes), -1)

		num1, err := strconv.Atoi(position_str[0])
		checkError(err, "Unable to convert given string to number")
		num2, err := strconv.Atoi(position_str[1])
		checkError(err, "Unable to convert given string to number")

		positions = append(positions, [2]int{num1, num2})
	}

	velocities := [][2]int{}
	for _, position_bytes := range velocities_temp {
		position_str := regex.FindAllString(string(position_bytes), -1)

		num1, err := strconv.Atoi(position_str[0])
		checkError(err, "Unable to convert given string to number")
		num2, err := strconv.Atoi(position_str[1])
		checkError(err, "Unable to convert given string to number")

		velocities = append(velocities, [2]int{num1, num2})
	}

	position_str := regex.FindAllString(string(map_size[0]), -1)

	num1, err := strconv.Atoi(position_str[0])
	checkError(err, "Unable to convert given string to number")
	num2, err := strconv.Atoi(position_str[1])
	checkError(err, "Unable to convert given string to number")
	map_props := [2]int{num1, num2}

	return positions, velocities, map_props
}

func custom_modulo(num1, num2 *int) {
	if *num1 < 0 {
		*num1 = *num2 + *num1
	} else {
		*num1 = *num1 % *num2
	}
}

func get_robots_in_quadrant(robo_x, robo_y, omitted_col, omitted_row *int) (int, int) {
	if *robo_x == *omitted_col || *robo_y == *omitted_row {
		return -1, -1
	}

	quadrant_x := 0
	quadrant_y := 0
	if *robo_x > *omitted_col {
		quadrant_x = 1
	}
	if *robo_y > *omitted_row {
		quadrant_y = 1
	}
	return quadrant_x, quadrant_y
}

func Part1(fileSrc string) int {
	positions, velocities, map_size := processInput(fileSrc)

	map_w := map_size[0]
	map_h := map_size[1]
	for i := 0; i < 100; i++ {
		for robo_num := 0; robo_num < len(positions); robo_num++ {
			positions[robo_num][0] += velocities[robo_num][0]
			positions[robo_num][1] += velocities[robo_num][1]

			custom_modulo(&positions[robo_num][0], &map_w)
			custom_modulo(&positions[robo_num][1], &map_h)
		}
	}

	map_w_half := map_w / 2
	map_h_half := map_h / 2

	quadrants := [2][2]int{}

	for _, position := range positions {
		quad_x, quad_y := get_robots_in_quadrant(&position[0], &position[1], &map_w_half, &map_h_half)
		if quad_x == -1 {
			continue
		}
		quadrants[quad_x][quad_y]++
	}

	result := quadrants[0][0] * quadrants[0][1] * quadrants[1][0] * quadrants[1][1]
	fmt.Println(result)
	return result
}

func print_map(map_img [][]bool) {
	for _, row := range map_img {
		for _, cell := range row {
			if cell == true {
				fmt.Print("# ")
			} else {
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

func Part2(fileSrc string) {
	positions, velocities, map_size := processInput(fileSrc)

	map_w := map_size[0]
	map_h := map_size[1]
	map_image := [][]bool{}
	for row := 0; row < map_h; row++ {
		map_row := make([]bool, map_w)
		map_image = append(map_image, map_row)
	}
	positions_len := len(positions)
	for i := 0; i >= 0; i++ {
		for robo_num := 0; robo_num < positions_len; robo_num++ {
			x := positions[robo_num][0]
			y := positions[robo_num][1]
			map_image[y][x] = false
			x += velocities[robo_num][0]
			y += velocities[robo_num][1]

			custom_modulo(&x, &map_w)
			custom_modulo(&y, &map_h)
			map_image[y][x] = true
			positions[robo_num][0] = x
			positions[robo_num][1] = y
		}
		sum_of_squares := [2]int{0, 0}
		sum_of_positions := [2]int{0, 0}
		for robo_num := 0; robo_num < positions_len; robo_num++ {
			x := positions[robo_num][0]
			y := positions[robo_num][1]

			sum_of_squares[0] += x * x
			sum_of_squares[1] += y * y

			sum_of_positions[0] += x
			sum_of_positions[1] += y
		}

		//Calculating standard deviation
		st_dev_x := math.Sqrt(float64(sum_of_squares[0]/positions_len - sum_of_positions[0]*sum_of_positions[0]/(positions_len*positions_len)))
		st_dev_y := math.Sqrt(float64(sum_of_squares[1]/positions_len - sum_of_positions[1]*sum_of_positions[1]/(positions_len*positions_len)))

		if st_dev_x < 18.2 && st_dev_y < 19.2 {

			fmt.Println(sum_of_squares[0]/positions_len - sum_of_positions[0]*sum_of_positions[0]/(positions_len*positions_len))
			fmt.Printf("\nStandard Deviation : \n\tx: %f\ty: %f\nSeconds : %d \n\n", st_dev_x, st_dev_y, i+1)

			print_map(map_image)

			fmt.Printf("\n\033[1;30;47mANSWER : %d\033[0m\n\n", i+1)
			break
		}
	}

}
