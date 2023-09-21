package portaler

func (r *PortalsRenderer) drawWallAsTexturedOnScreenTrapezoid(l *linedef, wall, fitIn *trapezoid, wallType wallTypeCode) {
	if wall.x2 < wall.x1 {
		panic("Inverted trapezoid!")
	}
	if fitIn.x2 < fitIn.x1 {
		panic("Inverted trapezoid!")
	}
	tex := r.scene.Textures[0]
	switch wallType {
	case wallTypeUpper:
		tex = r.scene.Textures[1]
	case wallTypeLower:
		tex = r.scene.Textures[2]
	}
	tw, th := tex.W, tex.H
	wallLength := l.getLength()
	onScreenWidth := wall.x2 - wall.x1
	xTextureStep := (wallLength * float64(tw)) / float64(onScreenWidth)
	texXCoord := 0.0
	for x := wall.x1; x <= wall.x2; x++ {
		// horizontal clipping
		if x < fitIn.x1 {
			texXCoord += xTextureStep * float64(fitIn.x1-x)
			x = fitIn.x1
		}
		if x > fitIn.x2 {
			return
		}
		if x < 0 {
			texXCoord += xTextureStep * float64(-x)
			x = 0
		}
		if x >= r.screenW {
			return
		}
		if r.renderedColumnsBuffer.isColumnFull(x) {
			texXCoord += xTextureStep
			continue
		}
		currLower, currUpper := wall.getLowerAndUpperYCoordAtX(x)
		onScreenHeight := currLower - currUpper
		texYStep := float64(th) / float64(onScreenHeight)
		texYCoord := -texYStep
		// vertical clipping
		clipLower, clipUpper := r.renderedColumnsBuffer.getLowerAndUpperAt(x)
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
			texYCoord += texYStep * float64(clipUpper-currUpper)
			currUpper = clipUpper
		}

		if wallType != wallTypeUpper {
			r.renderedColumnsBuffer.setNewLowerAt(currUpper, x)
		}
		if wallType != wallTypeLower {
			r.renderedColumnsBuffer.setNewUpperAt(currLower, x)
		}

		// the drawing itself:
		for y := currUpper; y < currLower; y++ {
			r.debugFlush()
			rr, gg, bb, _ := tex.Bitmap.At(int(texXCoord)%tw, int(texYCoord)%th).RGBA()
			r.io.SetColor(uint8(rr), uint8(gg), uint8(bb))
			r.io.DrawPoint(int32(x), int32(y))
			texYCoord += texYStep
		}
		texXCoord += xTextureStep
	}
}
