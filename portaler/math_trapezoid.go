package portaler

import (
	"fmt"
	"math"
)

// assumes left and right sides as parallel, and top and bottom as MAYBE not.
type trapezoid struct {
	x1, y1low, y1high int
	x2, y2low, y2high int

	// TODO: calculation of heights for "trimmed trapezoid"
	// like this one:
	//    ###########
	//   ############
	//  #############
	// ##############
	beforeTrimming *trapezoid
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
	if relativeX < 0 {
		// panic("Well, that's a crash")
	}
	if relativeX > width {
		// panic("Well, that's a crash too")
	}
	if width == 0 {
		return math.MinInt64, math.MaxInt64
	}
	low, up := t.y1low+(relativeX*(t.y2low-t.y1low))/width, t.y1high+(relativeX*(t.y2high-t.y1high))/width

	if t.beforeTrimming != nil {
		if t.beforeTrimming.beforeTrimming != nil {
			panic("Trimming is screwed up")
		}
		debugPrintf("\n YEAH \n")
		prevWidth := t.beforeTrimming.x2 - t.beforeTrimming.x1
		tlow, tup := t.beforeTrimming.getLowerAndUpperYCoordAtX(prevWidth * x / width)
		return min(low, tlow), max(up, tup)
	}
	return low, up
}

// returns true if t fits or intersects into/with frame
func (t *trapezoid) fitInto(frame *trapezoid) bool {
	if t.x1 < frame.x1 && t.x2 < frame.x1 {
		return false
	}
	if t.x1 > frame.x2 && t.x2 > frame.x2 {
		return false
	}
	if t.y1high > frame.y1low && t.y2high > frame.y2low {
		return false
	}
	if t.y1low < frame.y1high && t.y2low < frame.y2high {
		return false
	}
	untrimmed := newTrapezoid(t.x1, t.y1low, t.y1high, t.x2, t.y2low, t.y2high)
	if t.x1 < frame.x1 {
		tHl, tHu := t.getLowerAndUpperYCoordAtX(frame.x1)
		t.y1low = tHl
		t.y1high = tHu
		t.x1 = frame.x1
		t.beforeTrimming = untrimmed
	}
	if t.x2 > frame.x2 {
		tHl, tHu := t.getLowerAndUpperYCoordAtX(frame.x2)
		t.y2low = tHl
		t.y2high = tHu
		t.x2 = frame.x2
		t.beforeTrimming = untrimmed
	}
	t.beforeTrimming = nil
	return true
}
