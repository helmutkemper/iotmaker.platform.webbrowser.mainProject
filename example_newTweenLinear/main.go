// +build js

//
package main

import (
	"fmt"
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
	//imgHtml := factoryImage.NewHtmlImage(html, browserDocument.SelfDocument, map[string]interface{}{"id":  "player", "src": "./player_big.png"},true,true)
	//factoryImage.NewMultipleSpritesImage(&stage.Canvas, imgHtml,48,60,0,7, 90*time.Millisecond,50,70,48,60, density, densityManager).Crete()

	cl := factoryDraw.NewChartLinear(&stage.Canvas, 10, 10, 600, 500, density, densityManager)
	cl.Begin(10, 510)
	x := 10.0
	step := 630.0 / float64(fps.Get()) / 10.0
	factoryTween.NewEaseInOutCubic(
		time.Second*10,
		510,
		10,
		func(y float64) {
			cl.Point(int(x), int(y))
			x += step
		},
		func(y float64) {
			fmt.Printf("final y: %v\n", y)
			cl.End()
		},
	)

	<-done
}
