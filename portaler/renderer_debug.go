package portaler

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func getDebugColor(i int) (uint8, uint8, uint8) {
	colors := [][3]uint8{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
		{255, 255, 0},
		{0, 255, 255},
		{255, 0, 255},
		{255, 255, 255},
	}
	col := colors[i%len(colors)]
	return col[0], col[1], col[2]
}

func debugPrintf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (r *PortalsRenderer) debugFlush() {
	if r.debugOn {
		r.io.EndFrame()
		r.io.Flush()
		r.io.BeginFrame()
	}
}

func debugPrintOnScreen(x, y int, msg string, args ...interface{}) {
	if time.Now().Second()%2 == 0 {
		rl.DrawText(fmt.Sprintf(msg, args...), int32(x), int32(y), 16, rl.Black)
	} else {
		rl.DrawText(fmt.Sprintf(msg, args...), int32(x), int32(y), 16, rl.White)
	}
}
