package portaler

func (r *PortalsRenderer) getLowerAndUpperScreenYForTransformedVertex(x, y, lowerHeight, upperHeight float64, c *camera) (int, int) {
	// 0.5 here means "half of the screen"
	// It should be 0.5 PLUS, but screen Y coordinates go down, so we need to invert
	lower := 0.5 - r.aspectRatio*c.distToScreenPlane*(lowerHeight-c.Height)/x
	upper := 0.5 - r.aspectRatio*c.distToScreenPlane*(upperHeight-c.Height)/x
	return int(float64(r.screenH) * lower), int(float64(r.screenH) * upper)
}

func (r *PortalsRenderer) transformPointToScreenXCoord(x, y float64, c *camera) float64 {
	// angle := math.Atan2(y, x)
	// 0.5 here means "half of the screen"
	res := 0.5 + c.distToScreenPlane*y/x // *math.Tan(angle)
	return res
}

func (r *PortalsRenderer) transformPortalToScreenArea(l *linedef, floorH, ceilingH float64, c *camera, fitIn *trapezoid) (bool, *trapezoid) {
	x1, _, x2, _ := c.transformLinedefToCameraSpace(l)
	// culling:
	if x1 < c.distToScreenPlane || x2 < c.distToScreenPlane {
		if x1 > 0 || x2 > 0 {
			// yes, it's on purpose: allow the portal to be drawn full-screen and leave the clipping and stuff to the rendered columns buffer
			return true, newTrapezoid(0, r.screenH, 0, r.screenW, r.screenH, 0)
		}
		return false, nil
	}
	// TODO: optimize this call (separate projection func?)
	return r.transformLinedefToScreenArea(l, floorH, ceilingH, c, fitIn)
}

func (r *PortalsRenderer) transformLinedefToScreenArea(l *linedef, floorH, ceilingH float64, c *camera, fitIn *trapezoid) (bool, *trapezoid) {
	x1, y1, x2, y2 := c.transformLinedefToCameraSpace(l)
	// culling:
	if x1 < c.distToScreenPlane && x2 < c.distToScreenPlane {
		return false, nil
	}
	intersect, ix, iy := getLineIntersection(x1, y1, x2, y2, c.distToScreenPlane/8, 500, c.distToScreenPlane/8, -500)
	if intersect {
		if x1 < c.distToScreenPlane {
			x1, y1 = ix, iy
		} else if x2 < c.distToScreenPlane {
			x2, y2 = ix, iy
		}
	}
	// clipping ended

	screenX1 := int(float64(r.screenW) * r.transformPointToScreenXCoord(x1, y1, c))
	screenX2 := int(float64(r.screenW) * r.transformPointToScreenXCoord(x2, y2, c))

	// invert trapezoid so that it goes left -> right
	if screenX1 > screenX2 {
		// debugPrintf("Reverting linedef\n")
		t := screenX1
		screenX1 = screenX2
		screenX2 = t

		tf := y1
		y1 = y2
		y2 = tf
		tf = x1
		x1 = x2
		x2 = tf
	}
	// Cull against the given screenArea (horizontal only)
	if screenX1 == screenX2 {
		return false, nil
	}
	if screenX1 < fitIn.x1 && screenX2 <= fitIn.x1 {
		return false, nil
	} else if screenX1 >= fitIn.x2 && screenX2 > fitIn.x2 {
		return false, nil
	}

	ly1int, uy1int := r.getLowerAndUpperScreenYForTransformedVertex(x1, y1, floorH, ceilingH, c)
	ly2int, uy2int := r.getLowerAndUpperScreenYForTransformedVertex(x2, y2, floorH, ceilingH, c)

	return true, newTrapezoid(
		screenX1,
		ly1int,
		uy1int,
		screenX2,
		ly2int,
		uy2int,
	)
}
