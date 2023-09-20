package portaler

func (r *PortalsRenderer) areCoordsOnMinimap(x, y int) bool {
	return areCoordsInRect(x, y, 0, 0, r.minimapW, r.minimapH)
}

func (r *PortalsRenderer) drawMinimap(s *Scene, c *camera) {
	r.io.SetColor(0, 0, 0)
	r.io.FillRect(0, 0, r.minimapW, r.minimapH)
	r.io.SetColor(255, 255, 255)
	r.io.DrawRect(0, 0, r.minimapW, r.minimapH)
	r.io.DrawRect(r.minimapW/2-1, r.minimapH/2-1, 3, 3)
	for _, sec := range s.sectors {
		for lineNum, l := range sec.lines {
			r.io.SetColor(l.r, l.g, l.b)
			r.drawLinedefOnMinimap(l, c)
			lineNum++
		}
	}
}

func (r *PortalsRenderer) drawLinedefOnMinimap(l *linedef, c *camera) {
	x1, y1, x2, y2 := c.transformLinedefToCameraSpace(l)
	sx1 := y1*r.minimapScale + float64(r.minimapW)/2
	sy1 := -x1*r.minimapScale + float64(r.minimapH)/2
	sx2 := y2*r.minimapScale + float64(r.minimapW)/2
	sy2 := -x2*r.minimapScale + float64(r.minimapH)/2
	if r.areCoordsOnMinimap(int(sx1), int(sy1)) && r.areCoordsOnMinimap(int(sx2), int(sy2)) {
		r.io.DrawLine(int(sx1), int(sy1), int(sx2), int(sy2))
	}
}
