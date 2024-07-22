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
    Diameter float64
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

// checkCollisionWithWall returns true if the player collides with a given wall
func checkCollisionWithWall(player *Player, wall *Wall) bool {
    return (player.Position.X+player.Diameter/2 >= wall.Start.X && player.Position.X-player.Diameter/2 <= wall.End.X) &&
        (player.Position.Y+player.Diameter/2 >= wall.Start.Y && player.Position.Y-player.Diameter/2 <= wall.End.Y)
}

// checkWithinMazeBoundaries returns true if the player is within the maze boundaries
func checkWithinMazeBoundaries(player *Player) bool {
    return player.Position.X >= 0 && player.Position.X <= mazeWidth &&
        player.Position.Y >= 0 && player.Position.Y <= mazeHeight
}

// evaluateMazeConditions returns whether the player collides with any wall and is within the maze boundaries
func evaluateMazeConditions(player *Player, walls []Wall) (hasCollision bool, isInBounds bool) {
    for _, wall := range walls {
        if checkCollisionWithWall(player, &wall) {
            return true, checkWithinMazeBoundaries(player)
        }
    }
    return false, checkWithinMazeBoundaries(player)
}

func main() {
    player := Player{Position: Point{X: 10, Y: 10}, Diameter: 2}
    mazeWalls := []Wall{{Start: Point{X: 0, Y: 0}, End: Point{X: 10, Y: 1}}}
    collisionOccurred, withinBounds := evaluateMazeConditions(&player, mazeWalls)
    // Utilize collisionOccurred and withinBounds as needed
}