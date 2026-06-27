package main

import (
	"fmt"
	"gw2-translator/imgproc"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("images/gw005_clipped.jpg")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Call your modular package logic cleanly
	binaryImg := imgproc.ApplyAdaptiveThreshold(imaging.Grayscale(src), 15)

	err = imaging.Save(binaryImg, "images/gw005_adaptive.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	fmt.Println("Pipeline complete!")
}
