package day4

import (
	"bytes"
	"fmt"
	"os"
)

func processInput(fileSrc string) [][]byte {
	data, err := os.ReadFile(fileSrc)
	if err != nil {
		panic("Unable to read input file")
	}

	rows := bytes.Split(data, []byte("\n"))
	splittedRows := [][]byte{}
	for _, row := range rows {
		temp_row := bytes.Split(row, []byte(""))
		splittedRows = append(splittedRows, temp_row...)
	}

	return rows
}

func checkVertical(grid [][]byte, row, col, rowMax int, word string, count *int) {
	temp_str := make([]byte, len(word))
	//Check South
	if row+len(word) <= rowMax {
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row+ind][col]
		}
		if string(temp_str) == word {
			*count++
		}
	}
	//Check North
	clear(temp_str)
	if row-(len(word)-1) >= 0 {
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row-ind][col]
		}
		if string(temp_str) == word {
			*count++
		}
	}
}
func checkHorizontal(grid [][]byte, row, col, colMax int, word string, count *int) {
	// Checking East
	temp_str := make([]byte, len(word))
	if col+len(word) <= colMax {
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row][col+ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
	// Checking West
	if col-(len(word)-1) >= 0 {
		clear(temp_str)
		temp_str = make([]byte, len(word))
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row][col-ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
}

func checkDiagonal(grid [][]byte, row, col, rowMax, colMax int, word string, count *int) {

	temp_str := make([]byte, len(word))
	if col+len(word) <= colMax && row+len(word) <= rowMax {
		// Checking South-East
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row+ind][col+ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
	if col-(len(word)-1) >= 0 && row+len(word) <= rowMax {
		// Checking South-West
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row+ind][col-ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
	if col-(len(word)-1) >= 0 && row-(len(word)-1) >= 0 {
		// Checking North-West
		clear(temp_str)
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row-ind][col-ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
	if col+len(word) <= colMax && row-(len(word)-1) >= 0 {
		// Checking North-East
		clear(temp_str)
		for ind := 0; ind < len(word); ind++ {
			temp_str[ind] = grid[row-ind][col+ind]
		}
		if string(temp_str) == word {
			*count++
		}
	}
}

func Part1(fileSrc string) {
	wordGrid := processInput(fileSrc)
	word := "XMAS"
	wordMatchCount := 0
	rowsLen := len(wordGrid) - 1
	for ind := 0; ind < len(wordGrid)-1; ind++ {
		colsLen := len(wordGrid[ind])

		for ind2, letter := range wordGrid[ind] {
			if string(letter) != "X" {
				continue
			}
			checkHorizontal(wordGrid, ind, ind2, colsLen, word, &wordMatchCount)
			checkVertical(wordGrid, ind, ind2, rowsLen, word, &wordMatchCount)
			checkDiagonal(wordGrid, ind, ind2, rowsLen, colsLen, word, &wordMatchCount)
		}
	}
	fmt.Println(wordMatchCount)

}

func Part2(fileSrc string) {
	wordGrid := processInput(fileSrc)

	xmas_count := 0

	word1 := "MAS"
	word2 := "SAM"
	for row := 1; row < len(wordGrid)-2; row++ {
		for col := 1; col < len(wordGrid[row])-1; col++ {
			letter := string(wordGrid[row][col])
			diag1 := make([]byte, 3)
			diag2 := make([]byte, 3)
			if letter == "A" {
				//NW-SE diagonal
				diag1[0] = wordGrid[row-1][col-1]
				diag1[1] = wordGrid[row][col]
				diag1[2] = wordGrid[row+1][col+1]

				//NE-SW diagonal
				diag2[0] = wordGrid[row+1][col-1]
				diag2[1] = wordGrid[row][col]
				diag2[2] = wordGrid[row-1][col+1]
			}
			diagStr1 := string(diag1)
			diagStr2 := string(diag2)
			if (diagStr1 == word1 || diagStr1 == word2) && (diagStr2 == word1 || diagStr2 == word2) {
				xmas_count++
			}
		}
	}

	fmt.Println(xmas_count)
}
