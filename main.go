// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/eventMouse"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	webBrowserMouse "github.com/helmutkemper/iotmaker.platform.webbrowser/mouse"
	"github.com/helmutkemper/iotmaker.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.platform/fps"
	"github.com/helmutkemper/iotmaker.platform/mouse"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = canvas.Stage{}
)

var htmlElement iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace Html.Image

func prepareDataBeforeRun() {
	htmlElement = &Html.Html{}
	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		800,
		600,
		density,
		densityManager,
	)

	Html.PreLoadCursor(browserDocument.SelfDocument, Html.KTemplarianPath, Html.KTemplarianList)
	for _, v := range Html.PreLoadMouseList {
		htmlElement.Append(browserDocument.SelfDocument, v.Img.Get())
	}

	imgSpace = factoryBrowserImage.NewImage(
		htmlElement,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "visibleMousePointer",
			"src": "./fonts/Templarian/MaterialDesign/svg/cursor-default-outline.svg",
		},
		true,
		true,
	)
}

func main() {

	done := make(chan struct{}, 0)
	prepareDataBeforeRun()
	fps.Set(60)

	i := factoryImage.NewImage(
		&stage.Canvas,
		&stage.ScratchPad,
		imgSpace.Get(),
		-100,
		-100,
		24,
		24,
		density,
		densityManager,
	)
	i.SetDraggable(true)
	stage.Add(i.Draw)

	htmlElement.Remove(browserDocument.SelfDocument, imgSpace.Get())

	browserDocument.AddEventListener(eventMouse.KMouseMove, webBrowserMouse.SetMouseMoveManager(mouse.ManagerMouseMove))
	//mouse.AddFunctionPointer("bBox2", bx2.GetCollisionBox, bateu)

	<-done
}
