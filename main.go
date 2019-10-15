package main

import (
	"image"
	"time"

	"github.com/kbinani/screenshot"
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Hello")

	for {
		screen, _ := screenshot.CaptureRect(image.Rect(566, -1380, 1366, -780))
		img, _ := gocv.ImageToMatRGB(screen)
		window.IMShow(img)
		window.WaitKey(1)

		time.Sleep(5 * time.Second)
		pressKey(w)
		time.Sleep(5 * time.Second)
		releaseKey(w)
	}
}
