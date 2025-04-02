package main

import "fmt"

type EmptySquare struct {
	xCord int
	yCord int
}

func createBoard(size int) ([][]string, EmptySquare) {
	// Initialize the 2D array board && set the blank square position
	matrix := make([][]string, 0)
	num := 0
	vacantSquare := EmptySquare{xCord: size - 1, yCord: size - 1}

	for i := 0; i < size; i++ {
		row := make([]string, size)
		for j := 0; j < size; j++ {
			num++
			formatNum := fmt.Sprintf("%-3.2d", num) // Format number to have same num digits and trailing space 
			row[j] = formatNum
		}
		matrix = append(matrix, row)
	}
	// Add the identifier for the blank square
	matrix[vacantSquare.xCord][vacantSquare.yCord] = "-- "
	return matrix, vacantSquare
}

func displayBoard(matrix [][]string) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%s\n", row)
	}
	return result
}

func main() {
	board, blankSquare := createBoard(4)
	fmt.Print(displayBoard(board))
	fmt.Print(board[blankSquare.xCord][blankSquare.yCord])
}
