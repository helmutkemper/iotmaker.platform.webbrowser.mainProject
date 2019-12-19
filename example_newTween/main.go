// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform/factoryDraw"
	"github.com/helmutkemper/iotmaker.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"time"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
	html           iotmakerPlatformIDraw.IHtml
)

func main() {

	done := make(chan struct{}, 0)

	html = &Html.Html{}
	browserDocument := factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(browserDocument, "stage", 800, 600, density, densityManager)

	x := 10
	y := 10
	width := 600
	height := 500

	cl := factoryDraw.NewChartLinear(&stage.Canvas, x, y, width, height, density, densityManager)
	cl.Begin(x, y+height)

	fps.Set(60)

	factoryTween.NewEaseInOutSine(
		time.Second*5,
		float64(y+height),
		float64(y),
		func(y, p float64, ars []interface{}) {
			x := ars[0].(int)
			width := ars[1].(int)
			cl.Pixel(x+int(float64(width)*p), int(y))
		},
		func(y float64) {
			cl.End()
		},
		x,
		width,
	)

	<-done
}
