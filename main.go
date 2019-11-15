// +build js

package main

import (
	js "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

func main() {

	/*
	  el := js.NewElement()
	  d := el.Create("div", "test")
	  el.AppendChildToBody( d )
	*/

	c := js.NewCanvasWith2DContext("canvas_id", 50.0, 50.0)
	c.BeginPath()
	//c.StrokeStyle("#FF0000")
	c.MoveTo(0.0, 0.0)
	c.LineTo(50.0, 50.0)
	c.Stroke()
	c.AppendToDocumentBody()

}

/*
document := js.Global().Get("document")
p := document.Call("createElement", "p")
p.Set("innerHTML", "Hello WASM from Go!")
document.Get("body").Call("appendChild", p)
*/
