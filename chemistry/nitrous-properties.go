package chemistry

import (
	"math"
)

// Taken from https://web.archive.org/web/20220121041921/http://edge.rit.edu/edge/P07106/public/Nox.pdf
type nitrousProperties struct{}

// Vapor & Liquid Density at a given temperature.
// Vapor & Liquid enthalpy at a given temperature.
// molar mass.
// supercritical temperature

// In g/mol.
func (p *nitrousProperties) MolarMass() float64 {
	return 44.013
}

// In Kelvin.
func (p *nitrousProperties) CriticalTemperature() float64 {
	return 309.57
}

// In kg/mol.
func (p *nitrousProperties) CriticalDensity() float64 {
	return 452
}

// Returns vapor density and liquid density in kg/m^3.
func (p *nitrousProperties) Density(temp float64) (float64, float64) {
	T_r := temp / p.CriticalTemperature()

	b_1 := -1.00900
	b_2 := -6.28792
	b_3 := 7.50332
	b_4 := -7.90463
	b_5 := 0.629427

	third := float64(1) / 3
	power := b_1*math.Pow(1/T_r-1, third) + b_2*math.Pow(1/T_r-1, 2*third) + b_3*(1/T_r-1) + b_4*math.Pow(1/T_r-1, 4*third) + b_5*math.Pow(1/T_r-1, 5*third)
	vaporDensity := p.CriticalDensity() * math.Exp(power)

	b_1 = 1.72328
	b_2 = -0.83950
	b_3 = 0.51060
	b_4 = -0.10412

	power = b_1*math.Pow(1-T_r, third) + b_2*math.Pow(1-T_r, 2*third) + b_3*(1-T_r) + b_4*math.Pow(1-T_r, 4*third)
	liquidDensity := p.CriticalDensity() * math.Exp(power)

	return vaporDensity, liquidDensity
}

// Returns vapor density and liquid enthalpy in kJ/kg.
func (p *nitrousProperties) Enthalpy(temp float64) (float64, float64) {
	T_r := temp / p.CriticalTemperature()

	b_1 := -200.
	b_2 := 440.055
	b_3 := -459.701
	b_4 := 434.081
	b_5 := -485.338

	third := float64(1) / 3
	vaporEnthalpy := b_1 + b_2*math.Pow(1-T_r, third) + b_3*math.Pow(1-T_r, 2*third) + b_4*(1-T_r) + b_5*math.Pow(1-T_r, 4*third)

	b_1 = -200.
	b_2 = 116.043
	b_3 = -917.225
	b_4 = 794.779
	b_5 = -589.587
	liquidEnthalpy := b_1 + b_2*math.Pow(1-T_r, third) + b_3*math.Pow(1-T_r, 2*third) + b_4*(1-T_r) + b_5*math.Pow(1-T_r, 4*third)

	return vaporEnthalpy, liquidEnthalpy
}

var NitrousProperties = nitrousProperties{}
