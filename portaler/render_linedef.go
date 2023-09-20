package portaler

func (r *PortalsRenderer) renderLinedef(l *linedef, s *sector, c *camera, screenArea *trapezoid) {
	transformed, wallTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformed {
		return
	}
	r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
	r.drawWallAsOnScreenTrapezoid(wallTrapezoid, screenArea, false)
	r.setColorWithBrightness(100, 32, 0, s.lightLevel)
	r.drawFloorUnderOnscreenTrapezoid(wallTrapezoid, screenArea)
	r.setColorWithBrightness(64, 32, 64, s.lightLevel)
	r.drawCeilingOverOnscreenTrapezoid(wallTrapezoid, screenArea)
}

func (r *PortalsRenderer) renderPortalLinedef(l *linedef, s *sector, c *camera, screenArea *trapezoid) {
	transformed, portalTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformed {
		return
	}

	fit := portalTrapezoid.fitInto(screenArea)
	if !fit {
		return
	}

	// needed for the "outsticking sector" bug workaround (see below)
	initialPortalTrapezoid := *portalTrapezoid

	// Render upper/lower wall
	nextSectorFloor := l.getNextSectorFrom(s).floorHeight
	nextSectorCeiling := l.getNextSectorFrom(s).ceilingHeight
	// we should render lower wall if next sector's floor is higher than the current one
	if nextSectorFloor > s.floorHeight {
		_, wallTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, nextSectorFloor, c, screenArea)
		r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
		r.drawWallAsOnScreenTrapezoid(wallTrapezoid, screenArea, false)
		// resize current portal window (decrease lower y) to prevent next sector to be drawn over lower wall
		portalTrapezoid.y1low = wallTrapezoid.y1high
		portalTrapezoid.y2low = wallTrapezoid.y2high
	}
	// we should render upper wall if next sector's ceiling is lower than the current one
	if nextSectorCeiling < s.ceilingHeight {
		_, wallTrapezoid := r.transformLinedefToScreenArea(l, nextSectorCeiling, s.ceilingHeight, c, screenArea)
		r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
		r.drawWallAsOnScreenTrapezoid(wallTrapezoid, screenArea, false)
		// resize current portal window (increase upper y) to prevent next sector to be drawn over upper wall
		portalTrapezoid.y1high = wallTrapezoid.y1low
		portalTrapezoid.y2high = wallTrapezoid.y2low
	}

	r.io.SetColor(l.r, l.g, l.b)
	r.drawWallAsOnScreenTrapezoid(portalTrapezoid, screenArea, true)

	r.renderSector(l.getNextSectorFrom(s), c, portalTrapezoid)

	r.setColorWithBrightness(100, 32, 0, s.lightLevel)
	r.drawFloorUnderOnscreenTrapezoid(&initialPortalTrapezoid, screenArea)
	r.setColorWithBrightness(64, 32, 64, s.lightLevel)
	r.drawCeilingOverOnscreenTrapezoid(&initialPortalTrapezoid, screenArea)

	// BUG: The next sectors may stick out if they're rendered AFTER floors/ceilings of the current one.
	// It is related only to sectors rendered through portals through third portals, and direct portal shall not cause it.
	// Most probable cause: allowed screen area for that sector is not quadrilateral.
	// I now draw sector BEFORE the floor/ceiling, but it causes pixel overdraw, which can be severely costly with texturization.
	// Outcommented, moved to bottom of this func until further inverstigation.

	//r.renderSector(l.getNextSectorFrom(s), c, portalTrapezoid)
}
