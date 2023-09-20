package portaler

type wallTypeCode uint8

const (
	wallTypeLower wallTypeCode = iota
	wallTypeUpper
	wallTypeFull
)

func (r *PortalsRenderer) drawWallAsOnScreenTrapezoid(wall, fitIn *trapezoid, asOutline bool, wallType wallTypeCode) {
	if wall.x2 < wall.x1 {
		panic("Inverted trapezoid!")
	}
	if fitIn.x2 < fitIn.x1 {
		panic("Inverted trapezoid!")
	}

	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		if x < 0 {
			x = 0
		}
		if x >= r.screenW {
			return
		}
		if r.filledColumnsTable[x] {
			continue
		}
		if wallType == wallTypeFull {
			r.filledColumnsTable[x] = true
		}
		currLower, currUpper := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		clipLower, clipUpper := r.renderedColumnsTable[x][0], r.renderedColumnsTable[x][1] // fitIn.getLowerAndUpperYCoordAtX(x)
		if currLower > clipLower && currUpper > clipLower {
			continue
		}
		if currLower < clipUpper && currUpper < clipUpper {
			continue
		}
		if currLower > clipLower {
			currLower = clipLower
		}
		if currUpper < clipUpper {
			currUpper = clipUpper
		}

		if wallType != wallTypeUpper {
			r.renderedColumnsTable[x][0] = currUpper
		}
		if wallType != wallTypeLower {
			r.renderedColumnsTable[x][1] = currLower
		}

		// the drawing itself:
		if asOutline {
			if x == wall.x1 || x == wall.x2 {
				r.io.VerticalLine(x, currLower, currUpper)
			}
			r.io.DrawPoint(int32(x), int32(currLower))
			r.io.DrawPoint(int32(x), int32(currUpper))
		} else {
			r.debugFlush()
			r.io.VerticalLine(x, currUpper, currLower)
		}
	}
	// debugPrintOnScreen(fitIn.x1, (fitIn.y1high+fitIn.y1low)/2, "%+v", fitIn.getInfoString())
	// debugPrintOnScreen(fitIn.x1, (fitIn.y1high+fitIn.y1low)/2+16, "%+v", wall.getInfoString())
}

func (r *PortalsRenderer) drawFloorUnderOnscreenTrapezoid(wall, fitIn *trapezoid) {
	if wall.x2 < wall.x1 {
		panic("Inverted trapezoid!")
	}

	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		if x < 0 {
			x = 0
		}
		if x >= r.screenW {
			return
		}
		topY, _ := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		if topY < r.renderedColumnsTable[x][1] {
			topY = r.renderedColumnsTable[x][1]
		}
		bottomY := r.renderedColumnsTable[x][0]
		if topY > bottomY {
			continue
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
		r.debugFlush()
		r.renderedColumnsTable[x][0] = topY
	}
}

func (r *PortalsRenderer) drawCeilingOverOnscreenTrapezoid(wall, fitIn *trapezoid) {
	if wall.x2 < wall.x1 {
		panic("Inverted trapezoid!")
	}
	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		if x < 0 {
			x = 0
		}
		if x >= r.screenW {
			return
		}
		_, bottomY := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		if bottomY > r.renderedColumnsTable[x][0] {
			bottomY = r.renderedColumnsTable[x][0]
		}
		topY := r.renderedColumnsTable[x][1]
		if topY > bottomY {
			continue
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
		r.renderedColumnsTable[x][1] = bottomY
	}
}
