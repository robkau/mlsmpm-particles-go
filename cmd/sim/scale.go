package main

func gridToWorld(x, y, gridToWorldScale float64) (xW, yW float64) {
	xW = x * gridToWorldScale
	yW = y * gridToWorldScale
	return
}

func worldToGrid(x, y, gridToWorldScale float64) (xG, yG float64) {
	xG = x / gridToWorldScale
	yG = y / gridToWorldScale
	return
}

func inverted(gridHeight, pos float64) float64 {
	return gridHeight - pos
}
