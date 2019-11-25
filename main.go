// +build js

package main

import (
	iotmaker_platform_webbrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform/abstractType"
	"image/color"
)

func main() {

	_, stage := iotmaker_platform_webbrowser.NewStage(
		"stage",
		300,
		300,
		1,
	)

	abstractType.NewBasicBox(
		abstractType.BasicBox{
			Platform: &stage.Canvas,
			Id:       "line",
			Dimensions: abstractType.DimensionsBasicBox{
				X:         20,
				Y:         50,
				Width:     100,
				Height:    100,
				Border:    5,
				LineWidth: 2,
				Density:   1.0,
			},
			Shadow: abstractType.Shadow{
				Color: color.RGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 50,
				},
				Blur:    3,
				OffsetX: 1,
				OffsetY: 1,
			},
		},
	)

	abstractType.NewLineTo(
		&stage.Canvas,
		"line",
		1.0,
		0,
		0,
		300,
		300,
		1,
	)

}
