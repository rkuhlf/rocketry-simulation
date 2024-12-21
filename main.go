package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Create a new plot
	p := plot.New()

	// Set plot title and labels
	p.Title.Text = "Example Plot"
	p.X.Label.Text = "X-axis"
	p.Y.Label.Text = "Y-axis"

	// Generate data points (y = x^2)
	points := make(plotter.XYs, 11)
	for i := 0; i <= 10; i++ {
		x := float64(i)
		points[i].X = x
		points[i].Y = x * x
	}

	// Create a line plot from the points
	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatalf("failed to create line plot: %v", err)
	}

	// Add the line plot to the plot
	p.Add(line)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		log.Fatalf("failed to save plot: %v", err)
	}

	log.Println("Plot saved to plot.png")
}
