package main

import "fmt"

type EmptySquare struct{
	xCord 	int
	yCord	int
}

func createBoard(size int) ([][]string, EmptySquare){
	matrix := make([][]string, 0)
	num := 0
	vacant_square := EmptySquare{xCord: size - 1, yCord: size - 1}
	for range size{
		row := make([]string, 0)
		for range size{
			num++
			if num == size * size{
				row = append(row, "--")
				break
			}
			formatNum := fmt.Sprintf("%.2d", num)
			row = append(row, formatNum)
		}
		matrix = append(matrix, row )
	}
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
	board, blank := createBoard(3)
	fmt.Print(displayBoard(board))
	print(board[blank.xCord][blank.yCord])
}