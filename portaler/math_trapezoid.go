package portaler

import (
	"fmt"
	"math"
)

// assumes left and right sides as parallel, and top and bottom as MAYBE not.
type trapezoid struct {
	x1, y1low, y1high int
	x2, y2low, y2high int
}

func newTrapezoid(x1, y1l, y1h, x2, y2l, y2h int) *trapezoid {
	return &trapezoid{
		x1:     x1,
		y1low:  y1l,
		y1high: y1h,
		x2:     x2,
		y2low:  y2l,
		y2high: y2h,
	}
}

func (t *trapezoid) getInfoString() string {
	return fmt.Sprintf("{x %d ly %d uy %d} -> {x %d ly %d uy %d}; ", t.x1, t.y1low, t.y1high, t.x2, t.y2low, t.y2high)
}

func (t *trapezoid) getLowerAndUpperYCoordAtX(x int) (int, int) {
	relativeX := x - t.x1
	width := t.x2 - t.x1
	if width == 0 {
		return math.MinInt64, math.MaxInt64
	}
	return t.y1low + (relativeX*(t.y2low-t.y1low))/width, t.y1high + (relativeX*(t.y2high-t.y1high))/width
}
