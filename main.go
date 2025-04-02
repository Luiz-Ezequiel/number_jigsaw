package main

import "fmt"

type EmptySquare struct {
	xCord int
	yCord int
}

func createBoard(size int) ([][]string, EmptySquare) {
	matrix := make([][]string, 0)
	num := 0
	vacant_square := EmptySquare{xCord: size - 1, yCord: size - 1}
	
	for range size {
		row := make([]string, size)
		for i := range size{
			num++
			formatNum := fmt.Sprintf("%.2d", num)
			row[i] = formatNum
		}
		matrix = append(matrix, row)
	}
	matrix[vacant_square.xCord][vacant_square.yCord] = "--"
	return matrix, vacant_square
}


func displayBoard(matrix [][]string) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%v\n", row)
	}
	return result
}

func main() {
	board, blankSquare := createBoard(5)
	fmt.Print(displayBoard(board))
	fmt.Print(board[blankSquare.xCord][blankSquare.yCord])
}
