package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MazeGame struct {
	grid      [][]string
	playerRow, playerCol int
}

func NewMazeGame() *MazeGame {
	return &MazeGame{
		grid: [][]string{
			{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
			{"#", " ", " ", " ", "#", " ", " ", " ", " ", "#"},
			{"#", " ", "#", " ", "#", " ", "#", "#", " ", "#"},
			{"#", " ", "#", " ", " ", " ", "#", " ", " ", "#"},
			{"#", " ", "#", "#", "#", " ", "#", " ", "#", "#"},
			{"#", " ", " ", " ", "#", " ", " ", " ", " ", "#"},
			{"#", "#", "#", " ", "#", "#", "#", "#", " ", "#"},
			{"#", " ", " ", " ", " ", " ", " ", "#", " ", "#"},
			{"#", " ", "#", "#", "#", "#", " ", "#", " ", "#"},
			{"#", "#", "#", "#", "#", "#", "#", "#", "P", "#"},
		},
		playerRow: 9,
		playerCol: 8,
	}
}

func (game *MazeGame) DisplayMaze() {
	for _, row := range game.grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func (game *MazeGame) MovePlayer(direction string) {
	game.grid[game.playerRow][game.playerCol] = " "
	switch direction {
	case "UP":
		if game.playerRow > 0 && game.grid[game.playerRow-1][game.playerCol] != "#" {
			game.playerRow--
		}
	case "DOWN":
		if game.playerRow < len(game.grid)-1 && game.grid[game.playerRow+1][game.playerCol] != "#" {
			game.playerRow++
		}
	case "LEFT":
		if game.playerCol > 0 && game.grid[game.playerRow][game.playerCol-1] != "#" {
			game.playerCol--
		}
	case "RIGHT":
		if game.playerCol < len(game.grid[0])-1 && game.grid[game.playerRow][game.playerCol+1] != "#" {
			game.playerCol++
		}
	}
	game.grid[game.playerRow][game.playerCol] = "P"
}

func main() {
	mazeGame := NewMazeGame()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		mazeGame.DisplayMaze()
		fmt.Println("Move with UP, DOWN, LEFT, RIGHT or type EXIT to quit:")

		scanner.Scan()
		input := strings.ToUpper(scanner.Text())

		if input == "EXIT" {
			fmt.Println("Exiting game. Thank you for playing!")
			break
		}

		mazeGame.MovePlayer(input)
	}
}