package simulations

import "math"

func Cylinder(diameter, height float64) float64 {
	return math.Pi * diameter * diameter / 4 * height
}
