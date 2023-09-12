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
	c.Rotate(-3.14 / 2)
	for _, sec := range s.sectors {
		for lineNum, l := range sec.lines {
			r.io.SetColor(l.r, l.g, l.b)
			r.drawLinedefOnMinimap(l, c)
			lineNum++
		}
	}
	c.Rotate(3.14 / 2)
}

func (r *PortalsRenderer) drawLinedefOnMinimap(l *linedef, c *camera) {
	x1, y1, x2, y2 := c.transformLinedefToCameraSpace(l)
	x1 = x1*r.minimapScale + float64(r.minimapW)/2
	y1 = y1*r.minimapScale + float64(r.minimapH)/2
	x2 = x2*r.minimapScale + float64(r.minimapW)/2
	y2 = y2*r.minimapScale + float64(r.minimapH)/2
	if r.areCoordsOnMinimap(int(x1), int(y1)) && r.areCoordsOnMinimap(int(x2), int(y2)) {
		r.io.DrawLine(int(x1), int(y1), int(x2), int(y2))
	}
}
