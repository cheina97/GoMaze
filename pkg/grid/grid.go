package grid

import (
	"image/color"

	"github.com/cheina97/gomaze/pkg/matrix"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGrid(dst *ebiten.Image, m *matrix.Matrix) {
	for x := 0; x < m.Cols; x++ {
		for y := 0; y < m.Rows; y++ {
			switch m.Get(x, y) {
			case matrix.Empty:
				DrawGridCell(dst, x, y, emptyColor)
			case matrix.Wall:
				DrawGridCell(dst, x, y, wallColor)
			case matrix.Start:
				DrawGridCell(dst, x, y, startColor)
			case matrix.End:
				DrawGridCell(dst, x, y, endColor)
			case matrix.Path:
				DrawGridCell(dst, x, y, pathColor)
			}
		}
	}
}

func DrawGridCell(dst *ebiten.Image, x, y int, cl color.Color) {
	vector.DrawFilledRect(dst, float32(x*CellSize), float32(y*CellSize), float32(CellSize), float32(CellSize), cl, false)
}

func GetGridCellFromCoords(x, y int) (int, int) {
	return y / (CellSize), x / (CellSize)
}
