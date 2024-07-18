package main

import (
	"os"
	"strconv"
)

type Point struct {
	X, Y float64
}

type Wall struct {
	Start, End Point
}

type Player struct {
	Position Point
	Size     float64
}

var mazeWidth, mazeHeight float64

func init() {
	var err error
	mazeWidth, err = strconv.ParseFloat(os.Getenv("MAZE_WIDTH"), 64)
	if err != nil {
		panic("Failed to load MAZE_WIDTH from environment variables")
	}
	mazeHeight, err = strconv.ParseFloat(os.Getenv("MAZE_HEIGHT"), 64)
	if err != nil {
		panic("Failed to load MAZE_HEIGHT from environment variables")
	}
}

func checkCollisionWithWall(player *Player, wall *Wall) bool {
	return (player.Position.X+player.Size >= wall.Start.X && player.Position.X-player.Size <= wall.End.X) &&
		(player.Position.Y+player.Size >= wall.Start.Y && player.Position.Y-player.Size <= wall.End.Y)
}

func isPlayerWithinBounds(player *Player) bool {
	return player.Position.X >= 0 && player.Position.X <= mazeWidth &&
		player.Position.Y >= 0 && player.Position.Y <= mazeHeight
}

func checkCollisions(player *Player, walls []Wall) (bool, bool) {
	for _, wall := range walls {
		if checkCollisionWithWall(player, &wall) {
			return true, isPlayerWithinBounds(player)
		}
	}
	return false, isPlayerWithinMoveBounds(player)
}

func main() {
	player := Player{Position: Point{X: 10, Y: 10}, Size: 1}
	walls := []Wall{{Start: Point{X: 0, Y: 0}, End: Point{X: 10, Y: 1}}}
	collision, withinBounds := checkCollisions(&player, walls)
}