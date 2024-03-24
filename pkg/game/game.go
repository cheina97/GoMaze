package game

import (
	"os"

	"github.com/cheina97/gomaze/pkg/grid"
	"github.com/cheina97/gomaze/pkg/matrix"
	"github.com/dominikbraun/graph/draw"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type lastPosSet string

const (
	start lastPosSet = "start"
	end   lastPosSet = "end"
)

type Game struct {
	w, h         int
	Matrix       *matrix.Matrix
	nextPosToSet lastPosSet
	startFinding bool
}

func NewGame(w int, h int) (*Game, error) {
	ebiten.SetWindowSize(w*grid.CellSize, h*grid.CellSize)
	ebiten.SetWindowTitle("GoMaze")
	m := matrix.NewMatrix(h, w)
	if err := m.GenerateRandomWalls(); err != nil {
		return nil, err
	}
	if err := m.InitEdges(); err != nil {
		return nil, err
	}
	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(m.Graph, file)

	return &Game{
		w:            w,
		h:            h,
		Matrix:       m,
		nextPosToSet: start,
		startFinding: false,
	}, nil
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		i, j := grid.GetGridCellFromCoords(mx, my)

		p := g.Matrix.Get(j, i)
		if p == matrix.Invalid || p == matrix.Wall {
			return nil
		}

		switch g.nextPosToSet {
		case end:
			g.Matrix.CleanAllValues(matrix.End)
			g.Matrix.Set(j, i, matrix.End)
			g.Matrix.End = &matrix.Point{X: j, Y: i}
			g.nextPosToSet = start
			g.startFinding = true
		case start:
			g.Matrix.CleanAllValues(matrix.Start)
			g.Matrix.Set(j, i, matrix.Start)
			g.Matrix.Start = &matrix.Point{X: j, Y: i}
			g.nextPosToSet = end
			g.startFinding = true
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Matrix = matrix.NewMatrix(g.h, g.w)
		if err := g.Matrix.GenerateRandomWalls(); err != nil {
			return err
		}
		if err := g.Matrix.InitEdges(); err != nil {
			return err
		}
	}

	if g.Matrix.Start != nil && g.Matrix.End != nil && g.startFinding {
		g.Matrix.CleanAllValues(matrix.Path)
		if err := g.Matrix.FindMinimumPath(); err != nil {
			return err
		}
		g.startFinding = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	grid.DrawGrid(screen, g.Matrix)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
