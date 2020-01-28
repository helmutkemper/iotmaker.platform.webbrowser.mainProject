// +build js

//
package main

import (
	"fmt"
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorNames"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryFont"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factorySimpleBox"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryText"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage          *canvas.Stage
)

var eng engine.IEngine
var htmlElement iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace Html.Image

func prepareBeforeRun() {
	htmlElement = &Html.Html{}
	eng = &engine.Engine{}

	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		htmlElement,
		eng,
		browserDocument,
		"stage",
		density,
		densityManager,
	)
	stage.SetCursor(mouse.KCursorDefault)
	//stage.Engine(60)

	imgSpace = factoryBrowserImage.NewImage(
		htmlElement,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./spacecraft.png",
		},
		true,
		false,
	)
}

func main() {
	done := make(chan struct{})
	prepareBeforeRun()

	i := factoryImage.NewImage(
		"space",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		imgSpace.Get(),
		10,
		10,
		88,
		150,
		density,
		densityManager,
	)
	//i.SetDragMode(basic.KDragModeMobile)
	i.DragStart()
	stage.AddToDraw(i)

	rect := factorySimpleBox.NewBoxWithRoundedCorners(
		"boxDoKct",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		10,
		10,
		300,
		100,
		10,
		density,
		densityManager,
	)
	rect.DragStart()
	stage.AddToDraw(rect)

	t2 := factoryText.NewTextWithMaxWidth(
		"text",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
		nil,
		factoryColorNames.NewBlack(),
		factoryFont.NewFont(
			24.0,
			"px",
			factoryColorNames.NewBlack(),
			factoryFontFamily.NewArialBlack(),
			density,
			densityManager,
		),
		"Olá Mundo!",
		200,
		100,
		50,
		density,
		densityManager,
	)
	stage.AddToDraw(t2)

	t3 := factoryText.NewMeasureText(
		&stage.ScratchPad,
		factoryFont.NewFont(
			24.0,
			"px",
			factoryColorNames.NewBlack(),
			factoryFontFamily.NewArialBlack(),
			density,
			densityManager,
		),
		"Olá mundo",
	)
	fmt.Printf("width: %v\n", t3.Width)

	/*
		factoryTween.NewLinear(
			&engine.Engine{},
			time.Second*2,
			10.0,
			600.0,
			func(value float64, arguments ...interface{}) {
				fmt.Printf("onStartFunction()\n")
			},
			func(value float64, arguments ...interface{}) {
				//i.DragStart()
				fmt.Printf("onEndFunction()\n")
			},
			func(value float64, arguments ...interface{}) {
				fmt.Printf("onCycleStartFunction()\n")
			},
			func(value float64, arguments ...interface{}) {
				//i.DragStart()
				fmt.Printf("onCycleEndFunction()\n")
			},
			func(value float64, arguments ...interface{}) {
				fmt.Printf("onInvertFunction()\n")
			},
			func(value, percentToComplete float64, arguments ...interface{}) {
				i.Move(value, 10)
			},
			2,
		)
	*/

	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	<-done
}
