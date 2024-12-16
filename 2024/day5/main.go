package day5

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pQueue struct {
	num      int
	priority int
}

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
		rules_map[num2] = append(rules_map[num2], num1)
	}
	for num, rules := range rules_map {
		slices.Sort(rules)
		fmt.Printf("%d : %v\n", num, rules)
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

	// Creating a priority list based on rules_map
	priority_list := []int{}
	for key, less_priority_elements := range rules_map {
		index := 0
		for ind, element := range priority_list {
			if isElementIn(element, less_priority_elements) {
				// means that element is of less priority than the key
				index = ind
				break
			} else {
				index = ind + 1
			}
		}
		priority_list = slices.Insert(priority_list, index, key)
	}
	priority_map := make(map[int]int)
	for index, element := range priority_list {
		priority_map[element] = len(priority_list) - index
	}
	// if any page is not present while checking the pages order inside the priority_list
	// it automatically has the least priority

	fmt.Println(priority_map)
	middle_term_sum := 0
	for _, order := range printingOrderArray {
		most_recent_priority := priority_map[order[0]]
		isValid := true
		for _, page_num := range order {
			priority, _ := priority_map[page_num]
			if priority > most_recent_priority {
				isValid = false
				break
			} else {
				most_recent_priority = priority
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

}
