package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	GameRunning = iota
	GameOver
	GameWon
)

type Position struct {
	X, Y int
}

type Player struct {
	Position
}

type GameState struct {
	Player    Player
	Maze      [][]int
	GameState int
}

func initMaze(rows, cols int) [][]int {
	maze := make([][]int, rows)
	for i := range maze {
		maze[i] = make([]int, cols)
		for j := range maze[i] {
			maze[i][j] = rand.Intn(2)
		}
	}
	maze[0][0], maze[rows-1][cols-1] = 0, 0
	return maze
}

func initGame() GameState {
	rows, cols := 10, 10
	maze := initMaze(rows, cols)
	player := Player{Position{0, 0}}
	return GameState{player, maze, GameRunning}
}

func (game *GameState) updateGameState(move string) {
	directions := map[string]Position{
		"UP":    {0, -1},
		"DOWN":  {0, 1},
		"LEFT":  {-1, 0},
		"RIGHT": {1, 0},
	}

	if dir, ok := directions[move]; ok {
		newX, newY := game.Player.X+dir.X, game.Player.Y+dir.Y
		if newX >= 0 && newY >= 0 && newX < len(game.Maze[0]) && newY < len(game.Maze) && game.Maze[newY][newX] == 0 {
			game.Player.X, game.Player.Y = newX, newY
		}
	}

	if game.Player.X == len(game.Maze[0])-1 && game.Player.Y == len(game.Maze)-1 {
		game.GameState = GameWon
	}
}

func main() {
	rate, _ := time.ParseDuration(os.Getenv("RATE"))
	rand.Seed(time.Now().UnixNano())

	game := initGame()

	for game.GameState == GameRunning {
		var move string
		fmt.Println("Enter move (UP, DOWN, LEFT, RIGHT):")
		fmt.Scanln(&move)

		game.updateGameState(move)

		fmt.Println("Current Maze (Pseudo-representation):")
		for _, row := range game.Maze {
			for _, col := range row {
				fmt.Print(col, " ")
			}
			fmt.Println()
		}

		switch game.GameState {
		case GameWon:
			fmt.Println("Congratulations, you've won!")
			return
		}

		time.Sleep(rate)
	}
}