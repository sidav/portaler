package portaler

import (
	"image"
	"image/png"
	"os"
)

type Texture struct {
	Bitmap image.Image
	W, H   int
}

func InitTextureFromImageFile(filename string) *Texture {
	// first, read the image file
	imgfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()
	img, err := png.Decode(imgfile)
	imgfile.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	// return the "reader carriage" to zero
	imgfile.Seek(0, 0)
	// read file config (needed for w, h)
	cfg, _, err := image.DecodeConfig(imgfile)
	if err != nil {
		panic(err)
	}
	// init the map if needed

	return &Texture{
		Bitmap: img,
		W:      cfg.Width,
		H:      cfg.Height,
	}
}
