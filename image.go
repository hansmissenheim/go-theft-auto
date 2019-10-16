// File image implements manipulations to the captured image
// used for identification of objects in view
package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

var (
	// Points of the polygon encasing the region of interest
	vertices = []image.Point{
		image.Point{X: 0, Y: 600},
		image.Point{X: 0, Y: 325},
		image.Point{X: 200, Y: 275},
		image.Point{X: 600, Y: 275},
		image.Point{X: 800, Y: 325},
		image.Point{X: 800, Y: 550},
		image.Point{X: 370, Y: 550},
		image.Point{X: 215, Y: 600},
	}
)

// Masks the captured screen by the region we are interested in defined by the
// points above
func regionOfInterest(img gocv.Mat, vertices []image.Point) gocv.Mat {
	maskedImg := gocv.NewMat()
	mask := gocv.NewMatWithSizeFromScalar(
		gocv.NewScalar(0, 0, 0, 255),
		img.Rows(), img.Cols(), img.Type(),
	)
	gocv.FillPoly(&mask, [][]image.Point{vertices}, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	gocv.BitwiseAnd(mask, img, &maskedImg)
	return maskedImg
}

// Proccesses a given image through several manipulations and returns the
// processed image
func process(img gocv.Mat) gocv.Mat {
	processedImg := gocv.NewMat()
	gocv.CvtColor(img, &processedImg, gocv.ColorRGBToGray)
	gocv.Canny(img, &processedImg, 200.0, 300.0)
	return processedImg
}
