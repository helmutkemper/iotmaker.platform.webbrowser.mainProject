// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	webBrowserMouse "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mouse"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage          *canvas.Stage
)

var htmlElement iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace Html.Image

func prepareBeforeRun() {
	htmlElement = &Html.Html{}
	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		htmlElement,
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
	done := make(chan struct{}, 0)
	prepareBeforeRun()

	i := factoryImage.NewImage(
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
	i.DragStart()
	stage.AddToStage(i.Draw)

	browserDocument.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	<-done
}
