package main

import (
	"os"
	"strconv"
)

type Point struct {
	X, Y float64
}

type Wall struct {
	StartPoint, EndPoint Point
}

type Player struct {
	CurrentPosition Point
	Diameter        float64
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

func hasCollisionWithWall(player *Player, wall *Wall) bool {
	return (player.CurrentPosition.X+player.Diameter/2 >= wall.StartPoint.X && player.CurrentPosition.X-player.Diameter/2 <= wall.EndPoint.X) &&
		(player.CurrentPosition.Y+player.Diameter/2 >= wall.StartPoint.Y && player.CurrentPosition.Y-player.Diameter/2 <= wall.EndPoint.Y)
}

func isPlayerInsideMazeBoundaries(player *Player) bool {
	return player.CurrentPosition.X >= 0 && player.CurrentPosition.X <= mazeWidth &&
		player.CurrentPosition.Y >= 0 && player.CurrentPosition.Y <= mazeHeight
}

func assessCollisionsAndBounds(player *Player, mazeWalls []Wall) (hasCollision bool, isInBounds bool) {
	for _, wall := range mazeWalls {
		if hasCollisionWithWall(player, &wall) {
			return true, isPlayerInsideMazeBoundaries(player)
		}
	}
	return false, isTrainingWithInsideMazeBoundaries(player)
}

func main() {
	player := Player{CurrentPosition: Point{X: 10, Y: 10}, Diameter: 2}
	mazeWalls := []Wall{{StartPoint: Point{X: 0, Y: 0}, EndPoint: Point{X: 10, Y: 1}}}
	isColliding, withinMazeBounds := assessCollisionsAndBounds(&player, mazeWalls)
}