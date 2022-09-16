package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	. "github.com/dnwin/go-ray-tracing/raytrace"
)

func main() {
	f, err := os.Create("./assets/gradient.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := renderImg(f)
	png.Encode(f, img)
}

func renderImg(w io.Writer) image.Image {
	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := MakePoint3(0, 0, 0)
	o := Vec3(origin)
	horizontal := MakeVec3(viewportWidth, 0, 0)
	vertical := MakeVec3(0, viewportHeight, 0)
	//  origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length);
	lowerLeftCorner := o.Sub(horizontal.DivBy(2.0)).Sub(vertical.DivBy(2.0)).Sub(MakeVec3(0, 0, focalLength))

	// Render
	upLeft := image.Point{0, 0}
	lowRight := image.Point{imageWidth - 1, imageHeight - 1}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Printf("\rScanlines remaining: %d ", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(imageHeight-1-j) / float64(imageHeight-1)
			// lower_left_corner + u*horizontal + v*vertical - origin
			dir := lowerLeftCorner.Add(horizontal.MulBy(u)).Add(vertical.MulBy(v)).Sub(o)

			ray := Ray{Orig: origin, Dir: dir}
			clr := rayColor(ray)
			writeColor(img, i, j, clr)
		}
	}
	fmt.Println("\nDone.")

	return img
}

// rayColor linearly blends white and blue depending on the height of the y coordinate
// after scaling the ray direction to unit length (so −1.0<y<1.0).
func rayColor(ray Ray) Color {
	dir := ray.Direction()
	d := dir.UnitVector()

	t := 0.5 * (d.Y() + 1.0)
	white := MakeColor(1.0, 1.0, 1.0)
	blue := MakeColor(0.5, 0.7, 1.0)

	// linear blend/interpolation
	// blendedValue=(1−t)⋅startValue+t⋅endValue
	u := Vec3(white).MulBy((1.0 - t))
	v := Vec3(blue).MulBy(t)
	blend := u.Add(v)

	return Color(blend)
}

func writeColor(img *image.RGBA, x int, y int, clr Color) {
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
