// +build js

package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	webBrowserMouse "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryColor"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryGradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryImage"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryPoint"
	"github.com/helmutkemper/iotmaker.platform/abstractType/selectBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
	"github.com/helmutkemper/iotmaker.platform/mouse"
	"image/color"
	"time"
)

var (
	density                                   = 2.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	bx2                                       = &basicBox.BasicBox{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
	html           iotmakerPlatformIDraw.IHtml
)

func main() {

	done := make(chan struct{}, 0)

	html = &Html.Html{}
	browserDocument := factoryBrowserDocument.NewDocument()

	var colorShadow color.RGBA = colornames.DarkblueTransparent
	var blur int = 5
	var offsetX int = 2
	var offsetY int = 2
	var filter iotmakerPlatformIDraw.IFilterShadowInterface = shadow.NewShadowFilter(colorShadow, blur, offsetX, offsetY, density, densityManager)

	//mouse.AddFunctionPointer(bx1.GetAlphaChannel)

	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	colorWhite := factoryColor.NewColorPosition(colornames.Red, 0.5)
	colorBlack := factoryColor.NewColorPosition(colornames.Black, 1)
	colorList := factoryColor.NewColorList(colorBlack, colorWhite)

	coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)

	basicBox.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		"bbox_1",
		20,
		50,
		100,
		100,
		5,
		5,
		nil,
		nil,
		density,
		densityManager,
	)

	bx2 = basicBox.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		"bbox_2",
		20+50,
		50+50,
		100,
		100,
		10,
		8,
		filter,
		gradientFilter,
		density,
		densityManager,
	)

	selectBox.NewResizeBoxFromBasicBox(bx2, -3, -3, 6, 6, 1, density, densityManager)

	img := factoryImage.NewHtmlImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "player",
			"src": "./player_big.png",
		},
		true,
		true,
	)

	i := factoryImage.NewMultipleSpritesImage(
		&stage.Canvas,
		img,
		48,
		60,
		0,
		7,
		80*time.Millisecond,
		45,
		70,
		48,
		60,
		density,
		densityManager,
	)
	i.Crete()

	//fmt.Printf("over: %v\n", bx.GetAlphaChannel(0, 100))

	/*abstractType.NewLineTo(
		&stage.Canvas,
		"line",
		1.0,
		0,
		0,
		300,
		300,
		1,
	)*/

	browserDocument.SetMouseMoveListener(webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	mouse.SetPlatform(mouse.KPlatformWebBrowser)
	mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBySimpleBox, bateu)

	<-done
}

var lastCursor bool

func bateu(x, y int, collision bool) {
	if collision == false {
		webBrowserMouse.SetCursor(stage.SelfElement, mouse.KCursorAuto)
	} else {
		webBrowserMouse.SetCursor(stage.SelfElement, mouse.KCursorColResize)
	}

	lastCursor = collision

	//fmt.Printf("bateu: %v\n", collision)
}
