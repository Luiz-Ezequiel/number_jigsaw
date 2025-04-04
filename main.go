package main

import (
	"fmt"
	"math/rand"
)

type EmptySquare struct {
	yCord int
	xCord int
}

const validMoves = "wasd"

func randomMove() string {
	b := validMoves[rand.Intn(len(validMoves))]
	return string(b)
  }

// Move numbers around the empty square based on imput
func move(board [][]string, blankSquare EmptySquare, size int, direction string) ([][]string, EmptySquare, error) {
	var newX, newY int  
	switch direction {
	case "w":
		if blankSquare.yCord == 0 {
			return board, blankSquare, fmt.Errorf("can't move down")
		}
		newY, newX = blankSquare.yCord - 1, blankSquare.xCord
	case "s":
		if blankSquare.yCord == size - 1 {
			return board, blankSquare, fmt.Errorf("can't move up")
		}
		newY, newX = blankSquare.yCord + 1, blankSquare.xCord
	case "a":
		if blankSquare.xCord == 0 {
			return board, blankSquare, fmt.Errorf("can't move left")
		}
		newY, newX = blankSquare.yCord, blankSquare.xCord - 1
	case "d":
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

// Creates a solved puzzle and compares it to your board
func checkBoard(board [][]string, size int) bool {
	result := true
	puzzleSolved, _ := createBoard(size)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if puzzleSolved[i][j] != board[i][j]{
				result = false
			}
		}
	}
	return result
}

func gameLoop(board [][]string, blankSquare EmptySquare, size int) ([][]string, bool) {
	fmt.Println(displayBoard(board))

	var err error 
	input := ""
	isSolved := false

	for input != "q" {
		fmt.Scan(&input)
		 
		if input == "w" || input == "a" || input == "s" || input == "d" {
			board, blankSquare, err = move(board, blankSquare, size, input)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(displayBoard(board))
			isSolved = checkBoard(board, size)
			if isSolved {
				return board, isSolved
			}
		}
	}
	return board, isSolved
}

func shuffleBoard(board [][]string, blankSquare EmptySquare, size int) ([][]string, EmptySquare) {
	for i := 0; i < 1500; i++ {
		randomMove := randomMove()	
		board, blankSquare, _ = move(board, blankSquare, size, randomMove)
	} 
	return board, blankSquare
}

func main() {
	size := 0
	fmt.Print("Select the size of the board(>=3): ")
	for size <= 2 {
		fmt.Scan(&size)
	}

	board, blankSquare := createBoard(size)
	board, blankSquare = shuffleBoard(board, blankSquare, size) 

	_, result := gameLoop(board, blankSquare, size)
	if result {
		fmt.Print("Congratulations!! You won :)")
	} else {
		fmt.Print("Quitter. :P!")
	}
}
