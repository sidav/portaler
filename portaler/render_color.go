package portaler

func (rend *PortalsRenderer) setColorWithBrightness(r, g, b, brightness uint8) {
	r = uint8(min(255, int(r)*int(brightness)/255))
	g = uint8(min(255, int(g)*int(brightness)/255))
	b = uint8(min(255, int(b)*int(brightness)/255))
	rend.io.SetColor(r, g, b)
}
