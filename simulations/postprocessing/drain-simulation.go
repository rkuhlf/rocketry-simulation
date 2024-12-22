package postprocessing

import (
	"fmt"
	"log"
	"os"

	"github.com/rkuhlf/rocketry-simulation/simulations/simulators"
	"github.com/rkuhlf/rocketry-simulation/units"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func SaveDrainSimulation(data []simulators.DrainSimulationRecord, output string) {
	res := make([]map[string]string, 0, len(data))
	for _, row := range data {
		res = append(res, map[string]string{
			"Time (s)":        fmt.Sprintf("%f", row.Time),
			"Fluid Mass (kg)": fmt.Sprintf("%f", row.Mass),
			"Pressure (atm)":  fmt.Sprintf("%f", units.PaToAtm(row.Pressure)),
		})
	}

	saveResultToFile(res, output)
}

func VisualizeDrainSimulation(data []simulators.DrainSimulationRecord, folder string) {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	p := plotMass(data)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, folder+"/mass-plot.png"); err != nil {
		log.Fatalf("failed to save plot: %v", err)
	}

	p = plotPressure(data)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, folder+"/pressure-plot.png"); err != nil {
		log.Fatalf("failed to save plot: %v", err)
	}
}

func plotMass(data []simulators.DrainSimulationRecord) *plot.Plot {
	p := plot.New()

	p.Title.Text = "Tank Drain"
	p.X.Label.Text = "Time (s)"
	p.Y.Label.Text = "Mass (kg)"

	points := make(plotter.XYs, len(data))
	for i, row := range data {
		points[i].X = row.Time
		points[i].Y = row.Mass
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}

	p.Add(line)

	return p
}

func plotPressure(data []simulators.DrainSimulationRecord) *plot.Plot {
	p := plot.New()

	p.Title.Text = "Tank Drain"
	p.X.Label.Text = "Time (s)"
	p.Y.Label.Text = "Pressure (atm)"

	points := make(plotter.XYs, len(data))
	for i, row := range data {
		points[i].X = row.Time
		points[i].Y = units.PaToAtm(row.Pressure)
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}

	p.Add(line)

	return p
}
