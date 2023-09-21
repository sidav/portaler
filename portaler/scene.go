package portaler

import (
	"math"
)

type Scene struct {
	sectors  []*sector
	Textures []Texture
}

func (s *Scene) addSector(sec *sector) {
	sec.id = len(s.sectors)
	s.sectors = append(s.sectors, sec)
}

func (s *Scene) scale(factor float64) {
	for _, sec := range s.sectors {
		for _, l := range sec.lines {
			fac := factor
			if l.isPortal {
				fac = math.Sqrt(factor)
			}
			l.x1 *= fac
			l.y1 *= fac
			l.x2 *= fac
			l.y2 *= fac
		}
	}
}

func (s *Scene) FindSectorWithCoordinates(x, y float64) *sector {
	for _, sec := range s.sectors {
		if sec.areCoordsInSector(x, y) {
			return sec
		}
	}
	return nil
}

func (s *Scene) portalizeLinedefWithCoordinates(fx, fy, tx, ty float64) {
	var s1ind, s2ind int
	for _, s1 := range s.sectors {
		s1ind = s1.getIndexOfLinedefByCoords(fx, fy, tx, ty)
		if s1ind == -1 {
			continue
		}
		for _, s2 := range s.sectors {
			if s2 == s1 {
				continue
			}
			s2ind = s2.getIndexOfLinedefByCoords(fx, fy, tx, ty)
			if s2ind == -1 {
				continue
			} else {
				portal := createPortalLinedef(fx, fy, tx, ty, s1, s2)
				portal.setColor(128, 0, 128)
				s1.lines[s1ind] = portal
				s2.lines[s2ind] = portal
				return
			}
		}
	}
	panic("Can't find belonging sector!")
}
