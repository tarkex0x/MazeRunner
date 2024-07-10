package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	X int
	Y int
}

func (p *Player) MoveUp() {
	if p.Y > 0 {
		p.Y -= 1
	}
}

func (p *Player) MoveDown() {
	maxY, err := strconv.Atoi(os.Getenv("MAZE_HEIGHT"))
	if err != nil {
		fmt.Println("Error: Invalid MAZE_HEIGHT environment variable value")
		return
	}
	if p.Y < maxY-1 {
		p.Y += 1
	}
}

func (p *Player) MoveLeft() {
	if p.X > 0 {
		p.X -= 1
	}
}

func (p *Player) MoveRight() {
	maxX, err := strconv.Atoi(os.Getenv("MAZE_WIDTH"))
	if err != nil {
		fmt.Println("Error: Invalid MAZE_WIDTH environment variable value")
		return
	}
	if p.X < maxX-1 {
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
	player := Player{X: 0, Y: 0}

	player.UpdatePosition("right")
	player.UpdatePosition("down")
	player.DisplayPosition()
}