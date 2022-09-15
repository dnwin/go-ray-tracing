package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

func main() {
	// Write PNG to file
	const fn string = "/tmp/image.png"
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	renderImg(f)
}

// renderImg generates a rainbow PNG image.
func renderImg(w io.Writer) {
	// Progress Indicator
	stdout := bufio.NewWriter(os.Stdout)

	// Image
	const imageWidth int = 256
	const imageHeight int = 256

	upLeft := image.Point{0, 0}
	lowRight := image.Point{imageWidth - 1, imageHeight - 1}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Render as PNG
	for y := imageHeight - 1; y >= 0; y-- {
		fmt.Fprintf(stdout, "\rScanlines remaining: %d ", y)
		stdout.Flush()

		for x := 0; x < imageWidth; x++ {
			r := float64(x) / float64(imageWidth-1)
			g := float64(y) / float64(imageHeight-1)
			b := 0.25

			ir := uint8(255.99 * r)
			ig := uint8(255.99 * g)
			ib := uint8(255.99 * b)

			clr := color.RGBA{ir, ig, ib, 0xff} // Opaque

			img.Set(x, y, clr)
		}
	}
	png.Encode(w, img)

	fmt.Fprintf(stdout, "\nDone.\n")
	stdout.Flush()
}
