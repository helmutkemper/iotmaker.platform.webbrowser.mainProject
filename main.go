// +build js

package main

import (
	"fmt"
	iot_p "github.com/helmutkemper/iotmaker.platform"
)

func main() {

	/*
	  el := js.NewElement()
	  d := el.Create("div", "test")
	  el.AppendChildToBody( d )
	*/

	fmt.Printf("Olá mundo!")
	s := iot_p.Stage{}
	s.NewStageOnTheRoot("stage_id")

}

/*
document := js.Global().Get("document")
p := document.Call("createElement", "p")
p.Set("innerHTML", "Hello WASM from Go!")
document.Get("body").Call("appendChild", p)
*/
