package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Write PNG to file
	f, err := os.Create("./assets/output-image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := renderImg()
	png.Encode(f, img)
}

// renderImg generates a rainbow PNG image.
func renderImg() image.Image {
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
			// Red goes from fully off (0.0) to fully on (1.0) from left to right
			r := float64(x) / float64(imageWidth-1)
			// Green goes from fully off (0.0) to fully on (1.0) from bottom to top
			g := float64(imageHeight-1-y) / float64(imageHeight-1)
			b := 0.25

			ir := uint8(255.99 * r)
			ig := uint8(255.99 * g)
			ib := uint8(255.99 * b)

			clr := color.RGBA{ir, ig, ib, 0xff} // Opaque

			img.Set(x, y, clr)
		}
	}

	fmt.Fprintf(stdout, "\nDone.\n")
	stdout.Flush()

	return img
}
