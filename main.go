package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Maze struct {
	grid [][]string
	playerX, playerY int
}

func NewMaze() *Maze {
	return &Maze{
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
		playerX: 9,
		playerY: 8,
	}
}

func (m *Maze) PrintMaze() {
	for _, row := range m.grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

func (m *Maze) MovePlayer(direction string) {
	m.grid[m.playerX][m.playerY] = " "
	switch direction {
	case "UP":
		if m.playerX > 0 && m.grid[m.playerX-1][m.playerY] != "#" {
			m.playerX--
		}
	case "DOWN":
		if m.playerX < len(m.grid)-1 && m.grid[m.playerX+1][m.playerY] != "#" {
			m.playerX++
		}
	case "LEFT":
		if m.playerY > 0 && m.grid[m.playerX][m.playerY-1] != "#" {
			m.playerY--
		}
	case "RIGHT":
		if m.playerY < len(m.grid[0])-1 && m.grid[m.playerX][m.playerY+1] != "#" {
			m.playerY++
		}
	}
	m.grid[m.playerX][m.playerY] = "P"
}

func main() {
	maze := NewMaze()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		maze.PrintMaze()
		fmt.Println("Move with UP, DOWN, LEFT, RIGHT or type EXIT to quit:")

		scanner.Scan()
		input := strings.ToUpper(scanner.Text())

		if input == "EXIT" {
			fmt.Println("Exiting game. Thank you for playing!")
			break
		}

		maze.MovePlayer(input)
	}
}