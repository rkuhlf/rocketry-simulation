package simulators

import (
	"fmt"
)

type Vessel interface {
	UpdateState(timeStep float64, massChange float64) error
	FluidMass() float64
	Pressure() float64
}

type flowRateFunc func(Vessel) float64

/**
* Holds the setup of the simulation.
* The simulation can be run over and over again by calling Simulate().
 */
type drainSimulator struct {
	vessel   Vessel
	timeStep float64
	// Takes the vessel and returns what the flow rate out of it is.
	flowRateFunc flowRateFunc
}

func DrainSimulator(vessel Vessel, timeStep float64, flowRateFunc flowRateFunc) *drainSimulator {
	return &drainSimulator{
		timeStep:     timeStep,
		vessel:       vessel,
		flowRateFunc: flowRateFunc,
	}
}

func (d *drainSimulator) Simulate() []DrainSimulationRecord {
	fmt.Println("Starting drain simulation")
	result := make([]DrainSimulationRecord, 0)
	currentTime := float64(0)
	iters := 0

	for {
		massChange := d.flowRateFunc(d.vessel) * d.timeStep
		if d.vessel.FluidMass()+massChange < 0 {
			fmt.Printf("Stopping drain because massChange=%f and fluidMass=%f", massChange, d.vessel.FluidMass())
			break
		}

		d.vessel.UpdateState(d.timeStep, massChange)

		currentTime += d.timeStep
		record := DrainSimulationRecord{
			currentTime,
			d.vessel.FluidMass(),
			d.vessel.Pressure(),
		}

		result = append(result, record)

		iters++
		if iters%50 == 0 {
			fmt.Println(record)
		}
	}

	fmt.Printf("Finished drain simulation. Collected %d records. \n", len(result))

	return result
}

type DrainSimulationRecord struct {
	Time float64
	Mass float64
	// Pressure in Pa.
	Pressure float64
}

func (d DrainSimulationRecord) String() string {
	return fmt.Sprintf("{ Time: %.2f s, Mass: %.2f kg, Pressure: %.2f Pa }", d.Time, d.Mass, d.Pressure)
}
