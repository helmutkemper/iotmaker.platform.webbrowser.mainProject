// +build js

package main

import (
	iotmaker_platform_webbrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform/abstractType"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
	"image/color"
)

func main() {

	_, stage := iotmaker_platform_webbrowser.NewStage(
		"stage",
		300,
		300,
		1,
	)

	shadowBox := shadow.NewShadow(color.RGBA{R: 120, G: 0, B: 0, A: 255}, 10, 4, 4)

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
				LineWidth: 5,
				Density:   1.0,
			},
			Shadow: shadowBox,
			Gradient: abstractType.Gradient{
				Type: 0,
				ColorList: []abstractType.ColorStop{
					{
						Color: color.RGBA{
							R: 0,
							G: 0,
							B: 0,
							A: 255,
						},
						Stop: 0,
					},
					{
						Color: color.RGBA{
							R: 0,
							G: 0,
							B: 0,
							A: 0,
						},
						Stop: 1,
					},
				},
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
