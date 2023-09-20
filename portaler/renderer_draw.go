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
		if r.renderedColumnsBuffer.isColumnFull(x) {
			continue
		}
		currLower, currUpper := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		clipLower, clipUpper := r.renderedColumnsBuffer.getLowerAndUpperAt(x) // fitIn.getLowerAndUpperYCoordAtX(x)
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
			r.renderedColumnsBuffer.setNewLowerAt(currUpper, x)
		}
		if wallType != wallTypeLower {
			r.renderedColumnsBuffer.setNewUpperAt(currLower, x)
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
		if topY < r.renderedColumnsBuffer.getUpperAt(x) {
			topY = r.renderedColumnsBuffer.getUpperAt(x)
		}
		bottomY := r.renderedColumnsBuffer.getLowerAt(x)
		if topY > bottomY {
			continue
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
		r.debugFlush()
		r.renderedColumnsBuffer.setNewLowerAt(topY, x)
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
		if bottomY > r.renderedColumnsBuffer.getLowerAt(x) {
			bottomY = r.renderedColumnsBuffer.getLowerAt(x)
		}
		topY := r.renderedColumnsBuffer.getUpperAt(x)
		if topY > bottomY {
			continue
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
		r.renderedColumnsBuffer.setNewUpperAt(bottomY, x)
	}
}
