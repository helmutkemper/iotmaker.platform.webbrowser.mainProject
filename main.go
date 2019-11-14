// +build js

package main

import (
	js "github.com/helmutkemper/iotmaker.platform.webbrowser"
)

func main() {

	js.NewCanvasWith2DContext(50.0, 50.0)

}
