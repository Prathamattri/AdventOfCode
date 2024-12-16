package day11

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cache = map[[2]int]int{}

func processInput(fileSrc string) []int {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	fileString := string(fileBytes)
	// Breaking lines from newline character into array of strings
	//The rules are divided from printing order by an empty line
	//so finding the empty line and breaking the array into 2 lines is the target here
	// array_string := strings.Split(strings.Split(fileString, "\n")[0], " ")
	array_string := strings.Split(strings.Split(fileString, "\n")[0], " ")

	num_int_arr := []int{}
	for _, num_str := range array_string {
		num_int, _ := strconv.Atoi(num_str)
		num_int_arr = append(num_int_arr, num_int)
	}
	return num_int_arr
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(float64(num))) + 1
}

func splitStone(stone int, dig_count *int) (int, int) {
	split_help_num := int(math.Pow10(*dig_count))
	right_part := stone % split_help_num
	stone /= split_help_num
	return stone, right_part
}
func Part1(fileSrc string) int {
	stones_arr := processInput(fileSrc)

	for i := 0; i < 25; i++ {
		for ind := 0; ind < len(stones_arr); ind++ {
			stone := stones_arr[ind]
			if stone == 0 {
				stones_arr[ind] = 1
			} else if dig := countDigits(stone); dig%2 == 0 && dig > 1 {
				split_index := dig / 2
				left_stone, right_stone := splitStone(stone, &split_index)
				stones_arr[ind] = right_stone
				stones_arr = slices.Insert(stones_arr, ind, left_stone)
				ind++
			} else {
				stones_arr[ind] *= 2024
			}
		}
	}

	fmt.Println(len(stones_arr))
	return len(stones_arr)
}

func getStoneCount(stone, blink_count int) int {
	if val, exists := cache[[2]int{stone, blink_count}]; exists {
		return val
	}
	if blink_count == 0 {
		return 1
	}

	if stone == 0 {
		result := getStoneCount(1, blink_count-1)
		cache[[2]int{stone, blink_count}] = result
		return result
	} else if dig := countDigits(stone); dig%2 == 0 && dig > 1 {
		split_index := dig / 2
		left_stone, right_stone := splitStone(stone, &split_index)

		result := getStoneCount(left_stone, blink_count-1) + getStoneCount(right_stone, blink_count-1)
		cache[[2]int{stone, blink_count}] = result
		return result
	} else {
		result := getStoneCount(stone*2024, blink_count-1)
		cache[[2]int{stone, blink_count}] = result
		return result
	}
}

func Part2(fileSrc string) int {
	stone_arr := processInput(fileSrc)
	total_count := 0
	blink_count := 75

	for _, stone := range stone_arr {
		total_count += getStoneCount(stone, blink_count)
	}

	fmt.Printf("Total Count : %d\n", total_count)
	return total_count
}
