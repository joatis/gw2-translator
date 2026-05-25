package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// 1. Load the original screenshot from the /images folder
	inputPath := "images/gw005.jpg"
	src, err := imaging.Open(inputPath)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	// 2. Convert the image to a grayscale to remove color data
	grayImg := imaging.Grayscale(src)

	// 3. Apply Thresholding (Binarization)
	// We iterate through each pixel. If it's brite make it pure white. If dark pure black.
	binarizedimg := binarize(grayImg, 128) // 128 is the midpoint threshold (0-255)

	//  4. Save the processed image for Tesseract to use later
	outputPath := "images/processed_runes.png"
	err = imaging.Save(binarizedImg, outputPath)
	if err != nil {
		log.Fatalf("Failed to save image %v", err)
	}

	fmt.Printf("Success! Processed image saved to %s\n", outputPath)
}

// binarize applies a hard threshold to a grayscale image.
func binarize(img image.Image, threshold uint8) image.Image {
	bounds := img.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// Create a new blank grayscale image canvas
	result := image.NewGray(bounds)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Get the current pixels color
			oldColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)

			// Determine if the pixel is above or below our brightness threshold.
			var newColor uint8
			if oldColor.Y > threshold {
				newColor = 255 // Pure White
			} else {
				newColor = 0 // Pure Black
			}

			result.SetGray(x, y, color.Gray{Y: newColor})
		}
	}
	return result
}
