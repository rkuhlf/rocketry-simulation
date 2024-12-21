package simulations

import (
	"github.com/rkuhlf/rocketry-simulation/postprocessing"
	"github.com/rkuhlf/rocketry-simulation/rocketparts/pressurevessel"
	"github.com/rkuhlf/rocketry-simulation/simulators"
	"github.com/rkuhlf/rocketry-simulation/units"
)

func main() {
	vessel := pressurevessel.ConstantPressureVessel(
		60,
		units.AtmToPa(50),
	)

	sim := simulators.DrainSimulator(vessel, 0.1, simulators.ConstantFlowRate(2.5))

	res := sim.Simulate()
	postprocessing.SaveDrainSimulation(res)
	postprocessing.VisualizeDrainSimulation(res)
}
