package day13

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var cache = map[[2]int]int{}

func processInput(fileSrc string) ([][2]int, [][2]int, [][2]int) {
	file_bytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	// file_string := string(file_bytes)

	//regex for getting all the numbers from a line
	regex_num := regexp.MustCompile("[0-9]+")

	button_a_coeff_pairs := [][2]int{}

	regex := regexp.MustCompile(".* A.*\n")
	lines_with_btn_a := regex.FindAll(file_bytes, -1)

	for _, line := range lines_with_btn_a {
		nums_str := regex_num.FindAllString(string(line), -1)
		num1, _ := strconv.Atoi(nums_str[0])
		num2, _ := strconv.Atoi(nums_str[1])
		button_a_coeff_pairs = append(button_a_coeff_pairs, [2]int{num1, num2})
	}

	button_b_coeff_pairs := [][2]int{}

	regex = regexp.MustCompile(".* B.*\n")
	lines_with_btn_b := regex.FindAll(file_bytes, -1)
	for _, line := range lines_with_btn_b {
		nums_str := regex_num.FindAllString(string(line), -1)
		num1, _ := strconv.Atoi(nums_str[0])
		num2, _ := strconv.Atoi(nums_str[1])
		button_b_coeff_pairs = append(button_b_coeff_pairs, [2]int{num1, num2})
	}

	regex = regexp.MustCompile("Prize.*\n") // for getting line with prize coordinates

	lines_with_prize_coords := regex.FindAll(file_bytes, -1)
	target_coord := [][2]int{}
	for _, line := range lines_with_prize_coords {
		nums_str := regex_num.FindAllString(string(line), -1)
		num1, _ := strconv.Atoi(nums_str[0])
		num2, _ := strconv.Atoi(nums_str[1])
		target_coord = append(target_coord, [2]int{num1, num2})
	}
	return button_a_coeff_pairs, button_b_coeff_pairs, target_coord
}

func checkIfNumEqualItsFloat(num float64) bool {
	return math.Round(num) == math.Round(num*100)/100
}

func Part1(fileSrc string) float64 {
	button_a_coeff_pairs, button_b_coeff_pairs, target_coord := processInput(fileSrc)

	tokens := 0.0

	for machine_num := 0; machine_num < len(button_a_coeff_pairs); machine_num++ {
		btn_a_coeffs := button_a_coeff_pairs[machine_num]
		btn_b_coeffs := button_b_coeff_pairs[machine_num]

		//Equations are like :
		// 1: a1(X) + a2(Y) = c1 => a1(X) + a2(Y) - c1 = 0
		// 2: b1(X) + b2(Y) = c2 => b1(X) + b2(Y) - c2 = 0

		c1 := target_coord[machine_num][0]
		c2 := target_coord[machine_num][1]

		a1 := btn_a_coeffs[0]
		b1 := btn_b_coeffs[0]

		a2 := btn_a_coeffs[1]
		b2 := btn_b_coeffs[1]

		if a1*b2 != a2*b1 {

			delta := float64(a1*b2 - a2*b1)
			delta1 := a1*c2 - a2*c1
			delta2 := -(b1*c2 - b2*c1)

			button_a_presses := float64(delta2) / delta
			button_b_presses := float64(delta1) / delta

			if checkIfNumEqualItsFloat(button_a_presses) && checkIfNumEqualItsFloat(button_b_presses) {
				tokens += button_a_presses*3 + button_b_presses*1
			}
		} else {
			continue
		}
	}
	fmt.Printf("\nTokens : %d", int(tokens))

	return tokens
}

func Part2(fileSrc string) float64 {
	button_a_coeff_pairs, button_b_coeff_pairs, target_coord := processInput(fileSrc)

	tokens := 0.0

	for machine_num := 0; machine_num < len(button_a_coeff_pairs); machine_num++ {
		btn_a_coeffs := button_a_coeff_pairs[machine_num]
		btn_b_coeffs := button_b_coeff_pairs[machine_num]

		//Equations are like :
		// 1: a1(X) + a2(Y) = c1 => a1(X) + a2(Y) - c1 = 0
		// 2: b1(X) + b2(Y) = c2 => b1(X) + b2(Y) - c2 = 0

		c1 := target_coord[machine_num][0] + 10000000000000
		c2 := target_coord[machine_num][1] + 10000000000000

		a1 := btn_a_coeffs[0]
		b1 := btn_b_coeffs[0]

		a2 := btn_a_coeffs[1]
		b2 := btn_b_coeffs[1]

		if a1*b2 != a2*b1 {

			delta := float64(a1*b2 - a2*b1)
			delta1 := a1*c2 - a2*c1
			delta2 := -(b1*c2 - b2*c1)

			button_a_presses := float64(delta2) / delta
			button_b_presses := float64(delta1) / delta

			if checkIfNumEqualItsFloat(button_a_presses) && checkIfNumEqualItsFloat(button_b_presses) {
				tokens += button_a_presses*3 + button_b_presses*1
			}
		} else {
			continue
		}
	}
	fmt.Printf("\nTokens : %d\n", int(tokens))

	return tokens
}
