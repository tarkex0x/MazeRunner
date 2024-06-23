package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type MazeGame struct {
    grid                     [][]string
    playerRow, playerCol     int
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
            {"#", " ", " ", " ", " ", " ", " ", "#", " ", "G"},
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
    oldRow, oldCol := game.playerRow, game.playerCol 
    switch direction {
    case "UP":
        game.playerRow = max(game.playerRow-1, 0)
    case "DOWN":
        game.playerRow = min(game.playerRow+1, len(game.grid)-1)
    case "LEFT":
        game.playerCol = max(game.playerCol-1, 0)
    case "RIGHT":
        game.playerCol = min(game.playerCol+1, len(game.grid[0])-1)
    }

    if game.grid[game.playerRow][game.playerCol] == "#" {
        game.playerRow, game.playerCol = oldRow, oldCol
        return
    }

    if game.grid[game.playerRow][game.playerCol] == "G" {
        fmt.Println("Congratulations! You've reached the goal.")
        os.Exit(0)
    }

    game.grid[oldRow][oldCol] = " "
    game.grid[game.playerRow][game.playerCol] = "P"
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
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

        if input == "UP" || input == "DOWN" || input == "LEFT" || input == "RIGHT" {
            mazeStayed.MovePlayer(input)
        } else {
            fmt.Println("Invalid input. Please enter UP, DOWN, LEFT, RIGHT, or EXIT.")
        }
    }
}