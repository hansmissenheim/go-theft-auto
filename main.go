package main

import (
	"image"

	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Hello")

	for {
		screen, _ := screenshot.CaptureRect(
			image.Rect(screenRes[0], screenRes[1], screenRes[2], screenRes[3]),
		)
		img, _ := gocv.ImageToMatRGB(screen)
		window.IMShow(process(img))
		window.WaitKey(1)
	}
}
