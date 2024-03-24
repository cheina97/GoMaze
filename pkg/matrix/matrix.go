package matrix

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dominikbraun/graph"
)

type Point struct {
	X, Y int
}

type Matrix struct {
	Rows, Cols int
	Data       [][]int
	Graph      graph.Graph[string, string]
	Start, End *Point
}

// NewMatrix creates a new matrix with the given number of rows and columns.
func NewMatrix(rows, cols int) *Matrix {
	m := &Matrix{
		Rows: rows, Cols: cols,
		Graph: graph.New(graph.StringHash),
	}
	m.Data = make([][]int, rows)
	for i := range m.Data {
		m.Data[i] = make([]int, cols)
		for j := range m.Data[i] {
			m.Data[i][j] = Empty
		}
	}
	return m
}

// Set sets the value at the given row and column.
func (m *Matrix) Set(col, row, val int) {
	m.Data[row][col] = val
}

func (m *Matrix) CleanAllValues(v int) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if m.Data[i][j] == v {
				m.Set(j, i, Empty)
			}
		}
	}
}

func (m *Matrix) Get(col, row int) int {
	if col < 0 || col >= m.Cols || row < 0 || row >= m.Rows {
		return Invalid
	}
	return m.Data[row][col]
}

func (m *Matrix) GenerateRandomWalls() error {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if rand.Int()%3 == 0 {
				m.Set(j, i, Wall)
			} else {
				if err := m.Graph.AddVertex(fmt.Sprintf("%d/%d", i, j)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *Matrix) InitEdges() error {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			if m.Get(j, i) == Wall {
				continue
			}
			if m.Get(j, i-1) == Empty {
				if err := m.Graph.AddEdge(fmt.Sprintf("%d/%d", i, j), fmt.Sprintf("%d/%d", i-1, j)); err != nil && err != graph.ErrEdgeAlreadyExists {
					return err
				}
			}
			if m.Get(j, i+1) == Empty {
				if err := m.Graph.AddEdge(fmt.Sprintf("%d/%d", i, j), fmt.Sprintf("%d/%d", i+1, j)); err != nil && err != graph.ErrEdgeAlreadyExists {
					return err
				}
			}
			if m.Get(j-1, i) == Empty {
				if err := m.Graph.AddEdge(fmt.Sprintf("%d/%d", i, j), fmt.Sprintf("%d/%d", i, j-1)); err != nil && err != graph.ErrEdgeAlreadyExists {
					return err
				}
			}
			if m.Get(j+1, i) == Empty {
				if err := m.Graph.AddEdge(fmt.Sprintf("%d/%d", i, j), fmt.Sprintf("%d/%d", i, j+1)); err != nil && err != graph.ErrEdgeAlreadyExists {
					return err
				}
			}
		}
	}
	return nil
}

func getCoordsFromString(s string) (int, int) {
	var i, j int
	fmt.Sscanf(s, "%d/%d", &i, &j)
	return i, j
}

func (m *Matrix) FindMinimumPath() error {
	path, err := graph.ShortestPath(m.Graph, fmt.Sprintf("%d/%d", m.Start.Y, m.Start.X), fmt.Sprintf("%d/%d", m.End.Y, m.End.X))
	if err != nil && err != graph.ErrTargetNotReachable {
		return err
	}
	for k := range path {
		if k == 0 || k == len(path)-1 {
			continue
		}
		i, j := getCoordsFromString(path[k])
		m.Set(j, i, Path)
	}
	return nil
}
