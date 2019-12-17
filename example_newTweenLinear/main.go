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
	"github.com/helmutkemper/iotmaker.platform/mathTween"
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
	//imgHtml := factoryImage.NewHtmlImage(html, browserDocument.SelfDocument, map[string]interface{}{"id":  "player", "src": "./player_big.png"},true,true)
	//factoryImage.NewMultipleSpritesImage(&stage.Canvas, imgHtml,48,60,0,7, 90*time.Millisecond,50,70,48,60, density, densityManager).Crete()

	x := 10
	y := 10
	width := 600
	height := 500

	cl := factoryDraw.NewChartLinear(&stage.Canvas, x, y, width, height, density, densityManager)
	cl.Begin(x, y+height)

	fps.Set(10)

	interactionCurrent := 0.0
	interactionTotal := float64(height - y)
	f := mathTween.KEaseInOutCubic
	for {
		yGraph := f(interactionCurrent, interactionTotal, float64(y+height), float64(y+height))
		cl.Point(int(float64(x+width)*(interactionCurrent/interactionTotal)), int(yGraph))
		interactionCurrent += 1.0

		if yGraph >= float64(y) {
			break
		}
	}

	cl.End()
	cl.Begin(x, y+height)

	factoryTween.NewEaseInOutCubic(
		time.Second*10,
		float64(y+height),
		float64(y),
		func(y, p float64, ars []interface{}) {
			x := ars[0].(int)
			//y := ars[1].(int)
			width := ars[2].(int)
			//height := ars[3].(int)
			cl.Point(int(float64(x+width)*p), int(y))
		},
		func(y float64) {
			cl.End()
		},
		x,
		y,
		width,
		height,
	)

	<-done
}
