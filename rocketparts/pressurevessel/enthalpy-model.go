package pressurevessel

import (
	"fmt"
	"math"
)

// Things I need from the fluid:
// Vapor & Liquid Density at a given temperature.
// Vapor & Liquid enthalpy at a given temperature.
// molar mass.
// supercritical temperature

type phase int

const (
	// Supercritical happens beyond a certain temperature, and we have to use different models (generallY).
	Supercritical phase = iota
	// We only have gas when the temperature is not so high as to be supercritical, but there isn't enough mass for a liquid to form.
	GasOnly
	// We only have liquid in the case that the mass is so high that a gas is not allowed to form. This is very bad, because then the liquid is usually exerting a lot of pressure on the vessel trying to contain it.
	LiquidOnly
	// This is the typical case.
	Equilibrium
)

type EnthalpyPressureVessel struct {
	fluidProperties fluidProperties
	// Measured in kg.
	fluidMass float64
	// Measured in joules.
	enthalpy float64
	// in m^3
	volume float64
}

type fluidProperties interface {
	MolarMass() float64
	CriticalTemperature() float64
	Density(float64) (float64, float64)
	Enthalpy(float64) (float64, float64)
}

// Supplies the initial temperature.
func NewEnthalpyPressureVessel(fluidMass, temperature, volume float64, fluidProperties fluidProperties) *EnthalpyPressureVessel {
	vaporDensity, liquidDensity := fluidProperties.Density(temperature)

	quality, _ := computeQuality(fluidMass, temperature, volume, fluidProperties.MolarMass(), vaporDensity, liquidDensity, fluidProperties.CriticalTemperature())

	// TODO: Handle it appropriately depending on what phase it is in - I need to be able to make some approximation of the non-saturated enthalpy.

	vaporSpecificEnthalpy, liquidSpecificEnthalpy := fluidProperties.Enthalpy(temperature)
	totalEnthalpy := fluidMass*quality*vaporSpecificEnthalpy + fluidMass*(1-quality)*liquidSpecificEnthalpy

	fmt.Println(totalEnthalpy)

	return &EnthalpyPressureVessel{
		fluidMass:       fluidMass,
		volume:          volume,
		fluidProperties: fluidProperties,
		enthalpy:        totalEnthalpy,
	}
}

// The temperature of the mass being added or removed.
func (p *EnthalpyPressureVessel) UpdateState(massChange float64, temperature float64) error {
	if p.fluidMass+massChange < 0 {
		return fmt.Errorf("could not update state to have a negative fluid mass")
	}
	p.fluidMass += massChange

	return nil
}

func (p *EnthalpyPressureVessel) Pressure() float64 {
	return 0
}

func (p *EnthalpyPressureVessel) FluidMass() float64 {
	return p.fluidMass
}

func (p *EnthalpyPressureVessel) Temperature() float64 {
	return p.fluidMass
}

// Used only at initialization, this calculates the fraction of the mass that is vapor.
// We hold the temperature fixed.
// The densityFunc gives the density of a liquid and its vapor when it's at the given temperature and given time to saturate, assuming the amount of mass is appropriate.
func computeQuality(fluidMass, temperature, volume, molarMass, vaporDensity, liquidDensity, supercriticalTemperature float64) (float64, phase) {
	// Do a bisection search for the quality that gets us sufficiently close to the correct volume.
	// We need a check that we aren't supercritical.
	if temperature > supercriticalTemperature {
		return 1, Supercritical
	}

	// If the vapor by itself would not generate sufficient volume, we know it's all vapor.
	if 1/vaporDensity*fluidMass < volume {
		return 1, GasOnly
	}

	// If the liquid by itself generates too much volume, we know it's not in vapor-liquid equilibrium; instead it should all be water.
	if 1/liquidDensity*fluidMass > volume {
		return 0, LiquidOnly
	}

	// We use the property that vapor is always less dense than liquid.
	quality := bisectionSearch(func(estimateQuality float64) float64 {
		estimateVaporVolume := 1 / vaporDensity * fluidMass * estimateQuality
		estimateLiquidVolume := 1 / liquidDensity * fluidMass * (1 - estimateQuality)
		return estimateLiquidVolume + estimateVaporVolume
	}, volume, float64(0), float64(1), 0.001)

	return quality, Equilibrium
}

// f must be an increasing function.
func bisectionSearch(f func(float64) float64, target, min, max, precisionFraction float64) float64 {
	for {
		input := (min + max) / 2

		output := f(input)

		// If we've converged, we break out of the loop. Convergence criteria is 0.1%
		if errorFraction(output, target) < precisionFraction {
			return input
		}

		// If our output is too big, we need to decrease the input.
		if output > target {
			max = input
		} else {
			// Otherwise, we need to increase the input
			min = input
		}
	}
}

func errorFraction(trueValue, measuredValue float64) float64 {
	// Calculate the absolute difference between the true value and measured value
	absoluteError := math.Abs(trueValue - measuredValue)

	// Calculate the percent error
	errorFraction := (absoluteError / math.Abs(trueValue))

	return errorFraction
}
