package portaler

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func areCoordsInRect(x, y, rx, ry, rw, rh int) bool {
	return x >= rx && x < rx+rw && y >= ry && y < ry+rh
}

func getLineIntersection(l1x1, l1y1, l1x2, l1y2, l2x1, l2y1, l2x2, l2y2 float64) (bool, float64, float64) {
	s1x := l1x2 - l1x1
	s1y := l1y2 - l1y1
	s2x := l2x2 - l2x1
	s2y := l2y2 - l2y1

	s := (-s1y*(l1x1-l2x1) + s1x*(l1y1-l2y1)) / (-s2x*s1y + s1x*s2y)
	t := (s2x*(l1y1-l2y1) - s2y*(l1x1-l2x1)) / (-s2x*s1y + s1x*s2y)

	// collision detected
	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true, l1x1 + (t * s1x), l1y1 + (t * s1y)
	}

	return false, 0, 0 // No collision
}
