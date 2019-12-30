// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/Html"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.platform/mathUtil"
	"time"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = canvas.Stage{}
)

var html iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace interface{}

func prepareDataBeforeRun() {

	html = &Html.Html{}
	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		800,
		600,
		density,
		densityManager,
	)

	imgSpace = factoryBrowserImage.NewImage(
		html,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		false,
	)
}

func main() {

	done := make(chan struct{}, 0)
	prepareDataBeforeRun()

	for a := 0; a != 3000; a += 1 {
		i := factoryImage.NewImage(
			&stage.Canvas,
			&stage.ScratchPad,
			imgSpace,
			-100,
			-100,
			29,
			50,
			density,
			densityManager,
		)
		//i.SetDraggable(true)
		stage.Add(i.Draw)
		factoryTween.NewSelectRandom(
			time.Millisecond*time.Duration(mathUtil.Float64FomInt(500, 3000)),
			mathUtil.Float64FomInt(0, 800),
			mathUtil.Float64FomInt(0, 800),
			nil,
			nil,
			nil,
			func(x, p float64, ars ...interface{}) {
				i.Dimensions.X = x
				i.OutBoxDimensions.X = x
			},
			0,
			nil,
		)
		factoryTween.NewSelectRandom(
			time.Millisecond*time.Duration(mathUtil.Float64FomInt(500, 3000)),
			mathUtil.Float64FomInt(0, 600),
			mathUtil.Float64FomInt(0, 600),
			nil,
			nil,
			nil,
			func(y, p float64, ars ...interface{}) {
				i.Dimensions.Y = y
				i.OutBoxDimensions.Y = y
			},
			0,
			nil,
		)
	}

	<-done
}
