package main

import "fmt"

func main() {
	var matrix [][]int
	matrix = createMatrix(5,5)
	initMatrix(matrix)
	printMatrix(matrix)
}

func createMatrix(rows int, cols int) [][]int{
	m := make([][]int, rows)
	for r := range m {
		m[r] = make([]int, cols)
	}
	return m
}

func printMatrix(matrix [][]int) {
	for x, col := range matrix{
		for y, _ := range col{
			fmt.Print(matrix[x][y], " ")
		}
		fmt.Println("")
	}
}

func initMatrix(matrix [][]int){
	for x, col := range matrix{
		for y, _ := range  col{
			matrix[x][y] =  x*y
		}
	}
}
