package pressurevessel

import (
	"fmt"
)

type constantPressureVessel struct {
	// Measured in kg.
	fluidMass float64
	// Measured in Pa
	pressure float64
}

func ConstantPressureVessel(fluidMass float64, pressure float64) *constantPressureVessel {
	return &constantPressureVessel{
		fluidMass: fluidMass,
		pressure:  pressure,
	}
}

func (p *constantPressureVessel) UpdateState(timeStep float64, massChange float64) error {
	if p.fluidMass < massChange {
		return fmt.Errorf("Could not update state to have a negative fluid mass.")
	}
	p.fluidMass += massChange

	return nil
}

func (p *constantPressureVessel) Pressure() float64 {
	return p.pressure
}

func (p *constantPressureVessel) FluidMass() float64 {
	return p.fluidMass
}
