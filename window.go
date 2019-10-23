// File window creates the screen capture parameters based on the set
// game resolution and monitor the game is played on.
package main

import (
	"fmt"
	"strings"

	"github.com/kbinani/screenshot"
)

var (
	monitor   = 0
	gameRes   = []int{800, 600}
	screenRes = []int{0, 0, 0, 0}
)

func init() {
	bounds := screenshot.GetDisplayBounds(monitor)
	boundsString, dx, dy := bounds.String(), bounds.Dx(), bounds.Dy()

	boundsString = boundsString[1 : len(boundsString)-1]
	boundsString = strings.Replace(boundsString, ")-(", ",", 1)
	fmt.Sscanf(boundsString, "%d,%d,%d,%d",
		&screenRes[0], &screenRes[1], &screenRes[2], &screenRes[3],
	)

	screenRes[0] = (dx-gameRes[0])/2 + screenRes[0]
	screenRes[1] = (dy-gameRes[1])/2 + screenRes[1]
	screenRes[2] = (dx-gameRes[0])/-2 + screenRes[2]
	screenRes[3] = (dy-gameRes[1])/-2 + screenRes[3]
}
