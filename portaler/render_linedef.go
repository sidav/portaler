package portaler

func (r *PortalsRenderer) renderLinedef(l *linedef, s *sector, c *camera, screenArea *trapezoid) {
	transformed, wallTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformed {
		return
	}
	r.setColorWithBrightness(100, 32, 0, s.lightLevel)
	r.drawFloorUnderOnscreenTrapezoid(wallTrapezoid, screenArea)
	r.setColorWithBrightness(64, 32, 64, s.lightLevel)
	r.drawCeilingOverOnscreenTrapezoid(wallTrapezoid, screenArea)
	r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
	r.drawWallAsOnScreenTrapezoid(wallTrapezoid, screenArea, false, wallTypeFull)
}

func (r *PortalsRenderer) renderPortalLinedef(l *linedef, s *sector, c *camera, screenArea *trapezoid) {
	transformed, portalTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformed {
		return
	}

	// needed for the "outsticking sector" bug workaround (see below)
	initialPortalTrapezoid := *portalTrapezoid

	// Render upper/lower wall
	nextSectorFloor := l.getNextSectorFrom(s).floorHeight
	nextSectorCeiling := l.getNextSectorFrom(s).ceilingHeight
	var upperWallTrapezoid, lowerWallTrapezoid *trapezoid
	// we should render lower wall if next sector's floor is higher than the current one
	if nextSectorFloor > s.floorHeight {
		_, lowerWallTrapezoid = r.transformLinedefToScreenArea(l, s.floorHeight, nextSectorFloor, c, screenArea)
		// resize current portal window (decrease lower y) to prevent next sector to be drawn over lower wall
		portalTrapezoid.y1low = lowerWallTrapezoid.y1high
		portalTrapezoid.y2low = lowerWallTrapezoid.y2high
	}
	// we should render upper wall if next sector's ceiling is lower than the current one
	if nextSectorCeiling < s.ceilingHeight {
		_, upperWallTrapezoid = r.transformLinedefToScreenArea(l, nextSectorCeiling, s.ceilingHeight, c, screenArea)
		// resize current portal window (increase upper y) to prevent next sector to be drawn over upper wall
		portalTrapezoid.y1high = upperWallTrapezoid.y1low
		portalTrapezoid.y2high = upperWallTrapezoid.y2low
	}

	// r.io.SetColor(l.r, l.g, l.b)
	// r.drawWallAsOnScreenTrapezoid(portalTrapezoid, screenArea, true)

	r.setColorWithBrightness(100, 32, 0, s.lightLevel)
	r.drawFloorUnderOnscreenTrapezoid(&initialPortalTrapezoid, screenArea)
	r.setColorWithBrightness(64, 32, 64, s.lightLevel)
	r.drawCeilingOverOnscreenTrapezoid(&initialPortalTrapezoid, screenArea)
	r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
	if upperWallTrapezoid != nil {
		r.drawWallAsOnScreenTrapezoid(upperWallTrapezoid, screenArea, false, wallTypeUpper)
	}
	if lowerWallTrapezoid != nil {
		r.drawWallAsOnScreenTrapezoid(lowerWallTrapezoid, screenArea, false, wallTypeLower)
	}
	r.renderSector(l.getNextSectorFrom(s), c, portalTrapezoid)
}
