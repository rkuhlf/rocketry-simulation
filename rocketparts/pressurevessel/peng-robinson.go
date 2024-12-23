package pressurevessel

import (
	"errors"
	"fmt"
	"log"
	"math"

	"github.com/khezen/rootfinding"
)

type prPressureVessel struct {
	// Measured in kg.
	fluidMass float64
	// Measured in Pa.
	pressure float64
	// Measured in K.
	temperature float64
	// in g/mol
	molarMass float64
	// in m^3
	volume float64
}

/**
* Assumes constant temperature.
 */
func PrPressureVessel(fluidMass, pressure, temperature, molarMass, volume float64) *prPressureVessel {
	return &prPressureVessel{
		fluidMass:   fluidMass,
		pressure:    pressure,
		temperature: temperature,
		molarMass:   molarMass,
		volume:      volume,
	}
}

/**
* Takes the temperature of the mass change.
 */
func (p *prPressureVessel) UpdateState(timeStep float64, massChange float64) error {
	if p.fluidMass < massChange {
		return fmt.Errorf("could not update state to have a negative fluid mass")
	}

	p.fluidMass += massChange

	return nil
}

// Returns in m^3 / mol.
func (p *prPressureVessel) molarVolume() (float64, error) {
	if p.fluidMass == 0 {
		return 0, errors.New("division by zero error computing molar volume")
	}
	// Divide by 1000 to convert the molar mass into kg / mol.
	moles := p.fluidMass / (p.molarMass / 1000)
	return p.volume / moles, nil
}

func (p *prPressureVessel) Pressure() float64 {
	if p.fluidMass == 0 {
		return 0
	}

	v, err := p.molarVolume()
	if err != nil {
		log.Fatalf("Unexpected error from molarVolume %v", err)
	}

	return pressure_pr_eos(v, p.temperature)
}

func (p *prPressureVessel) FluidMass() float64 {
	return p.fluidMass
}

// These properties are from chatGPT.
var R = 8.314     // Universal gas constant, J/(mol·K)
var Tc = 304.2    // Critical temperature, K
var Pc = 7.38e6   // Critical pressure, Pa
var omega = 0.225 // Acentric factor

// Calculate PR EOS parameters for CO2
var b = 0.07780 * R * Tc / Pc
var a = 0.45724 * math.Pow(R*Tc, 2) / Pc

func alpha(temp float64) float64 {
	kappa := 0.37464 + 1.54226*omega - 0.26992*math.Pow(omega, 2)
	Tr := temp / Tc // Reduced temperature
	return math.Pow(1+kappa*(1-math.Sqrt(Tr)), 2)
}

// Compute pressure using Peng-Robinson EOS.
func pressure_pr_eos(v, temp float64) float64 {
	a_T := a * alpha(temp)
	return R*temp/(v-b) - a_T/(v*(v+b)+b*(v-b))
}

// Solve for molar volume using Peng-Robinson EOS.
func volume_pr_eos(P, temp float64) float64 {
	a_T := a * alpha(temp)

	// Define the cubic equation for v: P = f(v)
	pr_cubic := func(v float64) float64 {
		return (R*temp/(v-b) - a_T/(v*(v+b)+b*(v-b))) - P
	}

	// v_guess := R * temp / P

	root, err := rootfinding.Brent(pr_cubic, 0, 1000, 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(root)

	return root
}
