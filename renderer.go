package main

import (
	"fmt"
	"os"
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

var playerX, playerY int = 1, 5

func printMaze() {
	for y, row := range maze {
		for x, col := range row {
			if playerX == x && playerY == y {
				fmt.Print("P ")
			} else {
				switch col {
				case 0:
					fmt.Print("  ")
				case 1:
					fmt.Print("█ ")
				case 2:
					fmt.Print("E ")
				}
			}
		}
		fmt.Println()
	}
}

func setupEnv() {
}

func main() {
	setupEnv()
	printMaze()
}