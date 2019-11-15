// +build js

package main

import (
	js "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

func main() {

	el := js.NewExistentElementById("mycanvas")
	el.NewCanvas("canvas_id")
	//c := js.NewCanvasWith2DContext("MyCanvas", 50.0, 50.0)
	el.Get()
	c.LineTo(50.0, 50.0)
	c.StrokeStyle("#000000")
	c.Stroke()
}
