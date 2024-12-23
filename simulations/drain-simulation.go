package simulations

import (
	"github.com/rkuhlf/rocketry-simulation/chemistry"
	"github.com/rkuhlf/rocketry-simulation/rocketparts/pressurevessel"
	"github.com/rkuhlf/rocketry-simulation/simulations/postprocessing"
	"github.com/rkuhlf/rocketry-simulation/simulations/simulators"
	"github.com/rkuhlf/rocketry-simulation/units"
)

var updateEnthalpyPressureVesselAdapter = func(vessel *pressurevessel.EnthalpyPressureVessel, timeStep, massChange float64) error {
	return vessel.UpdateState(massChange, vessel.Temperature())
}

var updateConstantPressureVesselAdapter = func(vessel *pressurevessel.ConstantPressureVessel, timeStep, massChange float64) error {
	return vessel.UpdateState(massChange)
}

var updatePrPressureVesselAdapter = func(vessel *pressurevessel.PrPressureVessel, timeStep, massChange float64) error {
	return vessel.UpdateState(massChange)
}

func DrainSimulation() {
	startTemperature := units.FahrenheitToKelvin(73)
	tankVolume := Cylinder(
		units.InchesToMeters(7), // 7" diameter
		units.FeetToMeters(10),  // 10' high
	) // Volume in m^3

	vessel := pressurevessel.NewEnthalpyPressureVessel(
		60,
		startTemperature,
		tankVolume,
		&chemistry.NitrousProperties,
	)

	sim := simulators.NewDrainSimulator(vessel, updateEnthalpyPressureVesselAdapter, 0.1, simulators.ConstantFlowRate(-2.5))

	res := sim.Simulate()
	postprocessing.SaveDrainSimulation(res, "./.local/drain-output.csv")
	postprocessing.VisualizeDrainSimulation(res, "./.local/drain")
}

type Simulator interface {
	Simulate() []simulators.DrainSimulationRecord
}

func DrainSimulationBattery() {
	timeStep := 0.1 // s
	flowRate := simulators.ConstantFlowRate(-2.5)
	fluidMass := float64(60) // kg
	startPressure := units.AtmToPa(50)
	startTemperature := units.FahrenheitToKelvin(73)
	nitrousMolarMass := 44.013 // g/mol
	tankVolume := Cylinder(
		units.InchesToMeters(7), // 7" diameter
		units.FeetToMeters(10),  // 10' high
	) // Volume in m^3

	sims := map[string]Simulator{
		"Constant Pressure": simulators.NewDrainSimulator(pressurevessel.NewConstantPressureVessel(
			fluidMass,
			startPressure,
		), updateConstantPressureVesselAdapter, timeStep, flowRate),
		"PR Constant Temperature": simulators.NewDrainSimulator(pressurevessel.NewPrPressureVessel(
			fluidMass,
			startPressure,
			startTemperature,
			nitrousMolarMass,
			tankVolume,
		), updatePrPressureVesselAdapter, timeStep, flowRate),
	}

	for name, sim := range sims {
		res := sim.Simulate()
		postprocessing.SaveDrainSimulation(res, "./.local/drain-output-"+name+".csv")
		postprocessing.VisualizeDrainSimulation(res, "./.local/drain"+name)
	}
}
