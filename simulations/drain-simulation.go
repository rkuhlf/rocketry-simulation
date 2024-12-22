package simulations

import (
	"github.com/rkuhlf/rocketry-simulation/rocketparts/pressurevessel"
	"github.com/rkuhlf/rocketry-simulation/simulations/postprocessing"
	"github.com/rkuhlf/rocketry-simulation/simulations/simulators"
	"github.com/rkuhlf/rocketry-simulation/units"
)

func DrainSimulation() {
	vessel := pressurevessel.ConstantPressureVessel(
		60,
		units.AtmToPa(50),
	)

	sim := simulators.DrainSimulator(vessel, 0.1, simulators.ConstantFlowRate(-2.5))

	res := sim.Simulate()
	postprocessing.SaveDrainSimulation(res, "./.local/drain-output.csv")
	postprocessing.VisualizeDrainSimulation(res, "./.local/drain")
}

func DrainSimulationBattery() {
	fluidMass := float64(60) // kg
	startPressure := units.AtmToPa(50)
	startTemperature := units.FahrenheitToKelvin(73)
	nitrousMolarMass := 44.013 // g/mol
	tankVolume := Cylinder(
		units.InchesToMeters(7), // 7" diameter
		units.FeetToMeters(10),  // 10 ft high
	) // Volume in m^3

	vessels := map[string]simulators.Vessel{
		"Constant Pressure": pressurevessel.ConstantPressureVessel(
			fluidMass,
			startPressure,
		),
		"PR Constant Temperature": pressurevessel.PrPressureVessel(
			fluidMass,
			startPressure,
			startTemperature,
			nitrousMolarMass,
			tankVolume,
		),
	}

	for name, vessel := range vessels {
		sim := simulators.DrainSimulator(vessel, 0.1, simulators.ConstantFlowRate(-2.5))

		res := sim.Simulate()
		postprocessing.SaveDrainSimulation(res, "./.local/drain-output-"+name+".csv")
		postprocessing.VisualizeDrainSimulation(res, "./.local/drain"+name)
	}
}
