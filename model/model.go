package model

import (
	"fmt"
)

type SudokuGrid [9][9]int16
type CandidateGenerator func(*SudokuGrid, int, int) []int16


func (grid *SudokuGrid) Solve(generator CandidateGenerator) (bool, *SudokuGrid) {
	if !grid.HasEmptyCell() {
		return true, grid
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] != 0 {
				continue
			}

			candidates := generator(grid, row, col)
			if len(candidates) == 0 {
				return false, nil
			}

			for _, value := range candidates {
				grid[row][col] = int16(value)
				solved, result := grid.Solve(generator)
				if solved {
					return solved, result
				} else {
					grid[row][col] = 0
				}
			}

			// After trying all potential candidates for current cell,
			// if the grid remains unsolved, end the curren search path
			return false, nil
		}
	}

	return false, nil
}


func (grid *SudokuGrid) Print() {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("|")
		for col := 0; col < 9; col++ {
			if grid[row][col] != 0 {
				fmt.Printf(" %d", grid[row][col])
			} else {
				fmt.Print("  ")
			}
			if col == 2 || col == 5 {
				fmt.Print(" |")
			}
		}
		fmt.Println(" |")
		if row == 2 || row == 5 {
			fmt.Println("+-------+-------+-------+")
		}
	}
	fmt.Println("+-------+-------+-------+")
}


func (grid *SudokuGrid) RowContains(row int, target int16) bool {
	gridRow := grid[row]
	for _, val := range gridRow {
		if val == target {
			return true
		}
	}

	return false
}


func (grid *SudokuGrid) ColumnContains(col int, target int16) bool {
	for row := 0; row < 9; row++ {
		if grid[row][col] == target {
			return true
		}
	}
	return false
}


func (grid *SudokuGrid) SubgridContains(row, col int, target int16) bool {
	startRow := 3 * (row / 3)
	stopRow := 3 * ((row / 3) + 1)
	
	startCol := 3 * (col / 3)
	stopCol := 3 * ((col / 3) + 1)

	return segmentContains(grid, startRow, stopRow, startCol, stopCol, target)
}


func (grid *SudokuGrid) HasEmptyCell() bool {
	return segmentContains(grid, 0, 9, 0, 9, 0)
}


func segmentContains(grid *SudokuGrid, startRow, stopRow, startCol, stopCol int, target int16) bool {
	for row := startRow; row < stopRow; row++ {
		for col := startCol; col < stopCol; col++ {
			if grid[row][col] == target {
				return true
			}
		}
	}

	return false
}
