//go:build js
// +build js

//
package main

import (
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.platform/factoryDraw"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/tween"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = canvas.Stage{}

	cl *draw.ChartLinear
	id string
)

func main() {

	fps.Set(1)

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(browserDocument, "stage", 800, 600, density, densityManager)

	x := 10.0
	y := 10.0
	width := 600.0
	height := 500.0

	cl = factoryDraw.NewChartLinear(&stage.Canvas, x, y, width, height, density, densityManager)

	id = stage.Add(Draw)

	<-done
}

func Draw() {
	x := 10.0
	y := 10.0
	width := 600.0
	height := 500.0

	cl.Begin(x, y+height)

	interactionCurrent := 0.0
	interactionTotal := height - y
	startValue := y + height
	endValue := y
	delta := endValue - startValue
	f := tween.SelectRandom()
	for {
		yGraph := f(interactionCurrent, interactionTotal, startValue, delta)
		cl.Pixel((x+width)*(interactionCurrent/interactionTotal)+x, yGraph)
		interactionCurrent += 1

		if interactionCurrent/interactionTotal >= 1.0 {
			break
		}
	}

	//stage.Remove(id)
	cl.End()
}
