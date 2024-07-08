package main

import (
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
	maxY, _ := strconv.Atoi(os.Getenv("MAZE_HEIGHT"))
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
	maxX, _ := strconv.Atoi(os.Getenv("MAZE_WIDTH"))
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
		fmt.Println("Invalid direction. Please use 'up', 'down', 'left' or 'right'.")
	}
}

func (p *Player) DisplayPosition() {
	fmt.Printf("Player is at X: %d, Y: %d\n", p.X, p.Y)
}

func main() {
	player := Player{X: 0, Y: 0}

	player.UpdatePosition("right")
	player.UpdatePosition("down")
	player.DisplayPosition()
}