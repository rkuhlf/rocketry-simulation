package chemistryvisualization

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/rkuhlf/rocketry-simulation/chemistry"
	"github.com/rkuhlf/rocketry-simulation/units"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func VisualizeNitrous(folder string) {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	p := plotDensity()

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, folder+"/density-plot.png"); err != nil {
		log.Fatalf("failed to save plot: %v", err)
	}

	// p = plotEnthalpy()

	// // Save the plot to a PNG file
	// if err := p.Save(6*vg.Inch, 4*vg.Inch, folder+"/enthalpy-plot.png"); err != nil {
	// 	log.Fatalf("failed to save plot: %v", err)
	// }
}

func plotDensity() *plot.Plot {
	p := plot.New()

	p.Title.Text = "Nitrous Density"
	p.X.Label.Text = "Temperature (C)"
	p.Y.Label.Text = "Density (kg/m^3)"

	vaporPoints := make(plotter.XYs, 0)
	liquidPoints := make(plotter.XYs, 0)
	for temp := float64(-89); temp < 36; temp += 1 {
		// Vapor density is always NaN.
		vaporDensity, liquidDensity := chemistry.NitrousProperties.Density(units.CelsiusToKelvin(temp))

		vaporPoint := plotter.XY{
			X: temp,
			Y: vaporDensity,
		}
		vaporPoints = append(vaporPoints, vaporPoint)

		liquidPoint := plotter.XY{
			X: temp,
			Y: liquidDensity,
		}
		liquidPoints = append(liquidPoints, liquidPoint)
	}

	line, err := plotter.NewLine(vaporPoints)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}
	line.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 237, A: 255}

	p.Add(line)
	p.Legend.Add("Vapor", line)

	line, err = plotter.NewLine(liquidPoints)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}
	line.LineStyle.Color = color.RGBA{R: 0, G: 237, B: 0, A: 255}

	p.Add(line)
	p.Legend.Add("Liquid", line)

	return p
}

func plotEnthalpy() *plot.Plot {
	p := plot.New()

	p.Title.Text = "Nitrous Enthalpy"
	p.X.Label.Text = "Temperature (C)"
	p.Y.Label.Text = "Enthalpy (kJ/kg)"

	vaporPoints := make(plotter.XYs, 0)
	liquidPoints := make(plotter.XYs, 0)
	for temp := float64(-89); temp < 36; temp += 1 {
		vaporDensity, liquidDensity := chemistry.NitrousProperties.Enthalpy(units.CelsiusToKelvin(temp))

		vaporPoint := plotter.XY{
			X: temp,
			Y: vaporDensity,
		}
		vaporPoints = append(vaporPoints, vaporPoint)

		liquidPoint := plotter.XY{
			X: temp,
			Y: liquidDensity,
		}
		liquidPoints = append(liquidPoints, liquidPoint)
	}

	line, err := plotter.NewLine(vaporPoints)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}
	line.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 237, A: 255}

	p.Add(line)
	p.Legend.Add("Vapor", line)

	line, err = plotter.NewLine(liquidPoints)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}
	line.LineStyle.Color = color.RGBA{R: 0, G: 237, B: 0, A: 255}

	p.Add(line)
	p.Legend.Add("Liquid", line)

	return p
}
