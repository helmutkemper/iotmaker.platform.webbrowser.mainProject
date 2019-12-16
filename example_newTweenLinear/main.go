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

	cl := factoryDraw.NewChartLinear(&stage.Canvas, 10, 10, 200, 100, density, densityManager)
	cl.Begin(10, 110)
	initial := 10
	step := 100 / 200
	factoryTween.NewTweenEaseInQuadratic(
		time.Second*10,
		210,
		10,
		func(value float64) {
			cl.Point(int(value), initial)
			initial += step
		},
		func(value float64) {
			cl.End()
		},
	)

	<-done
}
