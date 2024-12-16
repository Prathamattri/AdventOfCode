package day8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processInput(fileSrc string) ([]int, [][]int) {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	fileString := string(fileBytes)
	// Breaking lines from newline character into array of strings
	//The rules are divided from printing order by an empty line
	//so finding the empty line and breaking the array into 2 lines is the target here
	lines := strings.Split(fileString, "\n")

	test_values := []int{}

	lines = lines[:len(lines)-1]
	operands_parent_arr := [][]int{}
	for _, line := range lines {
		//Splitting a line from ": " to get two parts
		//one containing target val
		//other containing array of operands

		//target value part
		line_parts := strings.Split(line, ": ")
		num, err := strconv.Atoi(line_parts[0])
		if err != nil {
			panic("Error converting string to integer")
		}
		test_values = append(test_values, num)

		//operands array part
		operands := strings.Split(line_parts[1], " ")
		operands_int := []int{}
		for _, operand := range operands {
			num, _ := strconv.Atoi(operand)
			operands_int = append(operands_int, num)
		}
		operands_parent_arr = append(operands_parent_arr, operands_int)

	}
	return test_values, operands_parent_arr
}

// Helper function for recursion in Part1
func isProducable(array []int, arr_len, test_val, index, curr_val int) bool {
	if index == arr_len {
		return curr_val == test_val
	}
	// Try addition
	curr_val += array[index]
	isAdd := isProducable(array, arr_len, test_val, index+1, curr_val)
	curr_val -= array[index]
	if isAdd {
		return true
	}

	// Try multiplication
	if index == 0 {
		curr_val = 1
	}
	curr_val *= array[index]
	isMul := isProducable(array, arr_len, test_val, index+1, curr_val)

	return isAdd || isMul
}

// Helper function for recursion in Part2
func isProducable2(array []int, arr_len, test_val, index, curr_val int) bool {
	if index == arr_len {
		return curr_val == test_val
	}
	// Try concatenation
	temp_curr_val, _ := strconv.Atoi(strconv.Itoa(curr_val) + strconv.Itoa(array[index]))
	isConcat := isProducable2(array, arr_len, test_val, index+1, temp_curr_val)
	if isConcat {
		return true
	}

	// Try addition
	curr_val += array[index]
	isAdd := isProducable2(array, arr_len, test_val, index+1, curr_val)
	curr_val -= array[index]

	if isAdd {
		return true
	}
	// Try multiplication
	if index == 0 {
		curr_val = 1
	}
	curr_val *= array[index]
	isMul := isProducable2(array, arr_len, test_val, index+1, curr_val)

	return isAdd || isMul || isConcat
}
func Part1(fileSrc string) int {
	total_calibration_res := 0
	target_result_arr, operands_parent_arr := processInput(fileSrc)

	for index, operands := range operands_parent_arr {
		if isProducable(operands, len(operands), target_result_arr[index], 0, 0) {
			total_calibration_res += target_result_arr[index]
		}
	}

	fmt.Println(total_calibration_res)
	return total_calibration_res
}

func Part2(fileSrc string) int {
	total_calibration_res := 0
	target_result_arr, operands_parent_arr := processInput(fileSrc)

	for index, operands := range operands_parent_arr {
		if isProducable2(operands, len(operands), target_result_arr[index], 0, 0) {
			total_calibration_res += target_result_arr[index]
		}
	}

	fmt.Println(total_calibration_res)
	return total_calibration_res
}
