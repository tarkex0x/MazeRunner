package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	X      int
	Y      int
	MaxX   int
	MaxY   int
}

func (p *Player) MoveUp() {
	if p.Y > 0 {
		p.Y -= 1
	}
}

func (p *Player) MoveDown() {
	if p.Y < p.MaxY-1 {
		p.Y += 1
	}
}

func (p *Player) MoveLeft() {
	if p.X > 0 {
		p.X -= 1
	}
}

func (p *Player) MoveRight() {
	if p.X < p.MaxX-1 {
		p.X += 1
	}
}

func (p *Player) UpdatePosition(direction string) {
	switch direction {
	case "up":
		p.MoveUp()
	case "down":
		p.MoveDown()
	case "left":
		p.MoveLeft()
	case "right":
		p.MoveRight()
	default:
		fmt.Println("Invalid direction. Please use 'up', 'down', 'left', or 'right'.")
	}
}

func (p *Player) DisplayPosition() {
	fmt.Printf("Player is at X: %d, Y: %d\n", p.X, p.Y)
}

func getEnvAsInt(key string) (int, error) {
	valStr := os.Getenv(key)
	if valStr == "" {
		return 0, errors.New("environment variable not found")
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("invalid value for environment variable %s: %v", key, err)
	}

	return val, nil
}

func main() {
	maxX, errX := getEnvAsInt("MAZE_WIDTH")
	maxY, errY := getEnvAsInt("MAZE_HEIGHT")
	if errX != nil || errY != nil {
		fmt.Println("Error reading MAZE dimensions: ", errX, errY)
		return
	}

	player := Player{X: 0, Y: 0, MaxX: maxX, MaxY: maxY}

	player.UpdatePosition("right")
	player.UpdatePosition("down")
	player.DisplayOutputPosition()
}