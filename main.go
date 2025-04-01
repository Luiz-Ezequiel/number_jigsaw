package main

import "fmt"

func createBoard(size int) [][]string{
	matrix := make([][]string, 0)
	num := 0
	for range size{
		row := make([]string, 0)
		for range size{
			num++
			if num == size * size{
				row = append(row, "  ")
				break
			}
			formatNum := fmt.Sprintf("%.2d", num)
			row = append(row, formatNum)
		}
		matrix = append(matrix, row )
	}
	return matrix
}

func displayBoard(matrix [][]string) string {
	var result string
	for _, row := range matrix {
		result += fmt.Sprintf("%v\n", row)
	}
	return result
}

func printBoard(size int) {
	count:= size
	num := 0

	for range count {
		for range count{
		num++
		if num == count * count{
			break
		}
		fmt.Printf("[%02d] ", num)		
		}
	fmt.Println()
	}
	fmt.Println("---------------")
}

func main() {
	printBoard(3)
	fmt.Print(displayBoard(createBoard(4)))
}