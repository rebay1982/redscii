package redscii

import (
	"image"
	"image/color"
	"testing"
)

func TestGreyScaleImage(t *testing.T) {
	bounds := image.Rect(0, 0, 5, 5)
	testCases := []struct {
		name     string
		image    image.Image
		expected image.Image
	}{
		{
			name:     "red_to_grey",
			image:    generateCustomSizeImage(bounds, color.RGBA{R: 255, G: 0, B: 0, A: 255}),
			expected: generateCustomSizeImage(bounds, color.RGBA{R: 85, G: 85, B: 85, A: 255}),
		},
		{
			name:     "green_to_grey",
			image:    generateCustomSizeImage(bounds, color.RGBA{R: 0, G: 255, B: 0, A: 255}),
			expected: generateCustomSizeImage(bounds, color.RGBA{R: 85, G: 85, B: 85, A: 255}),
		},
		{
			name:     "blue_to_grey",
			image:    generateCustomSizeImage(bounds, color.RGBA{R: 0, G: 0, B: 255, A: 255}),
			expected: generateCustomSizeImage(bounds, color.RGBA{R: 85, G: 85, B: 85, A: 255}),
		},
		{
			name:     "black_to_grey",
			image:    generateCustomSizeImage(bounds, color.RGBA{R: 0, G: 0, B: 0, A: 255}),
			expected: generateCustomSizeImage(bounds, color.RGBA{R: 0, G: 0, B: 0, A: 255}),
		},
		{
			name:     "white_to_grey",
			image:    generateCustomSizeImage(bounds, color.RGBA{R: 255, G: 255, B: 255, A: 255}),
			expected: generateCustomSizeImage(bounds, color.RGBA{R: 255, G: 255, B: 255, A: 255}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := GreyScaleImage(tc.image)

			for y := 0; y < bounds.Max.Y; y++ {
				for x := 0; x < bounds.Max.X; x++ {

					gotColor := got.At(x, y)
					expColor := tc.expected.At(x, y)

					gr, gg, gb, ga := gotColor.RGBA()
					er, eg, eb, ea := expColor.RGBA()

					if er != gr || eg != gg || eb != gb || ea != ga {
						t.Errorf("Expected color %v, got %v at pixel %d, %d.", expColor, gotColor, x, y)
					}
				}
			}
		})
	}
}

func TestGetAverageImgColor(t *testing.T) {
	testCases := []struct {
		name     string
		image    image.Image
		expected color.RGBA
	}{
		{
			name:     "single_pixel_image",
			image:    generateCustomSizeImage(image.Rect(0, 0, 1, 1), color.RGBA{R: 255, G: 0, B: 0, A: 255}),
			expected: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		},
		{
			name:     "red_horizontal_gradient",
			image:    generateHorizontalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 255, G: 0, B: 0, A: 255}),
			expected: color.RGBA{R: 127, G: 0, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "green_horizontal_gradient",
			image:    generateHorizontalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 255, B: 0, A: 255}),
			expected: color.RGBA{R: 0, G: 127, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "blue_horizontal_gradient",
			image:    generateHorizontalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 0, B: 255, A: 255}),
			expected: color.RGBA{R: 0, G: 0, B: 127, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "black_horizontal_gradient",
			image:    generateHorizontalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 0, B: 0, A: 255}),
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "white_horizontal_gradient",
			image:    generateHorizontalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 255, G: 255, B: 255, A: 255}),
			expected: color.RGBA{R: 127, G: 127, B: 127, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "red_vertical_gradient",
			image:    generateVerticalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 255, G: 0, B: 0, A: 255}),
			expected: color.RGBA{R: 127, G: 0, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "green_vertical_gradient",
			image:    generateVerticalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 255, B: 0, A: 255}),
			expected: color.RGBA{R: 0, G: 127, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "blue_vertical_gradient",
			image:    generateVerticalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 0, B: 255, A: 255}),
			expected: color.RGBA{R: 0, G: 0, B: 127, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "black_horizontal_gradient",
			image:    generateVerticalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 0, G: 0, B: 0, A: 255}),
			expected: color.RGBA{R: 0, G: 0, B: 0, A: 255}, // Gradient's average will always be the 255 / 2
		},
		{
			name:     "white_horizontal_gradient",
			image:    generateVerticalGradientImage(image.Rect(0, 0, 4, 4), color.RGBA{R: 255, G: 255, B: 255, A: 255}),
			expected: color.RGBA{R: 127, G: 127, B: 127, A: 255}, // Gradient's average will always be the 255 / 2
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := getAverageImgColor(tc.image)

			if tc.expected.R != got.R || tc.expected.G != got.G || tc.expected.B != got.B || tc.expected.A != got.A {
				t.Errorf("Expected color %v, got %v.", tc.expected, got)
			}
		})
	}
}

func TestGetRawPixelRGBA(t *testing.T) {
	var eR, eG, eB, eA uint8 = 255, 127, 64, 32
	testCases := []struct {
		name  string
		color color.Color
	}{
		{
			name:  "basic_rgba_test",
			color: color.RGBA{R: eR, G: eG, B: eB, A: eA},
		},
		{
			name:  "basic_nrgba_test",
			color: color.NRGBA{R: eR, G: eG, B: eB, A: eA},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, g, b, a := getRawPixelRGBA(tc.color)

			if uint8(r) != eR || uint8(g) != eG || uint8(b) != eB || uint8(a) != eA {
				t.Errorf("Expected color {%d, %d, %d, %d}, got {%d, %d, %d, %d}.", eR, eG, eB, eA, r, g, b, a)
			}
		})
	}
}

func generateCustomSizeImage(r image.Rectangle, c color.RGBA) image.Image {
	testImage := image.NewRGBA(r)
	xm := r.Max.X
	ym := r.Max.Y
	for y := 0; y < ym; y++ {
		for x := 0; x < xm; x++ {
			testImage.Pix[x*4+(y*testImage.Stride)] = c.R
			testImage.Pix[x*4+(y*testImage.Stride)+1] = c.G
			testImage.Pix[x*4+(y*testImage.Stride)+2] = c.B
			testImage.Pix[x*4+(y*testImage.Stride)+3] = c.A
		}
	}
	return testImage
}

func generateHorizontalGradientImage(r image.Rectangle, c color.RGBA) image.Image {
	testImage := image.NewNRGBA(r)
	xmin := r.Min.X
	xmax := r.Max.X
	ymin := r.Min.Y
	ymax := r.Max.Y

	xdelta := xmax - xmin - 1
	for y := ymin; y < ymax; y++ {
		for x := xmin; x < xmax; x++ {
			c := color.RGBA{
				R: uint8((x * int(c.R) / xdelta) & 0xFF),
				G: uint8((x * int(c.G) / xdelta) & 0xFF),
				B: uint8((x * int(c.B) / xdelta) & 0xFF),
				A: 0xFF,
			}
			testImage.Pix[x*4+(y*testImage.Stride)] = c.R
			testImage.Pix[x*4+(y*testImage.Stride)+1] = c.G
			testImage.Pix[x*4+(y*testImage.Stride)+2] = c.B
			testImage.Pix[x*4+(y*testImage.Stride)+3] = c.A
		}
	}
	return testImage
}

func generateVerticalGradientImage(r image.Rectangle, c color.RGBA) image.Image {
	testImage := image.NewNRGBA(r)
	xmin := r.Min.X
	xmax := r.Max.X
	ymin := r.Min.Y
	ymax := r.Max.Y

	ydelta := ymax - ymin - 1
	for y := ymin; y < ymax; y++ {
		c := color.RGBA{
			R: uint8((y * int(c.R) / ydelta) & 0xFF),
			G: uint8((y * int(c.G) / ydelta) & 0xFF),
			B: uint8((y * int(c.B) / ydelta) & 0xFF),
			A: 0xFF,
		}
		for x := xmin; x < xmax; x++ {
			testImage.Pix[x*4+(y*testImage.Stride)] = c.R
			testImage.Pix[x*4+(y*testImage.Stride)+1] = c.G
			testImage.Pix[x*4+(y*testImage.Stride)+2] = c.B
			testImage.Pix[x*4+(y*testImage.Stride)+3] = c.A
		}
	}
	return testImage
}
