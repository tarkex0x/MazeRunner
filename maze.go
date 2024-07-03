package main

import (
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

type MazeCell struct {
    X, Y     int  
    Visited  bool 
    Walls    [4]bool 
}

type Maze struct {
    Cells [][]MazeCell
    Width, Height int
}

func (m *Maze) Initialize() {
    m.Cells = make([][]MazeCell, m.Height)
    for i := range m.Cells {
        m.Cells[i] = make([]MazeCell, m.Width)
        for j := range m.Cells[i] {
            m.Cells[i][j] = MazeCell{X: j, Y: i, Visited: false, Walls: [4]bool{true, true, true, true}}
        }
    }
}

func (m *Maze) GenerateMaze() {
    rand.Seed(time.Now().UnixNano())
    var stack []MazeCell
    currentCell := &m.Cells[0][0]
    currentCell.Visited = true
    stack = append(stack, *current332Cell)

    for len(stack) > 0 {
        nextCell := m.GetNextCell(*currentCell)

        if nextCell != nil {
            stack = append(stack, *currentCell)
            m.RemoveWalls(currentCell, nextCell)

            currentCell = nextCell
            currentCell.Visited = true
        } else if len(stack) > 0 {
            currentCell = &stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
    }
}

func (m *Maze) GetNextCell(cell MazeCell) *MazeCell {
    neighbors := []MazeCell{}

    directions := []struct{ dx, dy int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} 
    for _, dir := range directions {
        x, y := cell.X+dir.dx, cell.Y+dir.dy

        if x >= 0 && y >= 0 && x < m.Width && y < m.Height && !m.Cells[y][x].Visited {
            neighbors = append(neighbors, m.Cells[y][x])
        }
    }

    if len(neighbors) > 0 {
        return &neighbors[rand.Intn(len(neighbors))]
    }
    return nil
}

func (m *Maze) RemoveWalls(current, next *MazeCell) {
    dx := current.X - next.X
    dy := current.Y - next.Y

    if dx == 1 {
        current.Walls[3] = false
        next.Walls[1] = false
    } else if dx == -1 {
        current.Walls[1] = false
        next.Walls[3] = false
    }

    if dy == 1 {
        current.Walls[0] = false
        next.Walls[2] = false
    } else if dy == -1 {
        current.Walls[2] = false
        next.Walls[0] = false
    }
}

func (m *Maze) Render() {
    for i := 0; i < m.Height; i++ {
        // Top walls
        for j := 0; j < m.Width; j++ {
            fmt.Print("+")
            if m.Cells[i][j].Walls[0] {
                fmt.Print("---")
            } else {
                fmt.Print("   ")
            }
        }
        fmt.Println("+")
        // Side walls
        for j := 0; j < m.Width; j++ {
            if m.Cells[i][j].Walls[3] {
                fmt.Print("|")
            } else {
                fmt.Print(" ")
            }
            fmt.Print("   ")
        }
        fmt.Println("|")
    }
    // Print the bottom wall
    fmt.Println(strings.Repeat("+---", m.Width) + "+")
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: mazerunner width height")
        os.Exit(1)
    }
    
    var width, height int
    fmt.Sscanf(os.Args[1], "%d", &width)
    fmt.Sscanf(os.Args[2], "%d", &height)

    maze := Maze{Width: width, Height: height}
    maze.Initialize()
    maze.GenerateMaze()
    maze.Render()
}