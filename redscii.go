package redscii

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

var (
	asciiMap = []string{" ", ".", "-", "=", "+", ":", "!", "8", "0", "#"}
)

func ASCIIfy(img image.Image) {
	img.ColorModel()
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			divider := 0xFF / (len(asciiMap) - 1)
			r, _, _, _ := img.At(x, y).RGBA()
			R := int(r>>8) / divider

			fmt.Print(asciiMap[R])
			fmt.Print(asciiMap[R])
		}
		fmt.Println()
	}
}

func DownscaleImage(img image.Image, scale float64) image.Image {
	// If the scale is larger than one, don't do any processing and just return the original image.
	if scale >= 1 {
		return img
	}

	imgBounds := img.Bounds()
	divider := int(math.Ceil(1 / scale))

	newImageBounds := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: imgBounds.Max.X / divider, Y: imgBounds.Max.Y / divider}}
	newImage := newImage(newImageBounds)

	// Downscale
	for y := imgBounds.Min.Y; y < (imgBounds.Max.Y / divider); y++ {
		for x := imgBounds.Min.X; x < (imgBounds.Max.X / divider); x++ {
			oX := x * divider
			oY := y * divider
			subRect := image.Rectangle{Min: image.Point{oX, oY}, Max: image.Point{oX + divider, oY + divider}}

			// TODO: This is an essumption. Figure out a way to detect this correctly at runtime.
			subImg := img.(*image.NRGBA).SubImage(subRect)
			avgColor := getAverageImgColor(subImg)

			newImage.Set(x, y, avgColor)
		}
	}
	return newImage
}

func GreyScaleImage(img image.Image) image.Image {

	// TODO: Use ITU-R BT7.09 luminance weights
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := getRawPixelRGBA(img.At(x, y))
			avg := uint8(((r + g + b) / 3) & 0xFF)
			alpha := uint8(a & 0xFF)

			img.(*image.RGBA).Set(x, y, color.RGBA{avg, avg, avg, alpha})
		}
	}

	return img
}

func newImage(rect image.Rectangle) *image.RGBA {
	return image.NewRGBA(rect)
}

func getAverageImgColor(img image.Image) color.RGBA {
	var R, G, B, A uint32
	cAvg := color.RGBA{}

	xmin := img.Bounds().Min.X
	ymin := img.Bounds().Min.Y
	xmax := img.Bounds().Max.X
	ymax := img.Bounds().Max.Y

	for y := ymin; y < ymax; y++ {
		for x := xmin; x < xmax; x++ {
			r, g, b, a := getRawPixelRGBA(img.At(x, y))

			R += r
			G += g
			B += b
			A += a
		}
	}

	divider := (xmax - xmin) * (ymax - ymin)

	cAvg.R = uint8((int(R) / divider) & 0xFF)
	cAvg.G = uint8((int(G) / divider) & 0xFF)
	cAvg.B = uint8((int(B) / divider) & 0xFF)
	cAvg.A = uint8((int(A) / divider) & 0xFF)

	return cAvg
}

func getRawPixelRGBA(c color.Color) (uint32, uint32, uint32, uint32) {
	r, g, b, a := c.RGBA()
	switch c.(type) {
	case color.NRGBA:
		// Transform these to non aphla-premultiplied values.
		a = a >> 8
		r *= 0xFF
		r /= a
		r = r >> 8

		g *= 0xFF
		g /= a
		g = g >> 8

		b *= 0xFF
		b /= a
		b = b >> 8

	case color.RGBA:
		// Values are assumed to be already premultiplied, return as is.
		a = a >> 8
		r = r >> 8
		g = g >> 8
		b = b >> 8

	default:
		// Don't do any manipulations for other color formats.
	}

	return r, g, b, a
}
