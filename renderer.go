package main

import (
	"fmt"
)

var maze = [][]int{
	{1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 1},
	{1, 0, 1, 0, 1, 0, 1},
	{1, 0, 1, 0, 0, 0, 1},
	{1, 0, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 2, 1},
	{1, 1, 1, 1, 1, 1, 1},
}

var playerX, playerY = 1, 5

func printMaze() {
	for y, row := range maze {
		for x, col := range row {
			switch {
			case playerX == x && playerY == y:
				fmt.Print("P ") 
			case col == 0:
				fmt.Print("  ") 
			case col == 1:
				fmt.Print("█ ") 
			case col == 2:
				fmt.Print("E ") 
			}
		}
		fmt.Println()
	}
}

func main() {
	printMaze() 
}