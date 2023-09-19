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
	aspectRatio        float64
	minimapW, minimapH int
	minimapScale       float64

	totalTimeForRendering time.Duration
	meanRenderingTime     time.Duration
	meanTimeEachFrames    int
	totalFramesRendered   int

	scene                *Scene
	renderedSectorsTable []bool   // TODO: get rid of it somehow
	renderedColumnsTable [][2]int // first - bottom-to-top (lower) rendered coord, second - top-to-bottom (upper)
	filledColumnsTable   []bool   // true if column can not be drawn into

	DebugOn bool
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
	if len(r.renderedColumnsTable) != r.screenW {
		r.renderedColumnsTable = make([][2]int, r.screenW)
		r.filledColumnsTable = make([]bool, r.screenW)
	}
	// reset columns table
	for x := range r.renderedColumnsTable {
		r.renderedColumnsTable[x][0] = r.screenH
		r.renderedColumnsTable[x][1] = 0
	}
	for x := range r.filledColumnsTable {
		r.filledColumnsTable[x] = false
	}

	r.io.BeginFrame()
	rl.ClearBackground(rl.Black)

	debugPrintf("======= BEGIN FRAME %d =======\n", r.totalFramesRendered+1)

	t := time.Now()

	r.renderScene(s, c)

	r.totalFramesRendered++
	r.totalTimeForRendering += time.Since(t)
	debugPrintf("FRAME ENDED. Passed %8v\n", time.Since(t))
	if r.totalFramesRendered%r.meanTimeEachFrames == 0 {
		r.meanRenderingTime = r.totalTimeForRendering / time.Duration(r.meanTimeEachFrames)
		r.totalTimeForRendering = 0
	}
	rl.DrawText(fmt.Sprintf("Mean frame time %-8s", r.meanRenderingTime), int32(r.minimapW)+2, 0, 21, rl.White)
	rl.DrawText(c.GetInfo(), 0, int32(r.screenH-32), 30, rl.Blue)

	r.io.EndFrame()
	r.io.Flush()
}

func (r *PortalsRenderer) renderScene(s *Scene, c *camera) {
	r.resetRenderedSectorsTable()
	screenArea := newTrapezoid(0, r.screenH-1, 0, r.screenW-1, r.screenH-1, 0) // represents whole screen
	x, y := c.GetDirectionVector()
	x *= c.distToScreenPlane * 2.0 / 3.0
	y *= c.distToScreenPlane * 2.0 / 3.0
	sectorWithCameraPlaneIn := s.FindSectorWithCoordinates(c.x+x, c.y+y)
	if sectorWithCameraPlaneIn != nil {
		debugPrintf("Camera sector %d\n", sectorWithCameraPlaneIn.id)
		r.renderSector(sectorWithCameraPlaneIn, c, screenArea)
	} else {
		// do not crash
		debugPrintf("Camera not found in any sector\n")
		r.renderSector(s.sectors[0], c, screenArea)
	}
	r.drawMinimap(s, c)
}

func (r *PortalsRenderer) renderSector(currentSector *sector, c *camera, screenArea *trapezoid) {
	debugPrintf("Rendering sector #%d: ", currentSector.id)
	r.renderedSectorsTable[currentSector.id] = true
	// render non-portals first
	for _, l := range currentSector.lines {
		if !l.isPortal {
			debugPrintf("S%d line rendered, ", currentSector.id)
			r.renderLinedef(l, currentSector, c, screenArea)
		}
	}
	// render portals only
	for _, l := range currentSector.lines {
		if l.isPortal {
			// skip if the sector was already rendered
			if r.wasSectorRendered(l.getNextSectorFrom(currentSector)) {
				debugPrintf("S%d portal to S%d skipped, ", currentSector.id, l.getNextSectorFrom(currentSector).id)
				continue
			}
			debugPrintf("S%d portal rendering, ", currentSector.id)
			r.renderPortalLinedef(l, currentSector, c, screenArea)
		}
	}
	debugPrintf("\nSector #%d rendered.\n", currentSector.id)
}
