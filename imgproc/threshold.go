package imgproc

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
)

// adaptiveThreshold implements a local adaptive algorithiom.
// I uses a box blur to estimate the local neighborhood's average.
func ApplyAdaptiveThreshold(img image.Image, windowSize int) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// Create a new blank grayscale image for the result
	result := image.NewGray(bounds)

	// 1. Create a "Local Average" map using a box blur.
	// imaging.Blur computes a local average effectively for each pixel.
	localAvgImg := imaging.Blur(img, float64(windowSize/2)) // WindowSize needs finetuning.

	// C value is an offset to subtract from the local average.
	// Tunable parameter. 5-15 is often good.
	const constantOffsetC = 7

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// 2. Get the pixels actual brightness
			srcPixel := color.GrayModel.Convert(img.At(x, y)).(color.Gray).Y

			// 3. Get the local average brightness
			localAvg := color.GrayModel.Convert(localAvgImg.At(x, y)).(color.Gray).Y

			// 4. Calculate the ADAPTIVE threshold for this pixel
			pixelThreshold := localAvg - constantOffsetC

			// 5. Apply the result
			var newColor uint8
			if srcPixel < pixelThreshold {
				newColor = 0 // Pure Black (rune/edge)
			} else {
				newColor = 255 // Pure White (background/shadow)
			}
			result.SetGray(x, y, color.Gray{Y: newColor})
		}
	}
	return result
}
