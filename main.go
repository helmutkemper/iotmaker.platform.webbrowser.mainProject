// +build js

package main

import (
	iotmaker_platform_webbrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform/abstractType"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
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
	coordinateBox := gradient.NewCoordinate(0, 0, 120, 150)
	colorWhite := gradient.NewColorPosition(color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00}, 0.2)
	colorBlack := gradient.NewColorPosition(color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}, 1)
	colorList := gradient.NewColorList(colorBlack, colorWhite)
	gradientBox := gradient.NewStrokeGradientLinear(coordinateBox, colorList)

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
			Shadow:   shadowBox,
			Gradient: gradientBox,
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
