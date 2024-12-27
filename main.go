package main

import (
	"AdventOfCode/2024/day1"
	"AdventOfCode/2024/day10"
	"AdventOfCode/2024/day11"
	"AdventOfCode/2024/day12"
	"AdventOfCode/2024/day13"
	"AdventOfCode/2024/day14"
	"AdventOfCode/2024/day15"
	"AdventOfCode/2024/day2"
	"AdventOfCode/2024/day3"
	"AdventOfCode/2024/day4"
	"AdventOfCode/2024/day5"
	"AdventOfCode/2024/day6"
	"AdventOfCode/2024/day7"
	"AdventOfCode/2024/day8"
	"AdventOfCode/2024/day9"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	year := flag.Int("y", 2024, "Year Number of AoC [2024]")
	day := flag.Int("d", 0, "Day Number of AoC [1-25]")
	partNum := flag.Int("p", 0, "Part Number of day's problem of AoC [1 or 2]")
	test := flag.Bool("test", false, "If the code need to be tested, this flag is used")
	flag.Usage = func() {
		fmt.Println("Advent Of Code 2024\n\n[Options]:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if !(*day > 0 && *day <= 25) || *partNum == 0 || *partNum > 2 || *year == 0 || *year < 2024 {
		fmt.Printf("year : %d \nday : %d\npart : %d\ntest mode : %v", *year, *day, *partNum, *test)
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
	case 5:
		if *test {
			day5.Test(testInputSrc)
		} else if *partNum == 1 {
			day5.Part1(inputSrc)
		} else if *partNum == 2 {
			day5.Part2(inputSrc)
		}
		break
	case 6:
		if *test {
			day6.Test(testInputSrc)
		} else if *partNum == 1 {
			day6.Part1(inputSrc)
		} else if *partNum == 2 {
			day6.Part2(inputSrc)
		}
		break
	case 7:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day7.Part1(inputSrc)
		} else if *partNum == 2 {
			day7.Part2(inputSrc)
		}
		break
	case 8:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day8.Part1(inputSrc)
		} else if *partNum == 2 {
			day8.Part2(inputSrc)
		}
		break
	case 9:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day9.Part1(inputSrc)
		} else if *partNum == 2 {
			day9.Part2(inputSrc)
		}
		break
	case 10:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day10.Part1(inputSrc)
		} else if *partNum == 2 {
			day10.Part2(inputSrc)
		}
		break
	case 11:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day11.Part1(inputSrc)
		} else if *partNum == 2 {
			day11.Part2(inputSrc)
		}
		break
	case 12:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day12.Part1(inputSrc)
		} else if *partNum == 2 {
			day12.Part2(inputSrc)
		}
		break
	case 13:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day13.Part1(inputSrc)
		} else if *partNum == 2 {
			day13.Part2(inputSrc)
		}
		break
	case 14:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day14.Part1(inputSrc)
		} else if *partNum == 2 {
			day14.Part2(inputSrc)
		}
		break
	case 15:
		if *test {
			inputSrc = testInputSrc
		}
		if *partNum == 1 {
			day15.Part1(inputSrc)
		} else if *partNum == 2 {
			day15.Part2(inputSrc)
		}
		break
	default:
		fmt.Println("this day hasn't been done yet")
	}

}
