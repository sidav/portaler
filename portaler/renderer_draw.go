package portaler

func (r *PortalsRenderer) drawWallAsOnScreenTrapezoid(wall, fitIn *trapezoid, asOutline bool) {
	if wall.x2 < wall.x1 {
		panic("Inverted trapezoid!")
	}
	if fitIn.x2 < fitIn.x1 {
		panic("Inverted trapezoid!")
	}
	// fmt.Printf("Drawing: %s", wall.getInfoString())

	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		currLower, currUpper := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		clipLower, clipUpper := fitIn.getLowerAndUpperYCoordAtX(x)
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

		// the drawing itself:
		if asOutline {
			if x == wall.x1 || x == wall.x2 {
				r.io.VerticalLine(x, currLower, currUpper)
			}
			r.io.DrawPoint(int32(x), int32(currLower))
			r.io.DrawPoint(int32(x), int32(currUpper))
		} else {
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
	// fmt.Printf("Drawing: %s", wall.getInfoString())

	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		topY, _ := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		bottomY, highestAllowedY := fitIn.getLowerAndUpperYCoordAtX(x)
		if topY > bottomY {
			continue
		}
		if topY < highestAllowedY {
			bottomY = highestAllowedY
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
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
		_, bottomY := wall.getLowerAndUpperYCoordAtX(x)
		// vertical clipping
		lowestAllowedY, topY := fitIn.getLowerAndUpperYCoordAtX(x)
		if bottomY < topY {
			continue
		}
		if bottomY > lowestAllowedY {
			bottomY = lowestAllowedY
		}
		// the drawing itself:
		r.io.VerticalLine(x, topY, bottomY)
	}
}
