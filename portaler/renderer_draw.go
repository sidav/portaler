package portaler

func (r *PortalsRenderer) drawWallAsOnScreenTrapezoid(wall, fitIn *trapezoid, asOutline bool) {
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
				r.verticalLine(x, currLower, currUpper)
			}
			r.drawPoint(x, currLower)
			r.drawPoint(x, currUpper)
		} else {
			r.verticalLine(x, currUpper, currLower)
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
		r.verticalLine(x, topY, bottomY)
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
		r.verticalLine(x, topY, bottomY)
	}
}

func (r *PortalsRenderer) drawPoint(x, y int) {
	// if x < 0 || y < 0 || x >= r.screenW || y >= r.screenH {
	// 	return
	// }
	// if r.renderedPixelsBuffer[x][y] == r.totalFramesRendered {
	// 	return
	// }
	// r.renderedPixelsBuffer[x][y] = r.totalFramesRendered
	r.io.DrawPoint(int32(x), int32(y))
}

func (r *PortalsRenderer) verticalLine(x, y1, y2 int) {
	// if y1 > y2 {
	// 	t := y1
	// 	y1 = y2
	// 	y2 = t
	// }
	// for y := y1; y < y2; y++ {
	// 	r.drawPoint(x, y)
	// }
	r.io.VerticalLine(x, y1, y2)
}
