package main

import "fmt"

func main() {

	// Image

	const imageWidth int = 256
	const imageHeight int = 256

	// Render as PPM

	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.25

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			fmt.Printf("%d %d %d\n", ir, ig, ib)
		}
	}
}
