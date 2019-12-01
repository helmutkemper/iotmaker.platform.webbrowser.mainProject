// +build js

package main

import (
	webBrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform/abstractType"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
)

func main() {

	_, stage := webBrowser.NewStage(
		"stage",
		300,
		300,
		1,
	)

	colorDarkBlue := colornames.Darkblue
	colorDarkBlue.A = 0xff

	shadowFilter := shadow.NewShadowFilter(colorDarkBlue, 5, 2, 2)
	coordinateBox := gradient.NewCoordinate(0, 0, 120, 150)
	colorWhite := gradient.NewColorPosition(colornames.WhiteTransparent, 0.2)
	colorBlack := gradient.NewColorPosition(colornames.Black, 1)
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
			//Shadow:   &shadowFilter,
			//Gradient: &gradientBox,
		},
	)

	abstractType.NewBasicBox(
		abstractType.BasicBox{
			Platform: &stage.Canvas,
			Id:       "line",
			Dimensions: abstractType.DimensionsBasicBox{
				X:         20 + 50,
				Y:         50 + 50,
				Width:     100,
				Height:    100,
				Border:    5,
				LineWidth: 5,
				Density:   1.0,
			},
			Shadow:   shadowFilter,
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
