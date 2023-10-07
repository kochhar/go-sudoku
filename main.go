package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"example/sudoku/model"
)


func getCandidates(grid *model.SudokuGrid, row, col int) []int16 {
	candidates := []int16{}

	for i := 1; i <= 9; i++ {
		if grid.RowContains(row, int16(i)) {
			continue
		}

		if grid.ColumnContains(col, int16(i)) {
			continue
		}

		if grid.SubgridContains(row, col, int16(i)) {
			continue
		}

		candidates = append(candidates, int16(i))
	}

	return candidates
}


func readGrid(input string) model.SudokuGrid {
	grid := model.SudokuGrid{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			for true {
				scanner.Scan()
				trimmed := strings.TrimSpace(scanner.Text())
				if len(trimmed) > 0 {
					break
				}
			}
			
			i1, _ := strconv.Atoi(scanner.Text())
			grid[row][col] = int16(i1)
		}
	}
	return grid
}


func main() {
	grid := readGrid(`
		1 0 3 0 5 0 7 0 9 
		7 8 9 1 0 3 0 5 0 
		0 5 0 7 0 9 1 0 3 
		9 1 0 3 0 5 0 7 0 
		0 7 0 9 1 0 3 0 5 
		3 0 5 0 7 0 9 1 0 
		8 9 1 0 3 0 5 0 7 
		5 0 7 0 9 1 0 3 0 
		0 3 0 5 0 7 0 9 1`)
	grid.Print()
	solved, solution := grid.Solve(getCandidates)
	if solved {
		solution.Print()
	} else {
		fmt.Println("Could not solve grid")
	}
}

