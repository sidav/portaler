package backend

type RendererBackend interface {
	Init(int32, int32)
	BeginFrame()
	EndFrame()
	SetInternalResolution(w, h int32)
	Flush()
	SetColor(r, g, b uint8)
	DrawPoint(x, y int32)
	FillRect(x, y, w, h int)
	VerticalLine(x, y0, y1 int)
}
