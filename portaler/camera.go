package portaler

import (
	"fmt"
	"math"
)

type camera struct {
	x, y              float64
	radians           float64
	distToScreenPlane float64
	Height            float64

	sin, cos float64
}

func NewCamera() *camera {
	cam := &camera{distToScreenPlane: 0.5, Height: 0.5}
	cam.Rotate(0) // just to update sin/cos values
	return cam
}

func (c *camera) Rotate(r float64) {
	c.radians += r
	c.sin = math.Sin(c.radians)
	c.cos = math.Cos(c.radians)
}

func (c *camera) GetDirectionVector() (float64, float64) {
	return math.Sin(c.radians + 3.1416/2), math.Cos(c.radians + 3.1416/2)
}

func (c *camera) GetInfo() string {
	return fmt.Sprintf("Cam at %.2f,%.2f; rotation %.2f", c.x, c.y, c.radians)
}

func (c *camera) GetCoords() (float64, float64) {
	return c.x, c.y
}

func (c *camera) SetCoords(x, y float64) {
	c.x, c.y = x, y
}

func (c *camera) transformPointToCameraSpace(x, y float64) (rx, ry float64) {
	rx = x - c.x
	ry = y - c.y
	rxTemp := rx
	rx = rx*c.cos - ry*c.sin
	ry = rxTemp*c.sin + ry*c.cos
	return
}

func (c *camera) transformLinedefToCameraSpace(l *linedef) (x1, y1, x2, y2 float64) {
	x1, y1 = c.transformPointToCameraSpace(l.x1, l.y1)
	x2, y2 = c.transformPointToCameraSpace(l.x2, l.y2)
	return
}
