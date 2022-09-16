package raytrace

import (
	"image"
	"image/color"
)

func WriteColor(img *image.RGBA, x int, y int, clr Color) {
	v := Vec3(clr)
	r := v.X()
	g := v.Y()
	b := v.Z()
	ir := uint8(255.99 * r)
	ig := uint8(255.99 * g)
	ib := uint8(255.99 * b)

	rgba := color.RGBA{ir, ig, ib, 0xff} // Opaque
	img.Set(x, y, rgba)
}
