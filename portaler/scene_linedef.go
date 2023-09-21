package portaler

import "math"

type linedef struct {
	x1, y1, x2, y2  float64
	isPortal        bool
	connectsSectors [2]*sector

	r, g, b uint8
}

func createLinedef(fx, fy, tx, ty float64) *linedef {
	return &linedef{x1: fx, y1: fy, x2: tx, y2: ty, isPortal: false}
}

func (l *linedef) setColor(r, g, b uint8) {
	l.r, l.g, l.b = r, g, b
}

func (l *linedef) getLength() float64 {
	normX, normY := l.x2-l.x1, l.y2-l.y1
	return math.Sqrt(normX*normX + normY*normY)
}

func createPortalLinedef(fx, fy, tx, ty float64, s1, s2 *sector) *linedef {
	return &linedef{
		x1: fx, y1: fy, x2: tx, y2: ty,
		isPortal:        true,
		connectsSectors: [2]*sector{s1, s2},
	}
}

func (l *linedef) getNextSectorFrom(s *sector) *sector {
	if !l.isPortal {
		panic("Not a portal!")
	}
	if l.connectsSectors[0] == s {
		return l.connectsSectors[1]
	} else if l.connectsSectors[1] == s {
		return l.connectsSectors[0]
	} else {
		panic("Something is wrong with portals connection data!")
	}
}
