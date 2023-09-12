package portaler

import "portalrenderer/backend"

func NewRenderer(b backend.RendererBackend, screenW, screenH int, sc *Scene) *PortalsRenderer {
	r := &PortalsRenderer{
		io:      b.(*backend.RaylibBackend),
		screenW: screenW,
		screenH: screenH,
	}
	r.minimapH = r.screenH / 4
	r.minimapW = r.minimapH
	r.minimapScale = float64(r.minimapH) / 32
	r.setToScene(sc)
	return r
}

func (r *PortalsRenderer) setToScene(s *Scene) {
	r.scene = s
	r.renderedSectorsTable = make([]bool, len(s.sectors))
	// double-check scene's sector ids adequacy
	for i := range s.sectors {
		for j := i + 1; j < len(s.sectors); j++ {
			if s.sectors[i].id == s.sectors[j].id {
				panic("Sector ids are not unique!")
			}
		}
	}
}
