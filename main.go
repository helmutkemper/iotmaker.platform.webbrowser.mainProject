// +build js

package main

import (
	iotmaker_platform_webbrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

func main() {

	_, stage := iotmaker_platform_webbrowser.NewStage(
		"stage",
		300,
		300,
		1,
	)

	//var draw iotmaker_platform.ICanvas = &iotmaker_platform.Draw{}
	//abstractType.NewBasicBox(draw)

	stage.Canvas.BeginPath()
	stage.Canvas.MoveTo(0, 0)
	stage.Canvas.LineTo(300, 300)
	stage.Canvas.Stroke()

}
