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
	// r.drawWallAsOnScreenTrapezoid(wallTrapezoid, screenArea, false, wallTypeFull)
	r.drawWallAsTexturedOnScreenTrapezoid(l, wallTrapezoid, screenArea, wallTypeFull)
}

func (r *PortalsRenderer) renderPortalLinedef(l *linedef, s *sector, c *camera, screenArea *trapezoid) {
	transformedPortal, portalTrapezoid := r.transformPortalToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformedPortal {
		return
	}
	transformedWall, portalAsWallTrapezoid := r.transformLinedefToScreenArea(l, s.floorHeight, s.ceilingHeight, c, screenArea)
	if !transformedWall {
		r.renderSector(l.getNextSectorFrom(s), c, portalTrapezoid)
		return
	}

	// Render upper/lower wall
	nextSectorFloor := l.getNextSectorFrom(s).floorHeight
	nextSectorCeiling := l.getNextSectorFrom(s).ceilingHeight
	var upperWallTrapezoid, lowerWallTrapezoid *trapezoid
	// we should render lower wall if next sector's floor is higher than the current one
	if nextSectorFloor > s.floorHeight {
		_, lowerWallTrapezoid = r.transformLinedefToScreenArea(l, s.floorHeight, nextSectorFloor, c, screenArea)
	}
	// we should render upper wall if next sector's ceiling is lower than the current one
	if nextSectorCeiling < s.ceilingHeight {
		_, upperWallTrapezoid = r.transformLinedefToScreenArea(l, nextSectorCeiling, s.ceilingHeight, c, screenArea)
	}

	// r.io.SetColor(l.r, l.g, l.b)
	// r.drawWallAsOnScreenTrapezoid(portalTrapezoid, screenArea, true)

	r.setColorWithBrightness(100, 32, 0, s.lightLevel)
	r.drawFloorUnderOnscreenTrapezoid(portalAsWallTrapezoid, screenArea)
	r.setColorWithBrightness(64, 32, 64, s.lightLevel)
	r.drawCeilingOverOnscreenTrapezoid(portalAsWallTrapezoid, screenArea)
	r.setColorWithBrightness(l.r, l.g, l.b, s.lightLevel)
	if upperWallTrapezoid != nil {
		r.drawWallAsTexturedOnScreenTrapezoid(l, upperWallTrapezoid, screenArea, wallTypeUpper)
	}
	if lowerWallTrapezoid != nil {
		r.drawWallAsTexturedOnScreenTrapezoid(l, lowerWallTrapezoid, screenArea, wallTypeLower)
	}
	r.renderSector(l.getNextSectorFrom(s), c, portalTrapezoid)
}
