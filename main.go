package main

import (
	"fmt"
	"portalrenderer/backend"
	"portalrenderer/portaler"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	io backend.RendererBackend
)

func main() {
	const sw, sh = 1400, 800
	const pixelSize = 4
	io = &backend.RaylibBackend{}
	io.Init(sw, sh)
	io.SetInternalResolution(sw/pixelSize, sh/pixelSize)

	cam := portaler.NewCamera()
	scene := &portaler.Scene{}
	scene.Textures = append(scene.Textures, *portaler.InitTextureFromImageFile("wall.png"))
	scene.Textures = append(scene.Textures, *portaler.InitTextureFromImageFile("upperwall.png"))
	scene.Textures = append(scene.Textures, *portaler.InitTextureFromImageFile("lowerwall.png"))
	scene.InitTesting2()

	renderer := portaler.NewRenderer(io, sw/pixelSize, sh/pixelSize, scene)

	for !rl.WindowShouldClose() {
		renderer.Render(scene, cam)
		if rl.IsKeyDown(rl.KeyLeft) {
			cam.Rotate(0.05)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			cam.Rotate(-0.05)
		}
		if rl.IsKeyDown(rl.KeyUp) {
			cx, cy := cam.GetCoords()
			vx, vy := cam.GetDirectionVector()
			factor := 0.1
			cam.SetCoords(cx+vx*factor, cy+vy*factor)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			cx, cy := cam.GetCoords()
			vx, vy := cam.GetDirectionVector()
			factor := -0.1
			cam.SetCoords(cx+vx*factor, cy+vy*factor)
		}
		if rl.IsKeyDown(rl.KeyComma) {
			cam.Rotate(-3.1416 / 2)
			cx, cy := cam.GetCoords()
			vx, vy := cam.GetDirectionVector()
			factor := -0.1
			cam.SetCoords(cx+vx*factor, cy+vy*factor)
			cam.Rotate(3.1416 / 2)
		}
		if rl.IsKeyDown(rl.KeyPeriod) {
			cam.Rotate(3.1416 / 2)
			cx, cy := cam.GetCoords()
			vx, vy := cam.GetDirectionVector()
			factor := -0.1
			cam.SetCoords(cx+vx*factor, cy+vy*factor)
			cam.Rotate(-3.1416 / 2)
		}
		if rl.IsKeyDown(rl.KeyA) {
			cam.Height += 0.025
			fmt.Printf("CamH is now %.2f\n", cam.Height)
		}
		if rl.IsKeyDown(rl.KeyZ) {
			cam.Height -= 0.025
			fmt.Printf("CamH is now %.2f\n", cam.Height)
		}
		if rl.IsKeyDown(rl.KeyDelete) {
			renderer.DebugOn = !renderer.DebugOn
			if !renderer.DebugOn {
				time.Sleep(500 * time.Millisecond)
			}
		}
	}
	defer rl.CloseWindow()
}
