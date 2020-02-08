// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.interfaces/iStage"
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontStyle"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/abstractType/text"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/dimensions"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorGradient"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorNames"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryFont"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryGradient"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryPoint"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factorySimpleBox"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryText"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/gravity"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
	"time"
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
	eng.SetFPS(60)

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

	f := factoryFont.NewFont(
		24.0,
		factoryFontFamily.NewVerdana(),
		factoryFontStyle.NewNotSet(),
		density,
		densityManager,
	)

	colorGradientBlack := factoryColorNames.NewBlack()
	colorGradientBlack.A = 30
	//colorGradientWhite := factoryColorNames.NewWhite()
	colorNotSet := factoryColorNames.NewNotSet()

	//shadowColor := factoryColorNames.NewBlack()
	/*shadow := factoryShadow.NewShadow(
		shadowColor,
		8,
		2,
		2,
		density,
		densityManager,
	)*/

	p0 := factoryPoint.NewPoint(
		0,
		0,
		density,
		densityManager,
	)
	p1 := factoryPoint.NewPoint(
		300,
		300,
		density,
		densityManager,
	)
	colorStop0 := factoryColorGradient.NewColorPosition(colorGradientBlack, 0.0)
	colorStop1 := factoryColorGradient.NewColorPosition(colorGradientBlack, 1.0)
	colorList := factoryColorGradient.NewColorList(colorStop0, colorStop1)

	gradient := factoryGradient.NewGradientLinearToFill(
		p0,
		p1,
		colorList,
	)
	inkSetup := &ink.Ink{
		LineWidth: 1,
		Shadow:    nil,
		Gradient:  gradient,
		Color:     colorNotSet,
	}

	t2 := factoryText.NewText(
		"text",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		f,
		"Draggable text and rocket! - bug chato do kct",
		50,
		50,
		density,
		densityManager,
	)
	t2.DragStart()
	stage.AddToDraw(t2)

	i := factoryImage.NewImage(
		"space",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		imgSpace.Get(),
		10,
		10,
		88,
		150,
		density,
		densityManager,
	)
	//i.SetDragMode(basic.KDragModeMobile)
	//todo:cursor hand on mouse over.
	i.DragStart()
	stage.AddToDraw(i)

	var dX = 0
	var dXAdjust = 0

	factoryTween.NewLinear(
		&engine.Engine{},
		time.Second*2,
		10.0,
		500.0,
		func(value float64, arguments ...interface{}) {
			//fmt.Printf("onStartFunction()\n")
		},
		func(value float64, arguments ...interface{}) {
			//i.DragStart()
			//fmt.Printf("onEndFunction()\n")
		}, //
		func(value float64, arguments ...interface{}) {
			//fmt.Printf("onCycleStartFunction()\n")
		},
		func(value float64, arguments ...interface{}) {
			//i.DragStart()
			//fmt.Printf("onCycleEndFunction()\n")
		},
		func(value float64, arguments ...interface{}) {
			//fmt.Printf("onInvertFunction()\n")
		},
		func(value, percentToComplete float64, arguments ...interface{}) {
			if i.DragIsDragging() {
				return
			}

			i.MoveX(int(value) + dXAdjust)
		},
		-1,
	)

	i.SetOnDragStartFunc(func(x, y int) {
		//fmt.Printf("OnDragStartFunc id: %v\n", i.Id)
	})

	i.SetOnDragEndFunc(func(x, y int) {
		//fmt.Printf("OnDragEndFunc id: %v\n", i.Id)
		dX, _ = i.GetDragDelta()
		dXAdjust += dX
		//fmt.Printf("dx: %v\n", dX)
	})

	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	/*font := factoryFont.NewFont(
		18,
		factoryFontFamily.NewVerdana(),
		factoryFontStyle.NewNotSet(),
		density,
		densityManager,
	)*/

	/*textToButton := factoryText.NewTextToButton(
			"textFomButton",
			stage,
			&stage.Canvas,
			&stage.ScratchPad,
	    inkColorBlack,
			font,
			"click-me!",
			density,
			densityManager,
		)*/

	/*Button(
		"button",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		[]basic.ISpriteBasicElement{textToButton},
		200,
		100,
		200,
		33,
		5,
		density,
		densityManager,
	)*/

	containerList := dimensions.Test()
	father := containerList[0]
	containerA := containerList[1]
	containerB := containerList[2]

	rectFather := factorySimpleBox.NewBoxWithRoundedCorners(
		"father",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		father.X,
		father.Y,
		father.Width,
		father.Height,
		0,
		density,
		densityManager,
	)
	stage.AddToDraw(rectFather)

	rectContainerA := factorySimpleBox.NewBoxWithRoundedCorners(
		"contaonerA",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		containerA.X,
		containerA.Y,
		containerA.Width,
		containerA.Height,
		0,
		density,
		densityManager,
	)
	stage.AddToDraw(rectContainerA)

	rectContainerB := factorySimpleBox.NewBoxWithRoundedCorners(
		"contaonerB",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		inkSetup,
		containerB.X,
		containerB.Y,
		containerB.Width,
		containerB.Height,
		0,
		density,
		densityManager,
	)
	stage.AddToDraw(rectContainerB)

	<-done
}

func Button(

	id string,
	stage iStage.IStage,
	platform,
	scratchPad iotmakerPlatformIDraw.IDraw,
	ink ink.Interface,
	list []basic.ISpriteBasicElement,
	x,
	y,
	width,
	height,
	border int,
	density interface{},
	densityManager coordinateManager.IDensity,

) {

	for _, element := range list {

		switch converted := element.(type) {
		case *text.Text:

			switch converted.Gravity {
			case gravity.KBottom:
			case gravity.KCenter:
			case gravity.KCenterHorizontal:
			case gravity.KCenterVertical:
			}

			stage.AddToDraw(converted)
		}

	}

	rect := factorySimpleBox.NewBoxWithRoundedCorners(
		id+"Rect",
		stage,
		platform,
		scratchPad,
		ink,
		x,
		y,
		width,
		height,
		border,
		density,
		densityManager,
	)
	stage.AddToDraw(rect)

}
