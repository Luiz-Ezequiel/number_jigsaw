package main

import (
	"fmt"
)

type EmptySquare struct {
	yCord int
	xCord int
}

func move(board [][]string, blankSquare EmptySquare, size int, direction string) ([][]string, EmptySquare, error) {
	var newX, newY int  
	switch direction {
	case "s":
		if blankSquare.yCord == 0 {
			return board, blankSquare, fmt.Errorf("can't move down")
		}
		newY, newX = blankSquare.yCord - 1, blankSquare.xCord
	case "w":
		if blankSquare.yCord == size - 1 {
			return board, blankSquare, fmt.Errorf("can't move up")
		}
		newY, newX = blankSquare.yCord + 1, blankSquare.xCord
	case "d":
		if blankSquare.xCord == 0 {
			return board, blankSquare, fmt.Errorf("can't move left")
		}
		newY, newX = blankSquare.yCord, blankSquare.xCord - 1
	case "a":
		if blankSquare.xCord == size - 1 {
			return board, blankSquare, fmt.Errorf("can't move right")
		}
		newY, newX = blankSquare.yCord, blankSquare.xCord + 1
	default:
		fmt.Println("Invalid direction")
		return board, blankSquare, nil
	}

	board[blankSquare.yCord][blankSquare.xCord] = board[newY][newX]
	board[newY][newX] = "-- "
	blankSquare.yCord, blankSquare.xCord = newY, newX
	return board, blankSquare, nil
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
	size := 3
	board, blankSquare := createBoard(size)
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err := move(board, blankSquare, size, "s")
	if err != nil{
		fmt.Print(err)
		return
	}
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err = move(board, blankSquare, size, "s")
	if err != nil{
		fmt.Print(err)
		return
	}
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err = move(board, blankSquare, size, "s")
	if err != nil{
		fmt.Print(err)
		return
	}
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err = move(board, blankSquare, size, "d")
	if err != nil{
		return
	}
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err = move(board, blankSquare, size, "a")
	if err != nil{
		return
	}
	fmt.Print(displayBoard(board))
	fmt.Println("___________________")
	board, blankSquare, err = move(board, blankSquare, size, "w")
	if err != nil{
		return
	}
	fmt.Print(displayBoard(board))
	// fmt.Print(board[blankSquare.yCord-1][blankSquare.xCord])
}
