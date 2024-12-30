package day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isElementIn(elem int, slice []int) bool {
	for _, num := range slice {
		if elem == num {
			return true
		}
	}
	return false
}

func processInput(fileSrc string) (map[int][]int, [][]int) {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}
	fileString := string(fileBytes)
	// Breaking lines from newline character into array of strings
	//The rules are divided from printing order by an empty line
	//so finding the empty line and breaking the array into 2 parts is the target here
	parts := strings.Split(fileString, "\n\n")

	rules := []string{}
	rules = append(rules, strings.Split(parts[0], "\n")...)

	rules_map := make(map[int][]int)

	// Creating a map to simplify the rules
	for _, pair := range rules {
		parts := strings.Split(pair, "|")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		rules_map[num1] = append(rules_map[num1], num2)
	}
	for _, rules := range rules_map {
		slices.Sort(rules)
	}

	printingOrderArray := [][]int{}
	printOrderList := strings.Split(parts[1], "\n")
	for _, printOrder := range printOrderList {
		temp_int_list := []int{}
		num_list_str := strings.Split(printOrder, ",")
		if printOrder != "" {
			for _, num_str := range num_list_str {
				num_int, _ := strconv.Atoi(num_str)
				temp_int_list = append(temp_int_list, num_int)
			}
			printingOrderArray = append(printingOrderArray, temp_int_list)
		}
	}

	return rules_map, printingOrderArray
}

func Part1(fileSrc string) int {
	rules_map, printingOrderArray := processInput(fileSrc)

	middle_term_sum := 0
	for _, order := range printingOrderArray {
		order_len := len(order)
		isValid := true
		for i := order_len - 1; i >= 1; i-- {
			key := order[i]
			for j := i - 1; j >= 0; j-- {
				_, isPresent := slices.BinarySearch(rules_map[key], order[j])
				if isPresent {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}

		if isValid {
			var mid_index int = len(order) / 2
			middle_term_sum += order[mid_index]
		}
	}
	fmt.Println(middle_term_sum)
	return middle_term_sum
}
func Part2(fileSrc string) {
	rules_map, printOrderArray := processInput(fileSrc)

	middle_term_sum := 0
	for order_ind, order := range printOrderArray {
		order_len := len(order)
		isValid := true
		for i := order_len - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				key := printOrderArray[order_ind][i]
				if _, ok := rules_map[key]; !ok {
					fmt.Println(key)
					break
				}
				_, isPresent := slices.BinarySearch(rules_map[key], printOrderArray[order_ind][j])
				if isPresent {
					isValid = false
					for k := j; k < i; k++ {
						temp := order[k]
						printOrderArray[order_ind][k] = printOrderArray[order_ind][k+1]
						printOrderArray[order_ind][k+1] = temp
					}
					j = i
				}
			}
		}
		if !isValid {
			var mid_index int = order_len / 2
			middle_term_sum += printOrderArray[order_ind][mid_index]
		}
	}

	fmt.Println(middle_term_sum)
}
