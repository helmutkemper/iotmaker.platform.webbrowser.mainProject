// +build js

package main

import iotmaker_platform "github.com/helmutkemper/iotmaker.platform"

func main() {

	/*
	  el := js.NewElement()
	  d := el.Create("div", "test")
	  el.AppendChildToBody( d )
	*/

	s := iotmaker_platform.Stage{}
	s.NewStageOnTheRoot("stage_id")

}

/*
document := js.Global().Get("document")
p := document.Call("createElement", "p")
p.Set("innerHTML", "Hello WASM from Go!")
document.Get("body").Call("appendChild", p)
*/
