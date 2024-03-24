package grid

import "image/color"

const (
	CellSize     = 30
	BorderOffset = 20
)

var (
	emptyColor = color.RGBA{255, 255, 255, 0}
	wallColor  = color.RGBA{0, 0, 0, 0}
	startColor = color.RGBA{75, 162, 71, 0}
	endColor   = color.RGBA{255, 20, 87, 0}
	pathColor  = color.RGBA{255, 219, 87, 0}
)
