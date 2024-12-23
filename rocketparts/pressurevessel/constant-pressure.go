package pressurevessel

import (
	"fmt"
)

type ConstantPressureVessel struct {
	// Measured in kg.
	fluidMass float64
	// Measured in Pa.
	pressure float64
}

func NewConstantPressureVessel(fluidMass float64, pressure float64) *ConstantPressureVessel {
	return &ConstantPressureVessel{
		fluidMass: fluidMass,
		pressure:  pressure,
	}
}

func (p *ConstantPressureVessel) UpdateState(massChange float64) error {
	if p.fluidMass < massChange {
		return fmt.Errorf("could not update state to have a negative fluid mass")
	}
	p.fluidMass += massChange

	return nil
}

func (p *ConstantPressureVessel) Pressure() float64 {
	return p.pressure
}

func (p *ConstantPressureVessel) FluidMass() float64 {
	return p.fluidMass
}
