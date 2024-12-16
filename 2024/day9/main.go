package day9

import (
	"fmt"
	"os"
	"strings"
)

func processInput(fileSrc string) []rune {
	fileBytes, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Error while reading input file")
	}

	fileString := strings.Split(string(fileBytes), "\n")[0]
	return []rune(fileString)
}

func calcChecksum(diskmap []rune) int {
	checksum := 0
	for ind, char := range diskmap {
		if char == '.' {
			continue
		}
		checksum += (ind * int(char-'0'))
	}
	return checksum
}

func Part1(fileSrc string) int {
	diskmap := processInput(fileSrc)
	memory := []rune{}
	file_num := 0
	for ind, count := range diskmap {
		num := int(count - '0')
		if ind&1 == 0 {
			//even num / it is a file block size
			file_num_rune := rune('0' + file_num)
			for ; num > 0; num-- {
				memory = append(memory, file_num_rune)
			}
			file_num++

		} else {
			//odd num / empty space size
			for ; num > 0; num-- {
				memory = append(memory, '.')
			}
		}
	}
	ptr2 := len(memory) - 1
	for ptr1 := 0; ptr1 < ptr2; ptr1++ {
		if memory[ptr1] == '.' {
			for memory[ptr2] == '.' {
				ptr2--
			}
			temp_rune := memory[ptr1]
			memory[ptr1] = memory[ptr2]
			memory[ptr2] = temp_rune

			ptr2--
		}
	}

	checksum := calcChecksum(memory)

	fmt.Printf("\nPart1 Checksum : %d\n", checksum)
	return checksum

}

func Part2(fileSrc string) int {
	diskmap := processInput(fileSrc)

	memory := []rune{}
	file_num := 0
	for ind, count := range diskmap {
		num := int(count - '0')
		if ind&1 == 0 {
			//even num / it is a file block size
			file_num_rune := rune('0' + file_num)
			for ; num > 0; num-- {
				memory = append(memory, file_num_rune)
			}
			file_num++

		} else {
			//odd num / empty space size
			for ; num > 0; num-- {
				memory = append(memory, '.')
			}
		}
	}

	memory_len := len(memory)

	var start_ind int
	ind_accumulator := memory_len - 1

	if len(diskmap)&1 == 0 {
		start_ind = len(diskmap) - 1
		ind_accumulator -= int(diskmap[start_ind] - '0')
	} else {
		start_ind = len(diskmap)
	}

	copy_diskmap := append([]rune{}, diskmap...)
	for ptr1 := start_ind - 1; ptr1 > 0; ptr1 -= 2 {
		starting_rune := memory[ind_accumulator]

		// find free space in diskmap

		req_free_space := int(copy_diskmap[ptr1] - '0')
		mem_index_accumulator := 0
		is_space_available := false

		for i := 1; i < ptr1; i += 2 {

			if req_free_space <= int(diskmap[i]-'0') {

				mem_index_accumulator += int(diskmap[i-1] - '0')

				diskmap[ptr1-1] = rune('0' + int(diskmap[ptr1-1]-'0') + int(diskmap[ptr1]-'0'))
				diskmap[ptr1] = '0'

				diskmap[i] = rune('0' + int(diskmap[i]-'0') - req_free_space)
				diskmap[i-1] = rune('0' + int(diskmap[i-1]-'0') + req_free_space)

				is_space_available = true

				break
			}
			mem_index_accumulator += (int(diskmap[i]-'0') + int(diskmap[i-1]-'0'))
		}

		if is_space_available {
			//swap in memory
			for i := 0; i < req_free_space; i++ {
				memory[ind_accumulator-i] = '.'
				memory[mem_index_accumulator+i] = starting_rune
			}
		}
		ind_accumulator -= (int(copy_diskmap[ptr1]-'0') + int(copy_diskmap[ptr1-1]-'0'))
	}
	checksum := calcChecksum(memory)

	fmt.Printf("\nPart2 Checksum : %d\n", checksum)
	return checksum
}
