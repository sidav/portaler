package portaler

// Struct for clipping against already drawn screen parts and to eliminate overdraw.
// Holds uppest and lowest non-yet-drawn screen Y coords for each screen X coord.
// It is allowed to draw between upper and lower Y values only (non-inclusive).
// Screen coords go from top to bottom, so upper are generally smaller values than lower.
type columnsBuffer struct {
	upperCoords []int // can only increase before clear
	lowerCoords []int // can only decrease
}

func (cb *columnsBuffer) reinitForWidth(width int) {
	if len(cb.lowerCoords) < width {
		cb.lowerCoords = make([]int, width)
	}
	if len(cb.upperCoords) < width {
		cb.upperCoords = make([]int, width)
	}
}

func (cb *columnsBuffer) clear(lowestScreenCoord int) {
	for i := range cb.lowerCoords {
		cb.lowerCoords[i] = lowestScreenCoord
		cb.upperCoords[i] = 0
	}
}

func (cb *columnsBuffer) isColumnFull(x int) bool {
	return cb.lowerCoords[x] <= cb.upperCoords[x]
}

func (cb *columnsBuffer) getLowerAndUpperAt(x int) (int, int) {
	return cb.lowerCoords[x], cb.upperCoords[x]
}

func (cb *columnsBuffer) getUpperAt(x int) int {
	return cb.upperCoords[x]
}

func (cb *columnsBuffer) getLowerAt(x int) int {
	return cb.lowerCoords[x]
}

func (cb *columnsBuffer) setNewUpperAt(value, x int) {
	if cb.upperCoords[x] > value {
		panic("Clipping error: something decreases upper value!")
	}
	cb.upperCoords[x] = value
}

func (cb *columnsBuffer) setNewLowerAt(value, x int) {
	if cb.lowerCoords[x] < value {
		panic("Clipping error: something increases lower value!")
	}
	cb.lowerCoords[x] = value
}
