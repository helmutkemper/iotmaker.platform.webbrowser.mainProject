// +build js

package main

import (
	iotmaker_platform_IDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
	"image/color"
)

var (
	colorShadow    color.RGBA
	filter         iotmaker_platform_IDraw.IFilterShadowInterface
	density                                   = 3.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	bx2                                       = &basicBox.BasicBox{}
	stage                                     = canvas.Stage{}
	gradientFilter iotmaker_platform_IDraw.IFilterGradientInterface
)

func main() {

	done := make(chan struct{}, 0)

	browserDocument := document.NewDocument()

	colorShadow = colornames.DarkblueTransparent
	filter = shadow.NewShadowFilter(colorShadow, 5, 2, 2, density, densityManager)

	//mouse.AddFunctionPointer(bx1.GetAlphaChannel)

	stage = canvas.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	colorWhite := gradient.NewColorPosition(colornames.WhiteTransparent, 0.5)
	colorBlack := gradient.NewColorPosition(colornames.Black, 1)
	colorList := gradient.NewColorList(colorBlack, colorWhite)

	coordinateP0 := gradient.NewPoint(0, 0, density, densityManager)
	coordinateP1 := gradient.NewPoint(120, 150, density, densityManager)
	gradientFilter = gradient.NewGradientLinearToStroke(coordinateP0, coordinateP1, colorList)

	basicBox.NewBasicBox(&stage.Canvas, &stage.ScratchPad, "bbox_1", 20, 50, 100, 100, 5, 5, nil, nil, density, densityManager)
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

	browserDocument.SetMouseMoveListener(mouse.GetDefaultFunction())
	mouse.AddFunctionPointer(bx2.GetAlphaChannel, bateu)

	<-done
}

func bateu(x, y int) {

	colorShadow = colornames.Darkred
	filter = shadow.NewShadowFilter(colorShadow, 5, 2, 2, density, densityManager)

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
}
