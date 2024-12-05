package main

import (
	"AdventOfCode/2024/day1"
	"AdventOfCode/2024/day2"
	"AdventOfCode/2024/day3"
	"AdventOfCode/2024/day4"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	year := flag.Int("y", 0, "Year Number of AoC [2024]")
	day := flag.Int("d", 0, "Day Number of AoC [1-25]")
	partNum := flag.Int("p", 0, "Part Number of day's problem of AoC [1 or 2]")
	test := flag.Bool("test", false, "If the code need to be tested, this flag is used")
	flag.Usage = func() {
		fmt.Println("Advent Of Code 2024\n\n[Options]:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if !(*day > 0 && *day <= 25) || *partNum == 0 || *partNum > 2 || *year == 0 || *year < 2024 {
		flag.Usage()
		os.Exit(-1)
	}

	inputSrc := "./" + strconv.Itoa(*year) + "/day" + strconv.Itoa(*day) + "/input.txt"
	testInputSrc := "./" + strconv.Itoa(*year) + "/day" + strconv.Itoa(*day) + "/test_input.txt"
	switch *day {
	case 1:
		fmt.Println(inputSrc)
		if *test {
			day1.Test(testInputSrc)
		} else if *partNum == 1 {
			day1.Part1(inputSrc)
		} else if *partNum == 2 {
			day1.Part2(inputSrc)
		}
		break
	case 2:
		if *test {
			day2.Test(testInputSrc)
		} else if *partNum == 1 {
			day2.Part1(inputSrc)
		} else if *partNum == 2 {
			day2.Part2(inputSrc)
		}
		break
	case 3:
		if *test {
			day3.Test(testInputSrc)
		} else if *partNum == 1 {
			day3.Part1(inputSrc)
		} else if *partNum == 2 {
			day3.Part2(inputSrc)
		}
		break
	case 4:
		if *test {
			day4.Test(testInputSrc)
		} else if *partNum == 1 {
			day4.Part1(inputSrc)
		} else if *partNum == 2 {
			day4.Part2(inputSrc)
		}
		break
	default:
		fmt.Println("this day hasn't been done yet")
	}

}
