package main

import (
	"os"

	"github.com/01-edu/z01"
)

func solveSudoku(grid [][]rune) bool {
	row, col := 0, 0
	var found bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == '0' {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num

			if solveSudoku(grid) {
				return true
			}

			grid[row][col] = '0'
		}
	}

	return false
}

func isSafe(grid [][]rune, row, col int, num rune) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	boxRow := row - row%3
	boxCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

func Safe(grid [][]rune, row, col int, num rune) bool {
	if grid[row][col] != '0' {
		// fmt.Printf("A: %v %v %v \n", row, col, num)
		counter := 0
		for i := 0; i < 9; i++ {
			if grid[row][i] == num || grid[i][col] == num {
				counter++
			}
			if counter > 2 {
				return false
			}
		}
		// fmt.Printf("B: %v %v %v \n", row, col, num)
		boxRow := row - row%3
		boxCol := col - col%3
		counter = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if grid[boxRow+i][boxCol+j] == num {
					counter++
				}
				if counter > 1 {
					return false
				}
			}
		}
		// fmt.Printf("C: %v %v %v \n", row, col, num)
	}

	return true
}

func PrintError() {
	s := []rune("Error")

	for i := 0; i < len(s); i++ {
		z01.PrintRune(s[i])
	}
	z01.PrintRune('\n')
}

func main() {
	s := os.Args[1:]

	if len(s) != 9 {
		PrintError()
		return
	}

	for i := 0; i < len(s); i++ {
		if len(s[i]) != 9 {
			PrintError()
			return
		}
	}

	grid := make([][]rune, len(s))

	for i := 0; i < len(s); i++ {
		grid[i] = make([]rune, len(s[i]))
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == '.' {
				grid[i][j] = '0'
			} else if s[i][j] >= '1' && s[i][j] <= '9' {
				grid[i][j] = rune(s[i][j])
			} else {
				PrintError()
				return
			}
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if !Safe(grid, i, j, grid[i][j]) {
				PrintError()
				return
			}
		}
	}

	if solveSudoku(grid) {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(s[i]); j++ {
				z01.PrintRune(rune(grid[i][j]))
				if j+1 != len(s[i]) {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		PrintError()
	}
}
