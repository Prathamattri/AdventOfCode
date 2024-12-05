package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ProcessInput(inputPath string, regexStr string) [][]byte {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		panic("Unable to read input file")
	}

	re, err := regexp.CompilePOSIX(regexStr)

	arr := re.FindAll(data, -1)
	return arr
}

func Part1(inputPath string) {
	funcCallArr := ProcessInput(inputPath, "mul\\([0-9]+,[0-9]+\\)")

	re := regexp.MustCompile("[0-9]+")

	totalAnswer := 0
	for _, str_bytes := range funcCallArr {
		nums_as_bytes := re.FindAll(str_bytes, -1)
		num1, _ := strconv.Atoi(string(nums_as_bytes[0]))
		num2, _ := strconv.Atoi(string(nums_as_bytes[1]))

		totalAnswer += num1 * num2
	}
	fmt.Println(totalAnswer)
}

func Part2(inputPath string) {
	inputCallsArr := ProcessInput(inputPath, "do\\(\\)|don't\\(\\)|mul\\([0-9]+,[0-9]+\\)")

	totalAnswer := 0
	re := regexp.MustCompile("[0-9]+")
	sourceArrLen := len(inputCallsArr)
	for ind := 0; ind < sourceArrLen; ind++ {
		str_string := string(inputCallsArr[ind])

		switch str_string {
		case "don't()":
			for ; ind < sourceArrLen && str_string != "do()"; ind++ {
				str_string = string(inputCallsArr[ind])
			}
			ind--
			break
		case "do()":
			break
		default:
			nums_as_bytes := re.FindAll(inputCallsArr[ind], -1)
			num1, _ := strconv.Atoi(string(nums_as_bytes[0]))
			num2, _ := strconv.Atoi(string(nums_as_bytes[1]))

			totalAnswer += num1 * num2
		}
	}
	fmt.Println(totalAnswer)
}
