package day1

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	item_1 int
	item_2 int
}

func (pair *Pair) getDistance() int {
	dist := pair.item_1 - pair.item_2
	if dist < 0 {
		return (dist * -1)
	}
	return dist
}

func (pair *Pair) printPair() {
	fmt.Printf("%d:%d\n", pair.item_1, pair.item_2)
}

// --- Day 1: Historian Hysteria ---

func ProcessInput(fileSrc string) ([]int, []int) {
	data, err := os.ReadFile(fileSrc)

	if err != nil {
		fmt.Println("Error encountered while reading input file")
	}

	dataAsStr := string(data)

	list1 := []int{}
	list2 := []int{}

	strList := strings.Split(dataAsStr, "\n")
	re := regexp.MustCompile("\\s+")
	for i := 0; i < len(strList)-1; i++ {
		numbers := re.Split(strList[i], -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

// --- Part1 ---
func Part1(fileSrc string) {

	list1, list2 := ProcessInput(fileSrc)
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	pairList := []Pair{}
	var totalDistance int = 0

	for i := 0; i < len(list1); i++ {
		pairList = append(pairList, Pair{list1[i], list2[i]})
	}
	// Calculating distance
	for i := 0; i < len(pairList); i++ {
		totalDistance += pairList[i].getDistance()
	}
	fmt.Println(totalDistance)
}

/*
		Part2

	 Calculate a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.
*/
func Part2(fileSrc string) {
	list1, list2 := ProcessInput(fileSrc)
	similarityScore := 0

	countList := make(map[int]int)

	for _, item1 := range list1 {
		if _, ok := countList[item1]; ok {
			continue
		}
		var count int = 0
		for _, item2 := range list2 {
			if item1^item2 == 0 {
				count++
			}
		}
		countList[item1] = count
	}

	for key, val := range countList {
		similarityScore += (key * val)
	}

	fmt.Println(similarityScore)
}
