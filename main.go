// +build js

package main

import (
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	webBrowser "github.com/helmutkemper/iotmaker.platform.webbrowser"
	"github.com/helmutkemper/iotmaker.platform/abstractType"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/gradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
)

func main() {

	var density = 2.0
	var densityManager coordinateManager.IDensity = &coordinateManager.Density{}

	stage := webBrowser.NewStage(
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	colorDarkBlue := colornames.Darkblue

	shadowFilter := shadow.NewShadowFilter(colorDarkBlue, 5, 2, 2, density, densityManager)
	coordinateP0 := gradient.NewPoint(0, 0, density, densityManager)
	coordinateP1 := gradient.NewPoint(120, 150, density, densityManager)
	colorWhite := gradient.NewColorPosition(colornames.WhiteTransparent, 0.2)
	colorBlack := gradient.NewColorPosition(colornames.Black, 1)
	colorList := gradient.NewColorList(colorBlack, colorWhite)
	gradientBox := gradient.NewGradientLinearToStroke(coordinateP0, coordinateP1, colorList)

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
				Density:   density,
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
				Density:   density,
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
