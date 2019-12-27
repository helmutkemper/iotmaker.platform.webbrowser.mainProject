// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/basic"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/fontFamily"
	webBrowserMouse "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/draw"
	"github.com/helmutkemper/iotmaker.platform/abstractType/genericTypes"
	"github.com/helmutkemper/iotmaker.platform/abstractType/selectBox"
	"github.com/helmutkemper/iotmaker.platform/factoryColor"
	"github.com/helmutkemper/iotmaker.platform/factoryDraw"
	"github.com/helmutkemper/iotmaker.platform/factoryFont"
	"github.com/helmutkemper/iotmaker.platform/factoryGradient"
	"github.com/helmutkemper/iotmaker.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.platform/factoryPoint"
	"github.com/helmutkemper/iotmaker.platform/factoryShadow"
	"github.com/helmutkemper/iotmaker.platform/factoryText"
	"github.com/helmutkemper/iotmaker.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"github.com/helmutkemper/iotmaker.platform/mouse"
	"image/color"
	"time"
)

var (
	density                                   = 3.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	bx2                                       = &draw.BasicBox{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
)

var html iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace interface{}
var imgPlayer interface{}

func prepareDataBeforeRun() {

	html = &Html.Html{}
	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	imgSpace = factoryBrowserImage.NewImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./spacecraft.png",
		},
		true,
		false,
	)

	imgPlayer = factoryBrowserImage.NewImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "player",
			"src": "./player_big.png",
		},
		true,
		false,
	)
}

func main() {

	//todo: canvasDrawImage()

	done := make(chan struct{}, 0)
	fps.Set(60)
	prepareDataBeforeRun()

	var colorShadow = colornames.BlackHalfTransparent
	var blur float64 = 10
	var offsetX float64 = 2
	var offsetY float64 = 2
	var shadowFilter = factoryShadow.NewShadowFilter(colorShadow, blur, offsetX, offsetY, density, densityManager)

	//mouse.AddFunctionPointer(bx1.GetAlphaChannel)

	colorWhite := factoryColor.NewColorPosition(colornames.Red, 0.5)
	colorBlack := factoryColor.NewColorPosition(colornames.Black, 1)
	colorList := factoryColor.NewColorList(colorBlack, colorWhite)

	coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)

	fontText := factoryFont.NewFont(15, "px", color.RGBA{}, fontFamily.KVerdana, density, densityManager)
	factoryFont.SetGlobal(
		&stage.Canvas,
		fontText,
	)

	factoryText.NewText(
		&stage.Canvas,
		shadowFilter,
		gradientFilter,
		nil,
		"Olá Mundo!",
		25,
		20,
		density,
		densityManager,
	)

	fontText = factoryFont.NewFont(15, "px", color.RGBA{}, fontFamily.KArial, density, densityManager)
	factoryText.NewTextWithFont(
		&stage.Canvas,
		nil,
		nil,
		colornames.Cadetblue,
		fontText,
		"Olá Mundo!",
		25,
		40,
		density,
		densityManager,
	)

	_ = basic.Sprite{
		Id:         "sprite",
		Platform:   &stage.Canvas,
		Dimensions: genericTypes.Dimensions{},
		OutBoxDimensions: genericTypes.Dimensions{
			X:      10,
			Y:      10,
			Width:  88,
			Height: 150,
		},
		Ink: genericTypes.Ink{},
		Drag: basic.Drag{
			IsDraggable: true,
		},
	}

	i := factoryImage.NewImage(
		&stage.Canvas,
		&stage.ScratchPad,
		imgSpace,
		10,
		10,
		460*0.055,
		783*0.055,
		density,
		densityManager,
	)
	stage.Add(i.Draw)

	factoryGradient.ResetStylesGlobal(&stage.Canvas)
	factoryDraw.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
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

	bx2 = factoryDraw.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
		"bbox_2",
		20+50,
		50+50,
		100,
		100,
		10,
		0,
		shadowFilter,
		gradientFilter,
		density,
		densityManager,
	)

	rz := selectBox.NewResizeBoxFromBasicBox(
		bx2,
		-3,
		-3,
		6,
		6,
		1,
		density,
		densityManager,
	)

	factoryGradient.ResetStylesGlobal(&stage.Canvas)

	/*factoryImage.NewMultipleSpritesImage(
		&stage.Canvas,
		imgPlayer,
		48,
		60,
		0,
		7,
		90*time.Millisecond,
		50,
		70,
		48,
		60,
		density,
		densityManager,
	)*/

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

	browserDocument.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)
	mouse.AddFunctionPointer("size", rz.GetCollisionBox, rz.ProcessMousePosition)

	factoryTween.NewLinearFiniteLoop(
		time.Second*2,
		2,
		10.0,
		300.0,
		func(x, p float64, ars []interface{}) {
			i.SetX(x)
		},
		nil,
		func(x float64) {
			//fps.AddRunner(func() { i.SetX(x) }, true)
		},
	)

	<-done
}

var lastCursor bool
