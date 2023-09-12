package portaler

type sector struct {
	id            int
	lines         []*linedef
	floorHeight   float64 // absolute
	ceilingHeight float64 // absolute
	lightLevel    uint8   // 0 - 255, 255 is the brightest
}

func (s *sector) addLinedef(fx, fy, tx, ty float64) {
	line := createLinedef(fx, fy, tx, ty)
	line.setColor(getDebugColor(len(s.lines)))
	s.lines = append(s.lines, line)
}

func (s *sector) appendLinedef(tx, ty float64) {
	prev := s.lines[len(s.lines)-1]
	s.addLinedef(prev.x2, prev.y2, tx, ty)
}

func (s *sector) getIndexOfLinedefByCoords(fx, fy, tx, ty float64) int {
	for i, l := range s.lines {
		if l.x1 == fx && l.y1 == fy && l.x2 == tx && l.y2 == ty {
			return i
		}
		if l.x1 == tx && l.y1 == ty && l.x2 == fx && l.y2 == fy {
			return i
		}
	}
	return -1
}

func (s *sector) areCoordsInSector(x, y float64) bool {
	intersections := 0
	for _, l := range s.lines {
		intersect, _, _ := getLineIntersection(l.x1, l.y1, l.x2, l.y2, x-10000, y, x, y)
		if intersect {
			intersections++
		}
	}
	return intersections%2 == 1
}
