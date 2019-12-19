// +build js

//
package main

import (
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform/factoryDraw"
	"github.com/helmutkemper/iotmaker.platform/tween"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = canvas.Stage{}
)

func main() {

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(browserDocument, "stage", 800, 600, density, densityManager)

	x := 10
	y := 10
	width := 600
	height := 500

	cl := factoryDraw.NewChartLinear(&stage.Canvas, x, y, width, height, density, densityManager)
	cl.Begin(x, y+height)

	interactionCurrent := 0.0
	interactionTotal := float64(height - y)
	startValue := float64(y + height)
	endValue := float64(y)
	delta := endValue - startValue
	f := tween.KEaseInOutCubic
	for {
		yGraph := f(interactionCurrent, interactionTotal, startValue, delta)
		cl.Pixel(int(float64(x+width)*(interactionCurrent/interactionTotal))+x, int(yGraph))
		interactionCurrent += 1

		if yGraph <= float64(y) {
			break
		}
	}

	cl.End()

	<-done
}
