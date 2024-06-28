package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	StatusRunning = iota
	StatusGameOver // Assuming you might want this for further development.
	StatusWon
)

type Coordinates struct {
	X, Y int
}

type Explorer struct {
	Coordinates
}

type MazeState struct {
	Explorer Explorer
	Maze     [][]int
	Status   int
}

func generateMaze(rows, columns int) [][]int {
	maze := make([][]int, rows)
	for i := range maze {
		maze[i] = make([]int, columns)
		for j := range maze[i] {
			maze[i][j] = rand.Intn(2) // 0 for path, 1 for wall
		}
	}
	// Start and finish should always be paths
	maze[0][0], maze[rows-1][columns-1] = 0, 0
	return maze
}

func startGame() MazeState {
	rows, columns := 10, 10
	maze := generateMaze(rows, columns)
	explorer := Explorer{Coordinates{0, 0}}
	return MazeState{Explorer: explorer, Maze: maze, Status: StatusRunning}
}

func (ms *MazeState) moveExplorer(direction string) {
	moves := map[string]Coordinates{
		"UP":    {0, -1},
		"DOWN":  {0, 1},
		"LEFT":  {-1, 0},
		"RIGHT": {1, 0},
	}

	if move, ok := moves[direction]; ok {
		newX, newY := ms.Explorer.X+move.X, ms.Explorer.Y+move.Y
		if newX >= 0 && newY >= 0 && newX < len(ms.Maze[0]) && newY < len(ms.Maze) && ms.Maze[newY][newX] == 0 {
			ms.Explorer.X, ms.Explorer.Y = newX, newY
		}
	}

	if ms.Explorer.X == len(ms.Maze[0])-1 && ms.Explorer.Y == len(ms.Maze)-1 {
		ms.Status = StatusWon
	}
}

func main() {
	rate, _ := time.ParseDuration(os.Getenv("RATE"))
	rand.Seed(time.Now().UnixNano())

	game := startGame()

	for game.Status == StatusRunning {
		var direction string
		fmt.Println("Enter move (UP, DOWN, LEFT, RIGHT):")
		fmt.Scanln(&direction)

		game.moveExplorer(direction)

		fmt.Println("Current Maze (Pseudo-representation):")
		for _, row := range game.Maze {
			for _, cell := range row {
				fmt.Print(cell, " ")
			}
			fmt.Println()
		}

		if game.Status == StatusWon {
			fmt.Println("Congratulations, you've won!")
			return
		}

		time.Sleep(rate)
	}
}