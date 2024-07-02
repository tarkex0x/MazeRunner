package main

import (
	"fmt"
	"math/rand"
	"os"
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
	stack = append(stack, *currentCell)

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
