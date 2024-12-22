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
