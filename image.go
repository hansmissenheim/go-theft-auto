// File image implements manipulations to the captured image
// used for identification of objects in view
package main

import (
	"image"
	"image/color"
	"math"

	"gocv.io/x/gocv"
)

var (
	// Points of the polygon encasing the region of interest
	vertices = []image.Point{
		image.Point{X: 10, Y: 525},
		image.Point{X: 10, Y: 325},
		image.Point{X: 200, Y: 275},
		image.Point{X: 600, Y: 275},
		image.Point{X: 790, Y: 325},
		image.Point{X: 790, Y: 525},
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
	gocv.GaussianBlur(processedImg, &processedImg, image.Point{5, 5}, 0, 0, 0)
	gocv.Canny(processedImg, &processedImg, 100.0, 300.0)
	gocv.Dilate(processedImg, &processedImg, gocv.GetStructuringElement(0, image.Point{3, 3}))
	processedImg = regionOfInterest(processedImg, vertices)

	lane := findLane(processedImg)
	drawLane(img, lane)
	return img
}

// Uses hough lines and averages the lines with same signed slopes in order
// to best guess the two lines of the current lane
func findLane(img gocv.Mat) [][]float64 {
	lines := gocv.NewMat()
	gocv.HoughLinesPWithParams(img, &lines, 1.0, math.Pi/180.0, 100, 175.0, 5.0)

	right, left := 0, 0
	lane := [][]float64{{0, 0}, {0, 0}}
	if !lines.Empty() {
		for i := 0; i < lines.Rows(); i++ {
			v := lines.GetVeciAt(0, i)
			slope := (float64(v[1]) - float64(v[3])) / (float64(v[0]) - float64(v[2]))
			intercept := float64(v[1]) - slope*float64(v[0])
			if slope > 0 {
				right++
				lane[1][0] += slope
				lane[1][1] += intercept
			} else {
				left++
				lane[0][0] += slope
				lane[0][1] += intercept
			}
		}
		lane[0][0] = lane[0][0] / float64(left)
		lane[0][1] = lane[0][1] / float64(left)
		lane[1][0] = lane[1][0] / float64(right)
		lane[1][1] = lane[1][1] / float64(right)
	}
	return lane
}

// Draws the lane lines on the given image
func drawLane(img gocv.Mat, lane [][]float64) {
	if !math.IsNaN(lane[0][0]) && !math.IsNaN(lane[0][1]) && lane[0][0] != 0 && lane[0][1] != 0 {
		gocv.Line(
			&img,
			image.Point{int(0), int(lane[0][1])},
			image.Point{int((275 - lane[0][1]) / lane[0][0]), int(275)},
			color.RGBA{R: 0, G: 255, B: 0, A: 255},
			5,
		)
	}
	if !math.IsNaN(lane[1][0]) && !math.IsNaN(lane[1][1]) && lane[1][0] != 0 && lane[1][1] != 0 {
		gocv.Line(
			&img,
			image.Point{int((275 - lane[1][1]) / lane[1][0]), int(275)},
			image.Point{int(800), int(lane[1][0]*800.0 + lane[1][1])},
			color.RGBA{R: 0, G: 255, B: 0, A: 255},
			5,
		)
	}
}
