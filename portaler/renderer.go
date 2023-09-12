package portaler

import (
	"fmt"
	"portalrenderer/backend"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PortalsRenderer struct {
	io                 *backend.RaylibBackend
	screenW, screenH   int
	minimapW, minimapH int
	minimapScale       float64

	totalTimeForRendering time.Duration
	totalFrames           int

	scene                *Scene
	renderedSectorsTable []bool // TODO: get rid of it somehow
}

func (r *PortalsRenderer) wasSectorRendered(s *sector) bool {
	return r.renderedSectorsTable[s.id]
}

func (r *PortalsRenderer) resetRenderedSectorsTable() {
	for i := range r.renderedSectorsTable {
		r.renderedSectorsTable[i] = false
	}
}

func (r *PortalsRenderer) Render(s *Scene, c *camera) {
	if rl.IsWindowResized() {
		r.io.SetInternalResolution(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	}
	r.io.BeginFrame()
	rl.ClearBackground(rl.Black)

	t := time.Now()

	r.renderScene(s, c)

	r.totalFrames++
	r.totalTimeForRendering += time.Since(t)
	debugPrintf("FRAME ENDED. Passed %8v\n", time.Since(t))
	rl.DrawText(fmt.Sprintf("Mean frame time %-8s", r.totalTimeForRendering/time.Duration(r.totalFrames)), int32(r.minimapW)+2, 0, 21, rl.White)
	rl.DrawText(c.GetInfo(), 0, int32(r.screenH-32), 30, rl.Blue)

	r.io.EndFrame()
	r.io.Flush()
}

func (r *PortalsRenderer) renderScene(s *Scene, c *camera) {
	r.resetRenderedSectorsTable()
	screenArea := newTrapezoid(0, r.screenH-1, 0, r.screenW-1, r.screenH, 0) // represents whole screen
	x, y := c.GetDirectionVector()
	x *= c.distToScreenPlane * 2.0 / 3.0
	y *= c.distToScreenPlane * 2.0 / 3.0
	sectorWithCameraPlaneIn := s.FindSectorWithCoordinates(c.x+x, c.y+y)
	if sectorWithCameraPlaneIn != nil {
		r.renderSector(sectorWithCameraPlaneIn, c, screenArea)
	} else {
		// do not crash
		r.renderSector(s.sectors[0], c, screenArea)
	}
	r.drawMinimap(s, c)
}

func (r *PortalsRenderer) renderSector(currentSector *sector, c *camera, screenArea *trapezoid) {
	debugPrintf("Rendering sector #%d: ", currentSector.id)
	r.renderedSectorsTable[currentSector.id] = true
nextLinedef:
	for _, l := range currentSector.lines {
		if l.isPortal {
			// skip if the sector was already rendered
			if r.wasSectorRendered(l.getNextSectorFrom(currentSector)) {
				debugPrintf("S%d portal to S%d skipped, ", currentSector.id, l.getNextSectorFrom(currentSector).id)
				continue nextLinedef
			}
			r.renderPortalLinedef(l, currentSector, c, screenArea)
		} else {
			debugPrintf("S%d line rendered, ", currentSector.id)
			r.renderLinedef(l, currentSector, c, screenArea)
		}
	}
	debugPrintf("\nSector #%d rendered.\n", currentSector.id)
}
